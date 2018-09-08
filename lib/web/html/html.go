package html

import (
	"fmt"
	"net/http"

	"github.com/anyandrea/newsfeed/lib/database/newsfeeddb"
	"github.com/anyandrea/newsfeed/lib/web"
	"github.com/mmcdole/gofeed"
)

func Index(db newsfeeddb.NewsFeedDB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		page := &Page{
			Title:  "Newsfeed",
			Active: "feeds",
		}

		// get feeds
		var feeds []*gofeed.Feed
		parser := gofeed.NewParser()
		urls := []string{
			"https://www.iracing.com/category/news/sim-racing-news/feed/",
			"https://www.heise.de/newsticker/heise-atom.xml",
			"https://news.ycombinator.com/rss",
			"https://www.reddit.com/r/iracing.rss",
			"https://www.reddit.com/r/simracing.rss",
		}
		for _, url := range urls {
			feed, err := parser.ParseURL(url)
			// if err != nil {
			// 	Error(rw, err)
			// 	return
			// }
			if err == nil {
				feeds = append(feeds, feed)
			}
		}
		page.Content = feeds

		web.Render().HTML(rw, http.StatusOK, "index", page)
	}
}

func NotFound(rw http.ResponseWriter, req *http.Request) {
	page := &Page{
		Title: "Newsfeed - Not Found",
	}
	web.Render().HTML(rw, http.StatusNotFound, "not_found", page)
}

func ErrorHandler(rw http.ResponseWriter, req *http.Request) {
	Error(rw, fmt.Errorf("Internal Server Error"))
}
func Error(rw http.ResponseWriter, err error) {
	page := &Page{
		Title:   "Newsfeed - Error",
		Content: err,
	}
	web.Render().HTML(rw, http.StatusInternalServerError, "error", page)
}

func Settings(db newsfeeddb.NewsFeedDB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		page := &Page{
			Title:  "Newsfeed - Settings",
			Active: "settings",
		}
		web.Render().HTML(rw, http.StatusOK, "settings", page)
	}
}
