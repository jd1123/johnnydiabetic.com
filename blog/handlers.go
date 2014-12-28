package blog

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jd1123/johnnydiabetic.com/helpers"
	"github.com/justinas/nosurf"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	type mySt struct {
		Posts  []BlogPost
		String string
	}
	posts := GetAllPosts()
	Os := mySt{Posts: posts}
	helpers.RunTemplateBase(w, "blog/index.html", Os)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	context := helpers.GetContext(r, S)
	context["token"] = nosurf.Token(r)
	session, err := S.Get(r, "session-name")
	if err != nil {
		log.Println("Session error", err)
		// redirect
	}
	fmt.Println(context, session)
	if session.Values["user_authenticated"] != nil {
		// user is logged in
		if r.Method == "POST" {
			// post it!
		} else if r.Method == "GET" {
			// show the post page
			helpers.RunTemplateBase(w, "blog/post.html", context)
		}
	} else {
		//user is not logged in
		//redirect to login page
		w.Write([]byte("You are not logged in and got to the post page"))
	}
}
