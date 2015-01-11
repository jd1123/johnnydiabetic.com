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

func EditPostHandler(w http.ResponseWriter, r *http.Request) {
	context := helpers.GetContext(r, S)
	context["token"] = nosurf.Token(r)
	context["Edit"] = true
	session, err := S.Get(r, "session-name")
	if err != nil {
		log.Println("Session error in EditPostHandler()", err)
	}
	if session.Values["user_authenticated"] != nil { // User is authenticated
		if r.Method == "POST" { // User is posting an edit
			postId, err := strconv.Atoi(mux.Vars(r)["key"])
			if err != nil {
				w.Write([]byte("key error!"))
			}
			err = EditPost(postId, r.FormValue("title"), r.FormValue("body"))
			if err != nil {
				w.Write([]byte("Post error!"))

			} else {
				http.Redirect(w, r, "/blog/view/"+strconv.Itoa(postId)+"/", 301)
			}
		} else if r.Method == "GET" { // User is trying to edit a post
			postId, err := strconv.Atoi(mux.Vars(r)["key"])
			if err != nil {
				w.Write([]byte("Key Error!"))
			} else {
				post := GetPostById(postId)
				context["Post"] = post
				context["Key"] = postId
				helpers.RunTemplateBase(w, "blog/post.html", context)
			}
		}
	} else { // User is not authenticated
		http.Redirect(w, r, "/login/", 301)
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
			http.Redirect(w, r, "/blog/", 301)
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
