package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	// TODO: Make it possible to disable.
	sigTermReceived = false
)

// Print a log line every second once we receive a termination signal.
func sigTermHeartbeater() {
	for {
		log.Println("Post-SIGTERM heartbeat")
		time.Sleep(time.Second)
	}
}

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

		// Start sending heartbeat logs after termination signal is received
		// so we can see how long it takes the task to be killed.
		// No, this fake semaphore is not really safe.
		if s == syscall.SIGTERM && !sigTermReceived {
			sigTermReceived = true
			go sigTermHeartbeater()
		}
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
