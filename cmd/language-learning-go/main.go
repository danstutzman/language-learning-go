package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		log.Fatal("HTTP_PORT can not be blank")
	}

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

		path := r.URL.Path
		if path == "/" {
			http.ServeFile(w, r, "web/index.html")
		} else if path == "/index.html" {
			http.ServeFile(w, r, "web/index.html")
		} else if path == "/favicon.ico" {
			http.ServeFile(w, r, "web/favicon.ico")
		} else if path == "/service-worker.js" {
			http.ServeFile(w, r, "web/service-worker.js")
		} else if strings.HasPrefix(path, "/static/") {
			http.ServeFile(w, r, "web"+path)
		} else {
			http.NotFound(w, r)
		}
	})

	server := &http.Server{
		Addr:    ":" + httpPort,
		Handler: serveMux,
	}
	log.Fatal(server.ListenAndServe())
}
