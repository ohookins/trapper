package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func printEnvironment() {
	log.Println("---- ENVIRONMENT ----")
	for _, value := range os.Environ() {
		log.Println(value)
	}
	log.Println("---- END ENVIRONMENT ----")
}

func signalHandler(sigChan chan os.Signal) {
	for s := range sigChan {
		log.Println("Received signal ", s)
	}
}

func setupHttpHandlers() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request on /")
		fmt.Fprint(w, "OK")
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request on /healthz")
		fmt.Fprint(w, "OK")
	})
}

func main() {
	// Handle all signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan)
	go signalHandler(sigChan)

	printEnvironment()

	setupHttpHandlers()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
