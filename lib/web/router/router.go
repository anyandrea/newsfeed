package router

import (
	"net/http"

	"github.com/alexedwards/scs"
	"github.com/anyandrea/newsfeed/lib/database/newsfeeddb"
	"github.com/anyandrea/newsfeed/lib/web/html"
	"github.com/gorilla/mux"
)

func New(db newsfeeddb.NewsFeedDB, sm *scs.Manager) *mux.Router {
	router := mux.NewRouter()
	setupRoutes(db, sm, router)
	return router
}

func setupRoutes(db newsfeeddb.NewsFeedDB, sm *scs.Manager, router *mux.Router) *mux.Router {
	// HTML
	router.NotFoundHandler = http.HandlerFunc(html.NotFound)

	router.HandleFunc("/", html.Index(db, sm))
	router.HandleFunc("/error", html.ErrorHandler)
	router.HandleFunc("/fetch", html.FetchFeeds(db, sm))

	router.HandleFunc("/login", html.Login(db, sm))
	router.HandleFunc("/logout", html.Logout(db, sm))

	router.HandleFunc("/settings", html.Settings(db, sm))
	router.HandleFunc("/account", html.Account(db, sm))

	return router
}
