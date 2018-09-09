package html

import (
	"net/http"

	"github.com/anyandrea/newsfeed/lib/database/newsfeeddb"
	"github.com/anyandrea/newsfeed/lib/web"
)

func Account(db newsfeeddb.NewsFeedDB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		page := &Page{
			Title:  "Newsfeed - Account",
			Active: "account",
		}
		web.Render().HTML(rw, http.StatusOK, "account", page)
	}
}
