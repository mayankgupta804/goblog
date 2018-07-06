package services

import (
	"log"
	"net/http"
)

// StartWebServer starts the web server on the provided port
func StartWebServer(port string) {

	r := NewRouter()
	http.Handle("/", r)

	log.Println("Starting web server on port: " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Println("An error occurred when starting HTTP server on port" + port)
		log.Fatalln("Error:", err.Error())
	}
}
