package html

import (
	"net/http"

	"github.com/alexedwards/scs"
	"github.com/anyandrea/newsfeed/lib/database/newsfeeddb"
	"github.com/anyandrea/newsfeed/lib/util"
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

		// get form login data
		if err := req.ParseForm(); err != nil {
			Error(sm, rw, req, err)
			return
		}
		email := req.FormValue("email")
		password := req.FormValue("password")

		// try to login user
		if len(email) > 0 && len(password) > 0 {
			user, err := db.GetUserByEmail(email)
			if err != nil {
				Error(sm, rw, req, err)
				return
			}

			if user.Id != 0 && util.ComparePasswords(user.Password, password) { // correct password?
				session.PutInt(rw, "user_id", user.Id) // store session
				session.PutString(rw, "user_name", user.Name)
			}
		}

		// get user_id from session
		userId, _ := session.GetInt("user_id")
		if userId != 0 { // valid session, redirect to account page
			http.Redirect(rw, req, "/account", http.StatusFound)
			return
		}

		page := &Page{
			Title: "Newsfeed - Login",
		}
		web.Render().HTML(rw, http.StatusOK, "login", page)
	}
}

func Logout(db newsfeeddb.NewsFeedDB, sm *scs.Manager) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		session := sm.Load(req)
		if err := session.Destroy(rw); err != nil {
			Error(sm, rw, req, err)
			return
		}
		http.Redirect(rw, req, "/login", http.StatusFound)
	}
}
