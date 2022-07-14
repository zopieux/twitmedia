package twitter

import "encoding/json"

type HomeRequest struct {
	Variables *HomeVariables `json:"variables"`
	Features  *HomeFeatures  `json:"features"`
	QueryID   string         `json:"queryId"`
}
type HomeVariables struct {
	Count                       int      `json:"count"`
	Cursor                      string   `json:"cursor,omitempty"`
	IncludePromotedContent      bool     `json:"includePromotedContent"`
	LatestControlAvailable      bool     `json:"latestControlAvailable"`
	WithSuperFollowsUserFields  bool     `json:"withSuperFollowsUserFields"`
	WithDownvotePerspective     bool     `json:"withDownvotePerspective"`
	WithReactionsMetadata       bool     `json:"withReactionsMetadata"`
	WithReactionsPerspective    bool     `json:"withReactionsPerspective"`
	WithSuperFollowsTweetFields bool     `json:"withSuperFollowsTweetFields"`
	SeenTweetIds                []string `json:"seenTweetIds,omitempty"`
}
type HomeFeatures struct {
	DontMentionMeViewAPIEnabled      bool `json:"dont_mention_me_view_api_enabled"`
	InteractiveTextEnabled           bool `json:"interactive_text_enabled"`
	ResponsiveWebUcGqlEnabled        bool `json:"responsive_web_uc_gql_enabled"`
	VibeTweetContextEnabled          bool `json:"vibe_tweet_context_enabled"`
	ResponsiveWebEditTweetAPIEnabled bool `json:"responsive_web_edit_tweet_api_enabled"`
	StandardizedNudgesMisinfo        bool `json:"standardized_nudges_misinfo"`
	ResponsiveWebEnhanceCardsEnabled bool `json:"responsive_web_enhance_cards_enabled"`
}

