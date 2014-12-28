package blog

import (
	"log"
	"net/http"
	"path"
	"sort"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jd1123/johnnydiabetic.com/middleware"
	"github.com/justinas/nosurf"
	"gopkg.in/mgo.v2"
)

var S sessions.Store

// App logic, stuff that doesnt fit into the framework files
func AddPost(b BlogPost) error {
	s, err := mgo.Dial("localhost")
	if err != nil {
		log.Println("Error:", err)
		return err
	}
	c := s.DB("test").C("blogposts")
	err = c.Insert(b)
	if err != nil {
		log.Println("Error:", err)
		return err
	}
	return nil
}

func GetAllPosts() []BlogPost {
	s, err := mgo.Dial("localhost")
	if err != nil {
		log.Println("Error:", err)
		return nil
	}
	c := s.DB("test").C("blogposts")
	var results []BlogPost
	err = c.Find(nil).All(&results)
	if err != nil {
		log.Println("Error:", err)
		return nil
	}
	sort.Sort(BlogPostCollection(results))
	return results
}

func RegisterHandlers(r *mux.Router, s sessions.Store) {
	if s != nil {
		S = s
	}
	for k, v := range routes {
		fullPath := path.Join(appPrefix, k) + "/"
		log.Println("registering", fullPath)
		r.HandleFunc(fullPath, v)
		http.Handle(fullPath, nosurf.New(middleware.LogRequest(r)))
	}
}
