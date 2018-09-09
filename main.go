package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/alexedwards/scs"
	"github.com/anyandrea/newsfeed/lib/database"
	"github.com/anyandrea/newsfeed/lib/database/newsfeeddb"
	"github.com/anyandrea/newsfeed/lib/env"
	"github.com/anyandrea/newsfeed/lib/web/router"
	"github.com/urfave/negroni"
)

func main() {
	db := setupDatabase()
	sm := setupSessionManager()

	// setup SIGINT catcher for graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// start a http server with negroni
	server := startHTTPServer(db, sm)

	// wait for SIGINT
	<-stop
	log.Println("Shutting down server...")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	server.Shutdown(ctx)
	log.Println("Server gracefully stopped")
}

func setupSessionManager() *scs.Manager {
	sm := scs.NewCookieManager("newsfeed-super-secret-cookie-key")
	sm.Lifetime(2 * time.Hour)
	sm.Persist(true)
	//sm.HttpOnly(false)
	//sm.Secure(true) // needs HTTPS
	return sm
}

func setupDatabase() newsfeeddb.NewsFeedDB {
	// setup weather database
	adapter := database.NewAdapter()
	if err := adapter.RunMigrations("lib/database/migrations"); err != nil {
		if !strings.Contains(err.Error(), "no change") {
			log.Println("Could not run database migrations")
			log.Fatal(err)
		}
	}
	db := newsfeeddb.NewNewsFeedDB(adapter)

	// background jobs
	go housekeeping(db)
	go feedCollection(db)

	return db
}

func housekeeping(db newsfeeddb.NewsFeedDB) {
	for {
		// retention policy of maximum 50 entries per feed
		if err := db.Housekeeping(50); err != nil {
			log.Println("Feed housekeeping failed")
			log.Fatal(err)
		}
		time.Sleep(12 * time.Hour)
	}
}

func feedCollection(db newsfeeddb.NewsFeedDB) {
	for {
		// go through all subscribed feeds and update them
		if err := db.FetchAllFeeds(); err != nil {
			log.Println("Feed collection failed")
			log.Fatal(err)
		}
		time.Sleep(1 * time.Hour)
	}
}

func startHTTPServer(db newsfeeddb.NewsFeedDB, sm *scs.Manager) *http.Server {
	handler := negroni.Classic()
	handler.UseHandler(sm.Use(router.New(db, sm)))

	addr := ":" + env.Get("PORT", "8080")
	server := &http.Server{Addr: addr, Handler: handler}

	go func() {
		log.Printf("Listening on http://0.0.0.0%s\n", addr)
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	return server
}
