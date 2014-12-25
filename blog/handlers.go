package blog

import (
	"log"
	"net/http"

	"github.com/jd1123/johnnydiabetic.com/helpers"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	type S struct {
		Posts  []BlogPost
		String string
	}
	BlogPosts := make([]BlogPost, 0)
	BlogPosts = append(BlogPosts, NewBlogPost("My First Post", "Merry Christmas-Eve!"))
	Os := S{Posts: BlogPosts, String: "This is a test"}
	helpers.RunTemplate(w, "blog/index.html", "header.html", "footer.html", Os)
	log.Println("IndexHandler in Blog app done.")
}
