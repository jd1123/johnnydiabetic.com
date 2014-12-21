package main

import (
	"fmt"
	"net/http"

	"github.com/jd1123/johnnydiabetic.com/server"
)

func main() {
	fmt.Println("Listening on port 8081")
	server.RegisterHandlers()
	//http.Handle("/static", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":8081", nil)
}
