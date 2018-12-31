package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		log.Fatal("HTTP_PORT can not be blank")
	}

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	server := &http.Server{
		Addr:    ":" + httpPort,
		Handler: serveMux,
	}
	log.Fatal(server.ListenAndServe())
}
