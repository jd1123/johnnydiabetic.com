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
	session, _ := Store.Get(r, "session-name")
	session.Values["user"] = "bar"
	err := session.Save(r, w)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(session)
	helpers.RunTemplate(w, "index.html", "header.html", "footer.html", nil)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "session-name")
	fmt.Println(session.Values["foo"])
	helpers.RunTemplate(w, "about.html", "header.html", "footer.html", nil)
}

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//fmt.Println(r.FormValue("username"))
		//fmt.Println(r.FormValue("password"))
		w.Write([]byte("to be implemented"))
	} else if r.Method == "GET" {
		helpers.RunTemplate(w, "login.html", "header.html", "footer.html", nil)
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
