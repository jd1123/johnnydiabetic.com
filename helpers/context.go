package helpers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

func GetContext(w http.ResponseWriter, r *http.Request, S sessions.Store) map[string]interface{} {
	context := make(map[string]interface{})
	session, err := S.Get(r, "session-name")
	if err != nil {
		log.Println("Session error", err)
	}
	fmt.Println(session.Values)
	if session.Values["user_authenticated"] != nil {
		context["User"] = session.Values["user_name"]
	}
	if session.Values["last2"] != nil {
		session.Values["last3"] = session.Values["last2"]
	}
	if session.Values["last"] != nil {
		session.Values["last2"] = session.Values["last"]
	}
	session.Values["last"] = r.URL.Path
	err = session.Save(r, w)
	if err != nil {
		fmt.Println("Session error:", err)
	}
	// Pass in the request and
	// get some stuff from it

	return context
}
