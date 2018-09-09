package html

import (
	"net/http"

	"github.com/alexedwards/scs"
	"github.com/anyandrea/newsfeed/lib/database/newsfeeddb"
	"github.com/anyandrea/newsfeed/lib/web"
)

func Settings(db newsfeeddb.NewsFeedDB, sm *scs.Manager) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		page := &Page{
			Title:  "Newsfeed - Settings",
			Active: "settings",
		}

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

		// TODO: display settings page for logged-in user

		web.Render().HTML(rw, http.StatusOK, "settings", page)
	}
}