type HomeResponse struct {
	Data *HomeResponseData `json:"data"`
}
type AffiliatesHighlightedLabel struct {
}
type EntityUrl struct {
	DisplayURL  string `json:"display_url"`
	ExpandedURL string `json:"expanded_url"`
	URL         string `json:"url"`
	Indices     []int  `json:"indices"`
}
type EntityURL struct {
	Urls []*EntityUrl `json:"urls,omitempty"`
}
type UserEntities struct {
	Description *EntityURL `json:"description,omitempty"`
	URL         *EntityURL `json:"url,omitempty"`
}
type Rgb struct {
	Blue  int `json:"blue"`
	Green int `json:"green"`
	Red   int `json:"red"`
}
type Palette struct {
	Percentage float64 `json:"percentage"`
	Rgb        *Rgb    `json:"rgb"`
}
type Ok struct {
	Palette []*Palette `json:"palette,omitempty"`
}
type R = json.RawMessage
type MediaColor struct {
	R *R `json:"r"`
}
type ProfileImageExtensionsMediaColor struct {
	Palette []*Palette `json:"palette,omitempty"`
}
type ProfileBannerExtensionsMediaColor struct {
	Palette []*Palette `json:"palette,omitempty"`
}
type ProfileBannerExtensions struct {
	MediaColor *MediaColor `json:"mediaColor,omitempty"`
	MediaStats *RAndTTL    `json:"mediaStats,omitempty"`
}
type ProfileImageExtensions struct {
	MediaColor *MediaColor `json:"mediaColor"`
	MediaStats *RAndTTL    `json:"mediaStats,omitempty"`
}
type User struct {
	ID                                           int                                `json:"id,omitempty"`
	IDStr                                        string                             `json:"id_str,omitempty"`
	Name                                         string                             `json:"name,omitempty"`
	ScreenName                                   string                             `json:"screen_name,omitempty"`
	Location                                     string                             `json:"location,omitempty"`
	Description                                  string                             `json:"description,omitempty"`
	URL                                          string                             `json:"url,omitempty"`
	Entities                                     *UserEntities                      `json:"entities,omitempty"`
	Protected                                    bool                               `json:"protected,omitempty"`
	FollowersCount                               int                                `json:"followers_count,omitempty"`
	FastFollowersCount                           int                                `json:"fast_followers_count,omitempty"`
	NormalFollowersCount                         int                                `json:"normal_followers_count,omitempty"`
	FriendsCount                                 int                                `json:"friends_count,omitempty"`
	ListedCount                                  int                                `json:"listed_count,omitempty"`
	CreatedAt                                    string                             `json:"created_at,omitempty"`
	FavouritesCount                              int                                `json:"favourites_count,omitempty"`
	UtcOffset                                    interface{}                        `json:"utc_offset,omitempty"`
	TimeZone                                     interface{}                        `json:"time_zone,omitempty"`
	GeoEnabled                                   bool                               `json:"geo_enabled,omitempty"`
	Verified                                     bool                               `json:"verified,omitempty"`
	StatusesCount                                int                                `json:"statuses_count,omitempty"`
	MediaCount                                   int                                `json:"media_count,omitempty"`
	Lang                                         interface{}                        `json:"lang,omitempty"`
	ContributorsEnabled                          bool                               `json:"contributors_enabled,omitempty"`
	IsTranslator                                 bool                               `json:"is_translator,omitempty"`
	IsTranslationEnabled                         bool                               `json:"is_translation_enabled,omitempty"`
	ProfileBackgroundColor                       string                             `json:"profile_background_color,omitempty"`
	ProfileBackgroundImageURL                    string                             `json:"profile_background_image_url,omitempty"`
	ProfileBackgroundImageURLHTTPS               string                             `json:"profile_background_image_url_https,omitempty"`
	ProfileBackgroundTile                        bool                               `json:"profile_background_tile,omitempty"`
	ProfileImageURL                              string                             `json:"profile_image_url,omitempty"`
	ProfileImageURLHTTPS                         string                             `json:"profile_image_url_https,omitempty"`
	ProfileBannerURL                             string                             `json:"profile_banner_url,omitempty"`
	ProfileImageExtensionsSensitiveMediaWarning  interface{}                        `json:"profile_image_extensions_sensitive_media_warning,omitempty"`
	ProfileImageExtensionsMediaAvailability      interface{}                        `json:"profile_image_extensions_media_availability,omitempty"`
	ProfileImageExtensionsAltText                interface{}                        `json:"profile_image_extensions_alt_text,omitempty"`
	ProfileImageExtensionsMediaColor             *ProfileImageExtensionsMediaColor  `json:"profile_image_extensions_media_color,omitempty"`
	ProfileImageExtensions                       *ProfileImageExtensions            `json:"profile_image_extensions,omitempty"`
	ProfileBannerExtensionsSensitiveMediaWarning interface{}                        `json:"profile_banner_extensions_sensitive_media_warning,omitempty"`
	ProfileBannerExtensionsMediaAvailability     interface{}                        `json:"profile_banner_extensions_media_availability,omitempty"`
	ProfileBannerExtensionsAltText               interface{}                        `json:"profile_banner_extensions_alt_text,omitempty"`
	ProfileBannerExtensionsMediaColor            *ProfileBannerExtensionsMediaColor `json:"profile_banner_extensions_media_color,omitempty"`
	ProfileBannerExtensions                      *ProfileBannerExtensions           `json:"profile_banner_extensions,omitempty"`
	ProfileLinkColor                             string                             `json:"profile_link_color,omitempty"`
	ProfileSidebarBorderColor                    string                             `json:"profile_sidebar_border_color,omitempty"`
	ProfileSidebarFillColor                      string                             `json:"profile_sidebar_fill_color,omitempty"`
	ProfileTextColor                             string                             `json:"profile_text_color,omitempty"`
	ProfileUseBackgroundImage                    bool                               `json:"profile_use_background_image,omitempty"`
	HasExtendedProfile                           bool                               `json:"has_extended_profile,omitempty"`
	DefaultProfile                               bool                               `json:"default_profile,omitempty"`
	DefaultProfileImage                          bool                               `json:"default_profile_image,omitempty"`
	PinnedTweetIds                               []int64                            `json:"pinned_tweet_ids,omitempty"`
	PinnedTweetIdsStr                            []string                           `json:"pinned_tweet_ids_str,omitempty"`
	HasCustomTimelines                           bool                               `json:"has_custom_timelines,omitempty"`
	CanDm                                        bool                               `json:"can_dm,omitempty"`
	CanMediaTag                                  bool                               `json:"can_media_tag,omitempty"`
	Following                                    bool                               `json:"following,omitempty"`
	FollowRequestSent                            bool                               `json:"follow_request_sent,omitempty"`
	Notifications                                bool                               `json:"notifications,omitempty"`
	Muting                                       bool                               `json:"muting,omitempty"`
	Blocking                                     bool                               `json:"blocking,omitempty"`
	BlockedBy                                    bool                               `json:"blocked_by,omitempty"`
	WantRetweets                                 bool                               `json:"want_retweets,omitempty"`
	AdvertiserAccountType                        string                             `json:"advertiser_account_type,omitempty"`
	AdvertiserAccountServiceLevels               []interface{}                      `json:"advertiser_account_service_levels,omitempty"`
	ProfileInterstitialType                      string                             `json:"profile_interstitial_type,omitempty"`
	BusinessProfileState                         string                             `json:"business_profile_state,omitempty"`
	TranslatorType                               string                             `json:"translator_type,omitempty"`
	WithheldInCountries                          []interface{}                      `json:"withheld_in_countries,omitempty"`
	FollowedBy                                   bool                               `json:"followed_by,omitempty"`
	ExtHasNftAvatar                              bool                               `json:"ext_has_nft_avatar,omitempty"`
	Ext                                          *UserExt                           `json:"ext,omitempty"`
	RequireSomeConsent                           bool                               `json:"require_some_consent,omitempty"`
}
type Professional struct {
	RestID           string        `json:"rest_id"`
	ProfessionalType string        `json:"professional_type"`
	Category         []interface{} `json:"category"`
}
type Result struct {
	Typename                   string                      `json:"__typename"`
	ID                         string                      `json:"key"`
	RestID                     string                      `json:"rest_id"`
	AffiliatesHighlightedLabel *AffiliatesHighlightedLabel `json:"affiliates_highlighted_label"`
	HasNftAvatar               bool                        `json:"has_nft_avatar"`
	Legacy                     *User                       `json:"legacy"`
	Professional               *Professional               `json:"professional"`
	SuperFollowEligible        bool                        `json:"super_follow_eligible"`
	SuperFollowedBy            bool                        `json:"super_followed_by"`
	SuperFollowing             bool                        `json:"super_following"`
}
type UserResults struct {
	Result *Result `json:"result"`
}
type Core struct {
	UserResults *UserResults `json:"user_results"`
}
type DownvotePerspective struct {
	IsDownvoted bool `json:"isDownvoted"`
}
type Size struct {
	Height int    `json:"h"`
	Width  int    `json:"w"`
	Resize string `json:"resize"`
}
type Sizes struct {
	Large  *Size `json:"large"`
	Medium *Size `json:"medium"`
	Small  *Size `json:"small"`
	Thumb  *Size `json:"thumb"`
}
type OriginalInfo struct {
	Height     int           `json:"height"`
	Width      int           `json:"width"`
	FocusRects []*FocusRects `json:"focus_rects,omitempty"`
}
type AdditionalMediaInfo struct {
	Monetizable bool `json:"monetizable"`
}
type ExtMediaColor struct {
	Palette []Palette `json:"palette"`
}
type ExtMediaAvailability struct {
	Status string `json:"status"`
}
type Variants struct {
	Bitrate     int    `json:"bitrate,omitempty"`
	ContentType string `json:"content_type"`
	URL         string `json:"url"`
}
type VideoInfo struct {
	AspectRatio    []int       `json:"aspect_ratio"`
	DurationMillis int         `json:"duration_millis"`
	Variants       []*Variants `json:"variants"`
}
type Media struct {
	ID                       int64                 `json:"id,omitempty"`
	IDStr                    string                `json:"id_str"`
	DisplayURL               string                `json:"display_url"`
	ExpandedURL              string                `json:"expanded_url"`
	Indices                  []int                 `json:"indices"`
	MediaKey                 string                `json:"media_key"`
	MediaURL                 string                `json:"media_url"`
	MediaURLHTTPS            string                `json:"media_url_https"`
	Type                     string                `json:"type"`
	URL                      string                `json:"url"`
	AdditionalMediaInfo      *AdditionalMediaInfo  `json:"additional_media_info"`
	MediaStats               *RAndTTL              `json:"mediaStats"`
	ExtMediaColor            *ExtMediaColor        `json:"ext_media_color"`
	ExtMediaAvailability     *ExtMediaAvailability `json:"ext_media_availability"`
	ExtSensitiveMediaWarning interface{}           `json:"ext_sensitive_media_warning,omitempty"`
	Features                 *HomeFeatures         `json:"features"`
	Sizes                    *Sizes                `json:"sizes"`
	OriginalInfo             *OriginalInfo         `json:"original_info"`
	VideoInfo                *VideoInfo            `json:"video_info"`
	ExtAltText               interface{}           `json:"ext_alt_text,omitempty"`
	Ext                      *MediaExt             `json:"ext,omitempty"`
}
type ExtendedEntities struct {
	Media []*Media `json:"media"`
}
type Tweet struct {
	CreatedAt                 string            `json:"created_at"`
	ConversationIDStr         string            `json:"conversation_id_str"`
	DisplayTextRange          []int             `json:"display_text_range"`
	Entities                  *TweetEntities    `json:"entities,omitempty"`
	ExtendedEntities          *ExtendedEntities `json:"extended_entities,omitempty"`
	FavoriteCount             int               `json:"favorite_count"`
	Favorited                 bool              `json:"favorited"`
	FullText                  string            `json:"full_text"`
	IsQuoteStatus             bool              `json:"is_quote_status"`
	Lang                      string            `json:"lang"`
	PossiblySensitive         bool              `json:"possibly_sensitive"`
	PossiblySensitiveEditable bool              `json:"possibly_sensitive_editable"`
	QuoteCount                int               `json:"quote_count"`
	ReplyCount                int               `json:"reply_count"`
	RetweetCount              int               `json:"retweet_count"`
	Retweeted                 bool              `json:"retweeted"`
	Source                    string            `json:"source"`
	UserIDStr                 string            `json:"user_id_str"`
	IDStr                     string            `json:"id_str"`
	ID                        int64             `json:"id,omitempty"`
	Truncated                 bool              `json:"truncated,omitempty"`
	InReplyToStatusID         interface{}       `json:"in_reply_to_status_id,omitempty"`
	InReplyToStatusIDStr      interface{}       `json:"in_reply_to_status_id_str,omitempty"`
	InReplyToUserID           interface{}       `json:"in_reply_to_user_id,omitempty"`
	InReplyToUserIDStr        interface{}       `json:"in_reply_to_user_id_str,omitempty"`
	InReplyToScreenName       interface{}       `json:"in_reply_to_screen_name,omitempty"`
	UserID                    int64             `json:"user_id,omitempty"`
	Geo                       interface{}       `json:"geo,omitempty"`
	Coordinates               interface{}       `json:"coordinates,omitempty"`
	Place                     interface{}       `json:"place,omitempty"`
	Contributors              interface{}       `json:"contributors,omitempty"`
	ConversationID            int64             `json:"conversation_id,omitempty"`
	SupplementalLanguage      interface{}       `json:"supplemental_language,omitempty"`
	Ext                       *TweetExt         `json:"ext,omitempty"`
	SelfThread                *SelfThread       `json:"self_thread,omitempty"`
	RetweetedStatus           *RetweetedStatus  `json:"retweeted_status_result,omitempty"`
}
type RetweetedStatus struct {
	Result *TweetResult `json:"result,omitempty"`
}
type TweetResult struct {
	Typename            string               `json:"__typename"`
	RestID              string               `json:"rest_id"`
	Core                *Core                `json:"core"`
	DownvotePerspective *DownvotePerspective `json:"downvotePerspective"`
	UnmentionInfo       *RAndTTL             `json:"unmention_info"`
	Legacy              *Tweet               `json:"legacy"`
}
type TweetResults struct {
	Result *TweetResult `json:"result"`
}
type ItemContent struct {
	ItemType         string        `json:"itemType"`
	TweetResults     *TweetResults `json:"tweet_results"`
	TweetDisplayType string        `json:"tweetDisplayType"`
}
type TimelinesDetails struct {
	InjectionType  string `json:"injectionType"`
	ControllerData string `json:"controllerData"`
}
type ClientEventInfoDetails struct {
	TimelinesDetails *TimelinesDetails `json:"timelinesDetails"`
}
type ClientEventInfo struct {
	Component string                  `json:"component,omitempty"`
	Element   string                  `json:"element,omitempty"`
	Details   *ClientEventInfoDetails `json:"detail,omitemptys"`
}
type ContentItem struct {
	ItemContent     *ItemContent     `json:"itemContent"`
	ClientEventInfo *ClientEventInfo `json:"clientEventInfo"`
}
type ContentItems struct {
	EntryID     string       `json:"entryId"`
	Dispensable bool         `json:"dispensable"`
	Item        *ContentItem `json:"item"`
}
type ConversationMetadata struct {
	AllTweetIds         []string `json:"allTweetIds"`
	EnableDeduplication bool     `json:"enableDeduplication"`
}
type Metadata struct {
	ConversationMetadata *ConversationMetadata `json:"conversationMetadata"`
}
type Content struct {
	ItemContent         *ItemContent     `json:"itemContent"`
	EntryType           string           `json:"entryType"`
	Items               []*ContentItems  `json:"items"`
	Metadata            *Metadata        `json:"metadata"`
	DisplayType         string           `json:"displayType"`
	ClientEventInfo     *ClientEventInfo `json:"clientEventInfo"`
	Value               string           `json:"value"`
	CursorType          string           `json:"cursorType"`
	StopOnEmptyResponse bool             `json:"stopOnEmptyResponse"`
}
type TlEntries struct {
	EntryID   string   `json:"entryId"`
	SortIndex string   `json:"sortIndex"`
	Content   *Content `json:"content,omitempty"`
}
type Instructions struct {
	Type    string       `json:"type"`
	Entries []*TlEntries `json:"entries"`
}
type ResponseObjects struct {
	FeedbackActions    []interface{} `json:"feedbackActions"`
	ImmediateReactions []interface{} `json:"immediateReactions"`
}
type HomeTimelineUrt struct {
	Instructions    []*Instructions  `json:"instructions"`
	ResponseObjects *ResponseObjects `json:"responseObjects"`
}
type Home struct {
	HomeTimelineUrt *HomeTimelineUrt `json:"home_timeline_urt"`
}
type HomeResponseData struct {
	Home *Home `json:"home"`
}

