package main

import (
	"log"
	"net/http"
)

type application struct {
}

func main() {
	// set up an app config
	app := application{}

	// get application routes
	mux := app.routes()

	// print out a message
	log.Println("Starting server on port 8080..")

	// start the server
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
