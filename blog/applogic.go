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
	"github.com/russross/blackfriday"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

func EditPost(id int) error {
	return nil
}

func GetPostById(id int) *BlogPost {
	s, err := mgo.Dial("localhost")
	if err != nil {
		log.Println("Error:", err)
		return nil
	}
	c := s.DB("test").C("blogposts")
	var result []BlogPost
	err = c.Find(bson.M{"id": id}).All(&result)
	if err != nil {
		log.Println("DB Error in GetPostById()", err)
		return nil
	}
	if len(result) > 0 {
		return &result[0]
	} else {
		return nil
	}
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
	for i := range results {
		PostMarkdown(&results[i])
	}
	return results
}

func PostMarkdown(b *BlogPost) {
	b.Content = string(blackfriday.MarkdownBasic([]byte(b.Content)))
}

func RegisterHandlers(r *mux.Router, s sessions.Store) {
	S = s
	for k, v := range routes {
		fullPath := path.Join(appPrefix, k) + "/"
		log.Println("registering", fullPath)
		r.HandleFunc(fullPath, v)
		http.Handle(fullPath, nosurf.New(middleware.LogRequest(r)))
	}
}
