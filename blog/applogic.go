package blog

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

// App logic, stuff that doesnt fit into the framework files

func AddPost() {
	s, err := mgo.Dial("localhost")
	if err != nil {
		log.Println("Error:")
		log.Fatal(err)
	}
	d := s.DB("test")
	fmt.Println(d)
}

func RegisterHandlers(r *mux.Router) {
	for k, v := range routes {
		fullPath := path.Join(appPrefix, k) + "/"
		log.Println("registering", fullPath)
		r.HandleFunc(fullPath, v)
		http.Handle(fullPath, r)
	}
}
