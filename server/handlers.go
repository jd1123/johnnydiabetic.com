package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jd1123/johnnydiabetic.com/helpers"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := S.Get(r, "session-name")
	session.Values["test"] = "saved!"
	err := session.Save(r, w)
	if err != nil {
		fmt.Println("Session save errror:", err)
	}
	helpers.RunTemplate(w, "index.html", "header.html", "footer.html", nil)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := S.Get(r, "session-name")
	fmt.Println(session.Values["foo"])
	helpers.RunTemplate(w, "about.html", "header.html", "footer.html", nil)
}

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	session, err := S.Get(r, "session-name")
	if err != nil {
		log.Println("Session error")
	}
	fmt.Println("SESSION VALUES", session.Values)
	if r.Method == "POST" {
		user := r.FormValue("username")
		pw := r.FormValue("password")
		u, err := login(user, pw)
		fmt.Println(u)
		if err != nil {
			w.Write([]byte("invalid username/password"))
		} else {
			session.Values["user_authenticated"] = true
			session.Values["user"] = u
			fmt.Println("Session Values:", session.Values)
			err = session.Save(r, w)
			if err != nil {
				fmt.Println(err)
			}
			w.Write([]byte("login successful"))
		}
		fmt.Println("trying to save session...")
	} else if r.Method == "GET" {
		if session.Values["user_authenticated"] == true {
			msg := "Welcome back"
			w.Write([]byte(msg))
		} else {
			helpers.RunTemplate(w, "login.html", "header.html", "footer.html", nil)
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
		delete(session.Values, "user")
		session.Save(r, w)
		w.Write([]byte("logged out"))
	} else {
		w.Write([]byte("Not logged in"))
	}

}

// Implement 404 logic here
func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	helpers.RunTemplate(w, "404.html", "header.html", "footer.html", nil)
}

func RobotsHandler(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("static/robots.txt")
	if err != nil {
		log.Printf("IO Error")
		w.Write([]byte("IO Error!"))
	} else {
		r, err := ioutil.ReadAll(f)
		if err != nil {
			log.Printf("IO Error")
			w.Write([]byte("IO Error!"))
		} else {
			w.Write(r)
		}
	}
}
