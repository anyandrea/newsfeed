package html

import (
	"net/http"

	"github.com/alexedwards/scs"
	"github.com/anyandrea/newsfeed/lib/database/newsfeeddb"
	"github.com/anyandrea/newsfeed/lib/web"
)

func Unauthorized(rw http.ResponseWriter) {
	page := &Page{
		Title: "Newsfeed - Unauthorized",
	}
	web.Render().HTML(rw, http.StatusUnauthorized, "unauthorized", page)
}

func Login(db newsfeeddb.NewsFeedDB, sm *scs.Manager) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		session := sm.Load(req)

		page := &Page{
			Title: "Newsfeed - Login",
		}

		// get login data
		if err := req.ParseForm(); err != nil {
			Error(rw, err)
			return
		}
		email := req.FormValue("email")
		password := req.FormValue("password")

		// try to login user
		if len(email) > 0 && len(password) > 0 {
			user, err := db.GetUserByEmail(email)
			if err != nil {
				Error(rw, err)
				return
			}
			if user.Id != 0 &&
				user.Password == password {
				if err := session.PutInt(rw, "user_id", user.Id); err != nil {
					Error(rw, err)
					return
				}
			}
		}

		// get user_id from session
		userId, err := session.GetInt("user_id")
		if err != nil {
			Error(rw, err)
			return
		}

		if userId != 0 {
			http.Redirect(rw, req, "/account", http.StatusFound)
			return
		}

		web.Render().HTML(rw, http.StatusOK, "login", page)
	}
}

func Logout(db newsfeeddb.NewsFeedDB, sm *scs.Manager) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		session := sm.Load(req)
		if err := session.Destroy(rw); err != nil {
			Error(rw, err)
			return
		}
		http.Redirect(rw, req, "/login", http.StatusFound)
	}
}
