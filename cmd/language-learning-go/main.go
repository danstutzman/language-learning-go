package main

import (
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

	tlsCertPath := os.Getenv("TLS_CERT_PATH")
	tlsKeyPath := os.Getenv("TLS_KEY_PATH")

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
	if tlsCertPath != "" && tlsKeyPath != "" {
		log.Printf("Serving HTTP with TLS on port %s...", httpPort)
		err := server.ListenAndServeTLS(tlsCertPath, tlsKeyPath)
		if err != nil {
			log.Fatalf("Error from ListenAndServe: %s", err)
		}
	} else {
		log.Printf("Serving HTTP on port %s...", httpPort)
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Error from ListenAndServe: %s", err)
		}
	}
}
