package html

import (
	"net/http"

	"github.com/anyandrea/newsfeed/lib/database/newsfeeddb"
	"github.com/anyandrea/newsfeed/lib/web"
)

func Settings(db newsfeeddb.NewsFeedDB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		page := &Page{
			Title:  "Newsfeed - Settings",
			Active: "settings",
		}
		web.Render().HTML(rw, http.StatusOK, "settings", page)
	}
}
