package helpers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

func GetContext(r *http.Request, S sessions.Store) map[string]interface{} {
	context := make(map[string]interface{})
	session, err := S.Get(r, "session-name")
	if err != nil {
		log.Println("Session error", err)
	}
	fmt.Println(session.Values)
	if session.Values["user_authenticated"] != nil {
		context["User"] = session.Values["user_name"]
	}

	// Pass in the request and
	// get some stuff from it

	return context
}
