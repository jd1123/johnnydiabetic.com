package blog

import (
	"net/http"

	"github.com/jd1123/johnnydiabetic.com/helpers"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	type S struct {
		Posts  []BlogPost
		String string
	}
	posts := GetAllPosts()
	Os := S{Posts: posts}
	helpers.RunTemplate(w, "blog/index.html", "header.html", "footer.html", Os)
}
