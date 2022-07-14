package twitter

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"sort"
	"strings"
	"sync"
	"time"
)

type TAuth struct {
	Token, Ct0, Bearer string
}

type Api struct {
	Client   *http.Client
	CacheDir string
	workQ    chan dlItem
	mu       sync.Mutex
}

func NewApi(cacheDir string) *Api {
	return &Api{
		Client:   http.DefaultClient,
		CacheDir: cacheDir,
		workQ:    make(chan dlItem),
		mu:       sync.Mutex{},
	}
}

type TwitAuthor struct {
	ScreenName string `json:"screen_name"`
	Name       string `json:"name"`
	Url        string `json:"url"`
	PictureUrl string `json:"picture_url"`
}

type TwitEntity struct {
	Key      string     `json:"key"`
	Url      string     `json:"url"`
	Type     string     `json:"type"`
	Size     []int      `json:"size"`
	Author   TwitAuthor `json:"author"`
	Likes    int        `json:"likes"`
	Retweets int        `json:"retweets"`
	MediaUrl string     `json:"media_url"`
	TweetUrl string     `json:"tweet_url"`
	Date     time.Time  `json:"date"`
}

type TwitResponse struct {
	Entities []*TwitEntity `json:"entities"`
	Cursor   string        `json:"cursor"`
	Eof      bool          `json:"eof"`
}

type Progress struct {
	Done  int `json:"done"`
	Total int `json:"total"`
}

const (
	idCachePrefixSize = 2
	userAgent         = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.53 Safari/537.36"
)

func hashOf(url string) string {
	s := sha1.New()
	s.Write([]byte(url))
	return hex.EncodeToString(s.Sum(nil))
}

type dlItem struct {
	key string
	url string
	wg  *sync.WaitGroup
}

func cacheKey(e string) string {
	return e[len(e)-idCachePrefixSize:] + "/" + e
}

type MediaGetter = func(ctx context.Context, dateFrom, dateTo *time.Time, cursor string, auth TAuth, prog chan<- Progress) (*TwitResponse, error)

func (a *Api) scheduleCache(key, url string, wg *sync.WaitGroup) string {
	wg.Add(1)
	a.workQ <- dlItem{
		key: key,
		url: url,
		wg:  wg,
	}
	return cacheKey(key)
}

func (a *Api) extractMedias(author *User, tweet *Tweet, wg *sync.WaitGroup) (out []*TwitEntity) {
	if tweet.ExtendedEntities == nil || len(tweet.ExtendedEntities.Media) < 1 {
		return []*TwitEntity{}
	}
	date, _ := time.Parse("Jan 02 15:04:05 -0700 2006", tweet.CreatedAt)
	base := &TwitEntity{
		Author: TwitAuthor{
			ScreenName: author.ScreenName,
			Name:       author.Name,
			Url:        "https://twitter.com/" + author.ScreenName,
			PictureUrl: a.scheduleCache(hashOf(author.ProfileImageURLHTTPS), author.ProfileImageURLHTTPS, wg),
		},
		Likes:    tweet.FavoriteCount,
		Retweets: tweet.RetweetCount,
		Date:     date,
		TweetUrl: "https://twitter.com/" + author.ScreenName + "/status/" + tweet.IDStr,
	}
	for _, media := range tweet.ExtendedEntities.Media {
		entity := *base
		entity.Key = media.IDStr
		entity.Type = media.Type
		entity.MediaUrl = media.ExpandedURL
		entity.Size = []int{media.Sizes.Large.Width, media.Sizes.Large.Height}
		switch media.Type {
		case "video":
			streams := media.VideoInfo.Variants
			sort.Slice(streams, func(i, j int) bool {
				return streams[i].Bitrate > streams[j].Bitrate
			})
			video := streams[0]
			entity.Url = a.scheduleCache(media.IDStr, video.URL, wg)
		case "photo":
			entity.Url = a.scheduleCache(media.IDStr, media.MediaURLHTTPS, wg)
		default:
			continue
		}
		out = append(out, &entity)
	}
	return out
}

func (a *Api) GetHomeMedia(ctx context.Context, _dateFrom, _dateTo *time.Time, cursor string, auth TAuth, prog chan<- Progress) (*TwitResponse, error) {
	tl, err := a.getHomeTimeline(ctx, cursor, auth)
	if err != nil {
		return nil, err
	}
	wg := sync.WaitGroup{}
	entries := tl.Data.Home.HomeTimelineUrt.Instructions[0].Entries
	total := len(entries)
	prog <- Progress{
		Done:  0,
		Total: total,
	}
	resp := &TwitResponse{Entities: []*TwitEntity{}, Eof: len(entries) <= 2}
	for _, e := range entries {
		total -= 1
		switch e.Content.EntryType {
		case "TimelineTimelineItem":
			if e.Content.ItemContent.ItemType != "TimelineTweet" {
				continue
			}
			if e.Content.ClientEventInfo == nil {
				// Ad?
				continue
			}
			result := e.Content.ItemContent.TweetResults.Result
			if result.Typename != "Tweet" {
				// Suspended account?
				continue
			}
			user := result.Core.UserResults.Result.Legacy
			tweet := tweetOrRetweet(result.Legacy)
			media := a.extractMedias(user, tweet, &wg)
			resp.Entities = append(resp.Entities, media...)
			total += len(media)
			break
		case "TimelineTimelineCursor":
			if e.Content.CursorType == "Bottom" {
				resp.Cursor = e.Content.Value
			}
			break
		}
		prog <- Progress{
			Done:  len(resp.Entities),
			Total: total,
		}
	}
	wg.Wait()
	return resp, nil
}

