package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jd1123/johnnydiabetic.com/server"
)

func main() {
	fmt.Println("Listening on port 8081")
	server.RegisterHandlers()
	//http.Handle("/static", http.FileServer(http.Dir("static")))
	log.Fatal(http.ListenAndServe(":8081", nil))
}
