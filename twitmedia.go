package main

import (
	"context"
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"net/http"
	"time"
	"twitmedia/twitter"
)

var (
	listen   = flag.String("addr", ":8080", "Address/port to listen on")
	cacheDir = flag.String("cache-dir", "/tmp/twitmedia-cache", "Media cache directory")
)

var (
	lineDelimiter = []byte("\n")
	errorMsg      = []byte("ERR\n")
	//go:embed public
	content embed.FS
)

var (
	api *twitter.Api
)

func getAuth(r *http.Request) twitter.TAuth {
	const prefix = "x-twit-cred-"
	return twitter.TAuth{
		Token:  r.Header.Get(prefix + "token"),
		Ct0:    r.Header.Get(prefix + "ct0"),
		Bearer: r.Header.Get(prefix + "bearer"),
	}
}

func timeline(w http.ResponseWriter, request *http.Request) {
	qs := request.URL.Query()
	cursor := qs.Get("cursor")
	var dateFrom, dateTo *time.Time
	dateFrom = &time.Time{}
	*dateFrom = time.Now().AddDate(0, 0, -7)
	if date, err := time.Parse("2006-01-02", qs.Get("from")); err == nil {
		*dateFrom = date
	}
	dateTo = nil
	if date, err := time.Parse("2006-01-02", qs.Get("to")); err == nil {
		dateTo = &time.Time{}
		*dateTo = date
	}
	auth := getAuth(request)
	w.Header().Set("X-Content-Type-Options", "nosniff")
	ctx, cancel := context.WithTimeout(request.Context(), 20*time.Second)
	defer cancel()

	prog := make(chan twitter.Progress, 1)

	var resp *twitter.TwitResponse = nil
	go func() {
		var err error
		//resp, err = api.GetHomeMedia(ctx, cursor, getAuth(request), prog)
		resp, err = api.GetSearchMedia(ctx, dateFrom, dateTo, cursor, auth, prog)
		if err != nil {
			fmt.Errorf("%s", err)
			w.Write(errorMsg)
		}
		close(prog)
	}()

	for p := range prog {
		if dat, err := json.Marshal(p); err == nil {
			w.Write(dat)
			w.Write(lineDelimiter)
			w.(http.Flusher).Flush()
		}
	}

	if resp == nil {
		return
	}

	dat, err := json.Marshal(resp)
	if err != nil {
		fmt.Errorf("%s", err)
		w.Write(errorMsg)
		return
	}
	w.Write(dat)
	w.Write(lineDelimiter)
}

func main() {
	flag.Parse()

	api = twitter.NewApi(*cacheDir)
	api.Start(5)

	http.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir(*cacheDir))))
	http.Handle("/tl", http.HandlerFunc(timeline))
	html, _ := fs.Sub(fs.FS(content), "public")
	http.Handle("/", http.FileServer(http.FS(html)))
	http.ListenAndServe(*listen, nil)
}