func (a *Api) GetSearchMedia(ctx context.Context, dateFrom, dateTo *time.Time, cursor string, auth TAuth, prog chan<- Progress) (*TwitResponse, error) {
	tl, err := a.getSearchTimeline(ctx, dateFrom, dateTo, cursor, auth)
	if err != nil {
		return nil, err
	}
	wg := sync.WaitGroup{}
	resp := &TwitResponse{Entities: []*TwitEntity{}}
	var entries []*SearchEntry
	for _, ins := range tl.Timeline.Instructions {
		if ins.AddEntries != nil {
			for _, entry := range ins.AddEntries.Entries {
				if entry.EntryID == "sq-cursor-bottom" {
					resp.Cursor = entry.Content.Operation.Cursor.Value
				} else if entry.Content.Item != nil {
					entries = append(entries, entry)
				}
			}
		} else if ins.ReplaceEntry != nil && ins.ReplaceEntry.EntryIDToReplace == "sq-cursor-bottom" {
			resp.Cursor = ins.ReplaceEntry.Entry.Content.Operation.Cursor.Value
		}
	}
	total := len(entries)
	resp.Eof = total == 0
	prog <- Progress{
		Done:  0,
		Total: total,
	}
	tweetMap := tl.GlobalObjects.Tweets
	userMap := tl.GlobalObjects.Users
	for _, e := range entries {
		total -= 1
		if e.Content.Item == nil {
			continue
		}
		tweetRef := e.Content.Item.Content.Tweet
		tweet, ok := tweetMap[tweetRef.ID]
		if !ok {
			continue
		}
		tweet = tweetOrRetweet(tweet)
		user, ok := userMap[tweet.UserIDStr]
		if !ok {
			continue
		}
		media := a.extractMedias(user, tweet, &wg)
		resp.Entities = append(resp.Entities, media...)
		total += len(media)
		prog <- Progress{
			Done:  len(resp.Entities),
			Total: total,
		}
	}
	wg.Wait()
	return resp, nil
}

func (a *Api) Start(workers int) {
	for i := 0; i < workers; i++ {
		go a.dlWorker()
	}
}

func (a *Api) dlWorker() {
	doIt := func(item dlItem) {
		defer item.wg.Done()
		tail := cacheKey(item.key)
		fname := path.Join(a.CacheDir, tail)
		a.mu.Lock()
		_, err := os.Stat(fname)
		if err == nil {
			a.mu.Unlock()
			return
		}
		_ = os.WriteFile(fname, []byte{}, 0o644)
		a.mu.Unlock()
		req, err := http.NewRequestWithContext(context.Background(), "GET", item.url, nil)
		if err != nil {
			return
		}
		req.Header.Set("User-Agent", userAgent)
		resp, err := a.Client.Do(req)
		if err != nil {
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return
		}
		os.MkdirAll(path.Dir(fname), 0o777)
		if err = os.WriteFile(fname, body, 0o644); err != nil {
			log.Printf("could not write media: %s", err)
		}
	}
	for item := range a.workQ {
		doIt(item)
	}
}

func (a *Api) addHeaders(referer string, auth TAuth, h *http.Header) {
	h.Set("Authority", "twitter.com")
	h.Set("Accept", "*/*")
	h.Set("Accept-Language", "en-US;q=0.8,en;q=0.7")
	h.Set("Authorization", "Bearer "+auth.Bearer)
	h.Set("Cache-Control", "no-cache")
	h.Set("Content-Type", "application/json")
	h.Set("Cookie", strings.NewReplacer("AUTHTOKEN", auth.Token, "Ct0", auth.Ct0).Replace(`ads_prefs="HBISAAA="; remember_checked_on=1; auth_token=AUTHTOKEN; ct0=Ct0; eu_cn=1; d_prefs=MjoxLGNvbnNlbnRfdmVyc2lvbjoyLHRleHRfdmVyc2lvbjoxMDAw; tweetdeck_version=legacy`))
	h.Set("Origin", "https://twitter.com")
	h.Set("Pragma", "no-cache")
	h.Set("Referer", referer)
	h.Set("Sec-Ch-Ua", "\"Chromium\";v=\"103\", \".Not/A)Brand\";v=\"99\"")
	h.Set("Sec-Ch-Ua-Mobile", "?0")
	h.Set("Sec-Ch-Ua-Platform", "\"Linux\"")
	h.Set("Sec-Fetch-Dest", "empty")
	h.Set("Sec-Fetch-Mode", "cors")
	h.Set("Sec-Fetch-Site", "same-origin")
	h.Set("User-Agent", userAgent)
	h.Set("X-Csrf-Token", auth.Ct0)
	h.Set("X-Twitter-Active-User", "yes")
	h.Set("X-Twitter-Auth-Type", "OAuth2Session")
	h.Set("X-Twitter-Client-Language", "en")
}

