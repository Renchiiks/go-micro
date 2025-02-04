package main

import (
	"log"
	"net/http"
)

const webPort = 8080

type Config struct{}

func main() {
	// Start the broker service
	app := Config{}
	log.Println("Starting broker service on port", webPort)

	// define http server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: app.routes(),
	}

	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
