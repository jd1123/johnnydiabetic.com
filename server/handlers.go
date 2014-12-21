package server

import (
	"net/http"

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
