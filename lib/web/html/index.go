package html

import (
	"fmt"
	"net/http"

	"github.com/anyandrea/newsfeed/lib/database/newsfeeddb"
	"github.com/anyandrea/newsfeed/lib/web"
)

func Index(db newsfeeddb.NewsFeedDB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		page := &Page{
			Title:  "Newsfeed",
			Active: "feeds",
		}

		// TODO: get feed from currently logged in user
		// TODO: if no user is logged in, get feeds from admin user
		users, err := db.GetUsers()
		if err != nil {
			Error(rw, err)
			return
		}

		feeds, err := db.GetFeedsByUserId(users[0].Id)
		if err != nil {
			Error(rw, err)
			return
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

func FetchFeeds(db newsfeeddb.NewsFeedDB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		if err := db.FetchAllFeeds(); err != nil {
			Error(rw, err)
			return
		}
		Index(db)(rw, req)
	}
}
