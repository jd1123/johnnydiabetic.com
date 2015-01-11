package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jd1123/johnnydiabetic.com/helpers"
	"github.com/justinas/nosurf"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	context := helpers.GetContext(w, r, S)
	context["token"] = nosurf.Token(r)
	helpers.RunTemplateBase(w, "index.html", context)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	context := helpers.GetContext(w, r, S)
	helpers.RunTemplateBase(w, "about.html", context)
}

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	context := helpers.GetContext(w, r, S)
	context["token"] = nosurf.Token(r)
	session, err := S.Get(r, "session-name")
	if err != nil {
		log.Println("Session error", err)
	}
	fmt.Println("SESSION VALUES", session.Values)
	if r.Method == "POST" {
		user := r.FormValue("username")
		pw := r.FormValue("password")
		u, err := login(user, pw)
		if err != nil {
			w.Write([]byte("invalid username/password"))
		} else {
			session.Values["user_authenticated"] = true
			session.Values["user_name"] = u.UserId
			session.Values["user"] = u
			err = session.Save(r, w)
			if err != nil {
				log.Println("Session Error in LoginHandler()", err)
			}
			redir := "/"
			if session.Values["last3"] != nil {
				redir = session.Values["last3"].(string)
			}
			http.Redirect(w, r, redir, 301)
			//w.Write([]byte("login successful"))
		}
	} else if r.Method == "GET" {
		if session.Values["user_authenticated"] != nil {
			//msg := "Welcome back"
			http.Redirect(w, r, "/", 301)
		} else {
			helpers.RunTemplateBase(w, "login.html", context)
		}
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := S.Get(r, "session-name")
	if err != nil {
		fmt.Println("Session error", err)
		w.Write([]byte("Not logged in"))
		return
	}
	if session.Values["user_authenticated"] != nil {
		delete(session.Values, "user_authenticated")
		delete(session.Values, "user_name")
		delete(session.Values, "user")
		session.Save(r, w)
		http.Redirect(w, r, "/", 301)
	} else {
		w.Write([]byte("Not logged in"))
	}

}

// Implement 404 logic here
func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	helpers.RunTemplateBase(w, "404.html", nil)
}

func RobotsHandler(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("static/robots.txt")
	if err != nil {
		log.Printf("IO Error in RobotsHandler()")
		http.NotFound(w, r)
	} else {
		fileContents, err := ioutil.ReadAll(f)
		if err != nil {
			log.Printf("IO Error in RobotsHandler()")
			http.NotFound(w, r)
		} else {
			w.Write(fileContents)
		}
	}
}
