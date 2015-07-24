package main

import (
	"gotraining/13-http/api/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	log.Println("main : Started")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	go func() {
		log.Println("listener : Started : Listening on: http://localhost:" + port)
		http.ListenAndServe(":"+port, routes.API())
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	log.Println("main : Completed")
}
