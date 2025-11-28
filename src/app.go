package main

import (
	"app/internal"
	"log"
	"os"
	"os/signal"
)

func main() {
	log.Println("Starting server")
	srv, err := internal.StartServer()
	if err != nil {
		log.Println("Error in starting server", err)
		return
	}
	log.Println("Server started")
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	log.Println("Closing server")
	srv.Close()
	log.Println("Server closed")
}
