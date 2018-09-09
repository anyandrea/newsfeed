package html

import (
	"fmt"
	"net/http"

	"github.com/alexedwards/scs"
	"github.com/anyandrea/newsfeed/lib/database/newsfeeddb"
	"github.com/anyandrea/newsfeed/lib/web"
)

func Index(db newsfeeddb.NewsFeedDB, sm *scs.Manager) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		page := &Page{
			Title:  "Newsfeed",
			Active: "feeds",
		}

		// get user_id from session
		session := sm.Load(req)
		userId, err := session.GetInt("user_id")
		if err != nil {
			Error(rw, err)
			return
		}

		// if no actual user is logged in, then display admin account's feed subscriptions
		if userId < 1 {
			user, err := db.GetUserByEmail("admin@localhost")
			if err != nil {
				Error(rw, err)
				return
			}
			userId = user.Id
		}

		// get feeds
		feeds, err := db.GetFeedsByUserId(userId)
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

func Unauthorized(rw http.ResponseWriter) {
	page := &Page{
		Title: "Newsfeed - Unauthorized",
	}
	web.Render().HTML(rw, http.StatusUnauthorized, "unauthorized", page)
}

func FetchFeeds(db newsfeeddb.NewsFeedDB, sm *scs.Manager) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		// get user_id from session
		session := sm.Load(req)
		userId, err := session.GetInt("user_id")
		if err != nil {
			Error(rw, err)
			return
		}

		if userId < 1 {
			Unauthorized(rw)
			return
		}

		if err := db.FetchAllFeeds(); err != nil {
			Error(rw, err)
			return
		}
		Index(db, sm)(rw, req)
	}
}