type SearchResponse struct {
	GlobalObjects *SearchObjects  `json:"globalObjects,omitempty"`
	Timeline      *SearchTimeline `json:"timeline,omitempty"`
}

type SearchEntities struct {
	Hashtags     []interface{} `json:"hashtags,omitempty"`
	Symbols      []interface{} `json:"symbols,omitempty"`
	UserMentions []interface{} `json:"user_mentions,omitempty"`
	Urls         []interface{} `json:"urls,omitempty"`
}
type RAndTTL struct {
	R         *R  `json:"r,omitempty"`
	TTL       int `json:"ttl,omitempty"`
	ViewCount int `json:"viewCount,omitempty"`
}
type Hashtags struct {
	Text    string `json:"text,omitempty"`
	Indices []int  `json:"indices,omitempty"`
}
type UserMentions struct {
	ScreenName string `json:"screen_name,omitempty"`
	Name       string `json:"name,omitempty"`
	ID         int64  `json:"id,omitempty"`
	IDStr      string `json:"id_str,omitempty"`
	Indices    []int  `json:"indices,omitempty"`
}
type FocusRects struct {
	X      int `json:"x,omitempty"`
	Y      int `json:"y,omitempty"`
	Width  int `json:"w,omitempty"`
	Height int `json:"h,omitempty"`
}
type Feature struct {
	Faces []*FocusRects `json:"faces,omitempty"`
}
type Features struct {
	Large  *Feature `json:"large,omitempty"`
	Orig   *Feature `json:"orig,omitempty"`
	Medium *Feature `json:"medium,omitempty"`
	Small  *Feature `json:"small,omitempty"`
}
type TweetEntities struct {
	Hashtags     []*Hashtags     `json:"hashtags,omitempty"`
	Symbols      []interface{}   `json:"symbols,omitempty"`
	UserMentions []*UserMentions `json:"user_mentions,omitempty"`
	Urls         []interface{}   `json:"urls,omitempty"`
	Media        []*Media        `json:"media,omitempty"`
}
type MediaExt struct {
	MediaStats *RAndTTL `json:"mediaStats,omitempty"`
}
type Urls struct {
	URL         string `json:"url,omitempty"`
	ExpandedURL string `json:"expanded_url,omitempty"`
	DisplayURL  string `json:"display_url,omitempty"`
	Indices     []int  `json:"indices,omitempty"`
}
type TypedURL struct {
	Type        string `json:"type,omitempty"`
	StringValue string `json:"string_value,omitempty"`
	ScribeKey   string `json:"scribe_key,omitempty"`
}
type TypeAndString struct {
	Type        string `json:"type,omitempty"`
	StringValue string `json:"string_value,omitempty"`
}
type UserValue struct {
	IDStr string        `json:"id_str,omitempty"`
	Path  []interface{} `json:"path,omitempty"`
}
type Site struct {
	Type      string     `json:"type,omitempty"`
	UserValue *UserValue `json:"user_value,omitempty"`
	ScribeKey string     `json:"scribe_key,omitempty"`
}
type ImageValue struct {
	URL    string `json:"url,omitempty"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
	Alt    string `json:"alt,omitempty"`
}
type Thumbnail struct {
	Type       string      `json:"type,omitempty"`
	ImageValue *ImageValue `json:"image_value,omitempty"`
}
type ImageColorValue struct {
	Palette []*Palette `json:"palette,omitempty"`
}
type ThumbnailImageColor struct {
	Type            string           `json:"type,omitempty"`
	ImageColorValue *ImageColorValue `json:"image_color_value,omitempty"`
}
type BindingValues struct {
	VanityURL              TypedURL            `json:"vanity_url,omitempty"`
	CardURL                TypedURL            `json:"card_url,omitempty"`
	Domain                 TypeAndString       `json:"domain,omitempty"`
	Title                  TypeAndString       `json:"title,omitempty"`
	ThumbnailImageAltText  TypeAndString       `json:"thumbnail_image_alt_text,omitempty"`
	Description            TypeAndString       `json:"description,omitempty"`
	Site                   Site                `json:"site,omitempty"`
	ThumbnailImageSmall    Thumbnail           `json:"thumbnail_image_small,omitempty"`
	ThumbnailImage         Thumbnail           `json:"thumbnail_image,omitempty"`
	ThumbnailImageLarge    Thumbnail           `json:"thumbnail_image_large,omitempty"`
	ThumbnailImageXLarge   Thumbnail           `json:"thumbnail_image_x_large,omitempty"`
	ThumbnailImageOriginal Thumbnail           `json:"thumbnail_image_original,omitempty"`
	ThumbnailImageColor    ThumbnailImageColor `json:"thumbnail_image_color,omitempty"`
}
type Device struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}
type Audience struct {
	Name   string      `json:"name,omitempty"`
	Bucket interface{} `json:"bucket,omitempty"`
}
type Platform struct {
	Device   *Device   `json:"device,omitempty"`
	Audience *Audience `json:"audience,omitempty"`
}
type CardPlatform struct {
	Platform *Platform `json:"platform,omitempty"`
}
type TweetExt struct {
	UnmentionInfo                  *RAndTTL `json:"unmentionInfo,omitempty"`
	SuperFollowMetadata            *RAndTTL `json:"superFollowMetadata,omitempty"`
	ReplyvotingDownvotePerspective *RAndTTL `json:"replyvotingDownvotePerspective,omitempty"`
}
type QuotedStatusPermalink struct {
	URL      string `json:"url,omitempty"`
	Expanded string `json:"expanded,omitempty"`
	Display  string `json:"display,omitempty"`
}
type SelfThread struct {
	ID    int64  `json:"id,omitempty"`
	IDStr string `json:"id_str,omitempty"`
}
type UserExt struct {
	HasNftAvatar        *RAndTTL `json:"hasNftAvatar,omitempty"`
	SuperFollowMetadata *RAndTTL `json:"superFollowMetadata,omitempty"`
	HighlightedLabel    *RAndTTL `json:"highlightedLabel,omitempty"`
}
type Moment struct{}
type OneCard struct{}
type Place struct{}
type OneMedia struct{}
type Broadcast struct{}
type Topic struct{}
type List struct{}
type SearchObjects struct {
	Tweets     map[string]*Tweet     `json:"tweets"`
	Users      map[string]*User      `json:"users"`
	Moments    map[string]*Moment    `json:"moments,omitempty"`
	Cards      map[string]*OneCard   `json:"cards,omitempty"`
	Places     map[string]*Place     `json:"places,omitempty"`
	Media      map[string]*OneMedia  `json:"media,omitempty"`
	Broadcasts map[string]*Broadcast `json:"broadcasts,omitempty"`
	Topics     map[string]*Topic     `json:"topics,omitempty"`
	Lists      map[string]*List      `json:"lists,omitempty"`
}
type Details struct {
	TimelinesDetails *TimelinesDetails `json:"timelinesDetails,omitempty"`
}
type SearchEntryTweetRef struct {
	ID          string `json:"id,omitempty"`
	DisplayType string `json:"displayType,omitempty"`
}
type SearchEntryItemContent struct {
	Tweet *SearchEntryTweetRef `json:"tweet,omitempty"`
}
type SearchEntryItem struct {
	Content         *SearchEntryItemContent `json:"content,omitempty"`
	ClientEventInfo *ClientEventInfo        `json:"clientEventInfo,omitempty"`
}
type SearchEntryContent struct {
	Item      *SearchEntryItem `json:"item,omitempty"`
	Operation *Operation       `json:"operation,omitempty"`
}
type SearchEntry struct {
	EntryID   string              `json:"entryId,omitempty"`
	SortIndex string              `json:"sortIndex,omitempty"`
	Content   *SearchEntryContent `json:"content,omitempty"`
}
type SearchAddEntries struct {
	Entries []*SearchEntry `json:"entries,omitempty"`
}
type Cursor struct {
	Value      string `json:"value,omitempty"`
	CursorType string `json:"cursorType,omitempty"`
}
type Operation struct {
	Cursor *Cursor `json:"cursor,omitempty"`
}
type OperationContent struct {
	Operation *Operation `json:"operation,omitempty"`
}
type SearchOperation struct {
	EntryID   string            `json:"entryId,omitempty"`
	SortIndex string            `json:"sortIndex,omitempty"`
	Content   *OperationContent `json:"content,omitempty"`
}
type SearchReplaceEntry struct {
	EntryIDToReplace string           `json:"entryIdToReplace,omitempty"`
	Entry            *SearchOperation `json:"entry,omitempty"`
}
type SearchInstructions struct {
	AddEntries   *SearchAddEntries   `json:"addEntries,omitempty"`
	ReplaceEntry *SearchReplaceEntry `json:"replaceEntry,omitempty"`
}
type SearchTimeline struct {
	ID           string                `json:"id,omitempty"`
	Instructions []*SearchInstructions `json:"instructions,omitempty"`
}
