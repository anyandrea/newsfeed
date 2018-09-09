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
		session := sm.Load(req)
		userName, _ := session.GetString("user_name")
		userId, _ := session.GetInt("user_id")

		// if no actual user is logged in, then display admin account's feed subscriptions
		if userId == 0 {
			user, err := db.GetUserByEmail("admin@localhost")
			if err != nil {
				Error(sm, rw, req, err)
				return
			}
			userId = user.Id
		}

		// get feeds
		feeds, err := db.GetFeedsByUserId(userId)
		if err != nil {
			Error(sm, rw, req, err)
			return
		}

		page := &Page{
			Title:   "Newsfeed",
			Active:  "feeds",
			User:    userName,
			Content: feeds,
		}
		web.Render().HTML(rw, http.StatusOK, "index", page)
	}
}

func NotFound(sm *scs.Manager) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		session := sm.Load(req)
		userName, _ := session.GetString("user_name")
		page := &Page{
			Title: "Newsfeed - Not Found",
			User:  userName,
		}
		web.Render().HTML(rw, http.StatusNotFound, "not_found", page)
	}
}

func ErrorHandler(sm *scs.Manager) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		Error(sm, rw, req, fmt.Errorf("Internal Server Error"))
	}
}
func Error(sm *scs.Manager, rw http.ResponseWriter, req *http.Request, err error) {
	session := sm.Load(req)
	userName, _ := session.GetString("user_name")
	page := &Page{
		Title:   "Newsfeed - Error",
		Content: err,
		User:    userName,
	}
	web.Render().HTML(rw, http.StatusInternalServerError, "error", page)
}

func FetchFeeds(db newsfeeddb.NewsFeedDB, sm *scs.Manager) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		session := sm.Load(req)
		userId, _ := session.GetInt("user_id")
		if userId == 0 {
			Unauthorized(rw)
			return
		}

		if err := db.FetchAllFeeds(); err != nil {
			Error(sm, rw, req, err)
			return
		}
		Index(db, sm)(rw, req)
	}
}
