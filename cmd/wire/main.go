package main

import (
	"dependencyInjection/remoteService/wire"
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	container := wire.CreateContainer(1)
	http.HandleFunc("/service1", func(w http.ResponseWriter, r *http.Request) {
		container.Remoteservice.Sleep()
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/service2", func(w http.ResponseWriter, r *http.Request) {
		container.Remoteservice2.Sleep()
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}