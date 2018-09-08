package router

import (
	"net/http"

	"github.com/anyandrea/newsfeed/lib/database/newsfeeddb"
	"github.com/anyandrea/newsfeed/lib/web/html"
	"github.com/gorilla/mux"
)

func New(db newsfeeddb.NewsFeedDB) *mux.Router {
	router := mux.NewRouter()
	setupRoutes(db, router)
	return router
}

func setupRoutes(db newsfeeddb.NewsFeedDB, router *mux.Router) *mux.Router {
	// HTML
	router.NotFoundHandler = http.HandlerFunc(html.NotFound)

	router.HandleFunc("/", html.Index(db))
	router.HandleFunc("/error", html.ErrorHandler)

	router.HandleFunc("/settings", html.Settings(db))

	return router
}