func (a *Api) getHomeTimeline(ctx context.Context, cursor string, auth TAuth) (*HomeResponse, error) {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
	data := HomeRequest{
		Variables: &HomeVariables{
			Count:                       20,
			Cursor:                      cursor,
			IncludePromotedContent:      true,
			LatestControlAvailable:      true,
			WithSuperFollowsUserFields:  true,
			WithDownvotePerspective:     true,
			WithReactionsMetadata:       false,
			WithReactionsPerspective:    false,
			WithSuperFollowsTweetFields: true,
			SeenTweetIds:                []string{"1545882924544827392"},
		},
		Features: &HomeFeatures{
			DontMentionMeViewAPIEnabled:      true,
			InteractiveTextEnabled:           true,
			ResponsiveWebUcGqlEnabled:        false,
			VibeTweetContextEnabled:          false,
			ResponsiveWebEditTweetAPIEnabled: false,
			StandardizedNudgesMisinfo:        false,
			ResponsiveWebEnhanceCardsEnabled: false,
		},
		QueryID: "iEMFB7-_JPy1LG00EWbYUw",
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequestWithContext(ctx, "POST", "https://twitter.com/i/api/twitter/iEMFB7-_JPy1LG00EWbYUw/HomeLatestTimeline", body)
	if err != nil {
		return nil, err
	}
	a.addHeaders("https://twitter.com/home", auth, &req.Header)
	resp, err := a.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("HTTP %d", resp.StatusCode))
	}
	defer resp.Body.Close()
	var response HomeResponse
	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (a *Api) getSearchTimeline(ctx context.Context, dateFrom, dateTo *time.Time, cursor string, auth TAuth) (*SearchResponse, error) {
	p := url.Values{}
	p.Add("include_profile_interstitial_type", "1")
	p.Add("include_blocking", "1")
	p.Add("include_blocked_by", "1")
	p.Add("include_followed_by", "1")
	p.Add("include_want_retweets", "1")
	p.Add("include_mute_edge", "1")
	p.Add("include_can_dm", "1")
	p.Add("include_can_media_tag", "1")
	p.Add("include_ext_has_nft_avatar", "1")
	p.Add("skip_status", "1")
	p.Add("cards_platform", "Web-12")
	p.Add("include_cards", "1")
	p.Add("include_ext_alt_text", "true")
	p.Add("include_quote_count", "true")
	p.Add("include_reply_count", "1")
	p.Add("tweet_mode", "extended")
	p.Add("include_ext_collab_control", "true")
	p.Add("include_entities", "true")
	p.Add("include_user_entities", "true")
	p.Add("include_ext_media_color", "true")
	p.Add("include_ext_media_availability", "true")
	p.Add("include_ext_sensitive_media_warning", "true")
	p.Add("include_ext_trusted_friends_metadata", "true")
	p.Add("send_error_codes", "true")
	p.Add("simple_quoted_tweet", "true")
	p.Add("pc", "1")
	p.Add("spelling_corrections", "1")
	p.Add("include_ext_edit_control", "false")
	p.Add("ext", "mediaStats,highlightedLabel,hasNftAvatar,replyvotingDownvotePerspective,voiceInfo,enrichments,superFollowMetadata,unmentionInfo,collab_control")
	p.Add("query_source", "typed_query")
	// Recent, chronological, not "top results".
	p.Add("tweet_search_mode", "live")
	// Start date.
	q := "-filter:replies"
	if dateFrom != nil {
		q += " since:" + dateFrom.Format("2006-01-02")
	}
	if dateTo != nil {
		q += " until:" + dateTo.Format("2006-01-02")
	}
	p.Add("q", q)
	// The important bit: follows only.
	p.Add("social_filter", "searcher_follows")
	p.Add("count", "20")
	if cursor != "" {
		p.Add("cursor", cursor)
	}
	req, err := http.NewRequestWithContext(ctx, "GET", "https://twitter.com/i/api/2/search/adaptive.json?"+p.Encode(), nil)
	if err != nil {
		return nil, err
	}
	a.addHeaders("https://twitter.com/search?", auth, &req.Header)
	resp, err := a.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("HTTP %d", resp.StatusCode))
	}
	defer resp.Body.Close()
	var response SearchResponse
	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}

func tweetOrRetweet(t *Tweet) *Tweet {
	if t.RetweetedStatus != nil && t.RetweetedStatus.Result != nil {
		return t.RetweetedStatus.Result.Legacy
	}
	return t
}
