package html

import (
	"net/http"

	"github.com/alexedwards/scs"
	"github.com/anyandrea/newsfeed/lib/database/newsfeeddb"
	"github.com/anyandrea/newsfeed/lib/web"
)

func Settings(db newsfeeddb.NewsFeedDB, sm *scs.Manager) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		session := sm.Load(req)
		userId, _ := session.GetInt("user_id")
		if userId == 0 {
			Unauthorized(rw)
			return
		}

		// TODO: display settings page for logged-in user
		userName, _ := session.GetString("user_name")
		page := &Page{
			Title:  "Newsfeed - Settings",
			Active: "settings",
			User:   userName,
		}
		web.Render().HTML(rw, http.StatusOK, "settings", page)
	}
}
