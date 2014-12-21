package server

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jd1123/johnnydiabetic.com/helpers"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	helpers.RunTemplate(w, helpers.Template("index.html"), helpers.Template("base.html"), nil)
	//w.Write([]byte("This is not working at all"))
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	helpers.RunTemplate(w, helpers.Template("about.html"), helpers.Template("base.html"), nil)
}

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func RobotsHandler(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("static/robots.txt")
	if err != nil {
		log.Printf("IO Error")
	} else {
		r, err := ioutil.ReadAll(f)
		if err != nil {
		}
		w.Write(r)
	}
}
