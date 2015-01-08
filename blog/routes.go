package blog

/*

This is the blog package.
It will handle all aspects of the blog, it is akin to a Django app and
follows a similar design pattern.
It will store blog posts into a mongo db

*/

import "net/http"

type HandlerFunction func(w http.ResponseWriter, r *http.Request)

var appPrefix = "/blog"

var routes = map[string]HandlerFunction{
	"/":           IndexHandler,
	"/post/":      PostHandler,
	"/view/{key}": ViewPostHandler,
	"/edit/{key}": EditPostHandler,
}
