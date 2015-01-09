package blog

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jd1123/johnnydiabetic.com/helpers"
	"github.com/justinas/nosurf"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	context := helpers.GetContext(r, S)
	posts := GetAllPosts()
	context["Posts"] = posts
	helpers.RunTemplateBase(w, "blog/index.html", context)
}

func ViewPostHandler(w http.ResponseWriter, r *http.Request) {
	context := helpers.GetContext(r, S)
	//session
	_, err := S.Get(r, "session-name")
	if err != nil {
		log.Println("Session error", err)
	}
	postId, err := strconv.Atoi(mux.Vars(r)["key"])
	if err != nil {
		w.Write([]byte("404!"))

	} else {
		// This sucks sooooo much. Do something less confusing.
		post := GetPostById(postId)
		if post != nil {
			PostMarkdown(post)
			context["Post"] = *post
			helpers.RunTemplateBase(w, "blog/blogpost.html", context)
		} else {
			w.Write([]byte("404!"))
		}
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	context := helpers.GetContext(r, S)
	context["token"] = nosurf.Token(r)
	session, err := S.Get(r, "session-name")
	if err != nil {
		log.Println("Session error", err)
		// redirect
	}
	if session.Values["user_authenticated"] != nil {
		// user is logged in
		if r.Method == "POST" {
			b := NewBlogPost(r.FormValue("title"), r.FormValue("body"))
			AddPost(b)
			w.Write([]byte("Post successful!"))
		} else if r.Method == "GET" {
			// show the post page
			helpers.RunTemplateBase(w, "blog/post.html", context)
		}
	} else {
		//user is not logged in
		//redirect to login page
		//w.Write([]byte("You are not logged in and got to the post page"))
		http.Redirect(w, r, "/login/", 301)
	}
}
