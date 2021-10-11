package main

import (
	"log"
	"net/http"
)

func main() {

	srv := &http.Server{Addr: ":8080", Handler: http.HandlerFunc(handleWithPush)}
	log.Printf("Serving on https://0.0.0.0:8000")
	log.Fatal(srv.ListenAndServeTLS("../certs/localhost.crt", "../certs/localhost.key"))
}
func handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received connection: %s", r.Proto)
	w.Write([]byte("Hello"))
}

func handleWithPush(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received connection: %s", r.Proto)

	if r.URL.Path == "/2nd" {
		log.Println("Handling 2nd")
		w.Write([]byte("Hello Again!"))
		return
	}

	// Handle 1st request
	log.Println("Handling 1st request")

	pusher, ok := w.(http.Pusher)
	if !ok {
		log.Println("Client push failed.")
	} else {
		err := pusher.Push("/2nd", nil)
		if err != nil {
			log.Printf("Failed push: %v", err)
		}
	}

	// Send response body
	w.Write([]byte("Welcome to HTTP2"))
}
