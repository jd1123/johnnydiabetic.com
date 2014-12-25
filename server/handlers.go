package server

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jd1123/johnnydiabetic.com/helpers"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	helpers.RunTemplate(w, "index.html", "header.html", "footer.html", nil)
	//w.Write([]byte("This is not working at all"))
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	helpers.RunTemplate(w, "about.html", "header.html", "footer.html", nil)
}

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
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
