package main

import (
	"fmt"
	"github.com/nytimes/gziphandler"
	"github.com/shurcool/httpgzip"
	"log"
	"net/http"
	"os"
	"strings"
)

var withoutGz = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=\"utf-8\"")
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
})
var withGz = gziphandler.GzipHandler(withoutGz)

var fileServer = httpgzip.FileServer(
	http.Dir("web"),
	httpgzip.FileServerOptions{IndexHTML: true})

func main() {
	httpPort := os.Getenv("HTTP_PORT")
	httpsPort := os.Getenv("HTTPS_PORT")
	httpsCertPath := os.Getenv("HTTPS_CERT_PATH")
	httpsKeyPath := os.Getenv("HTTPS_KEY_PATH")

	if httpPort != "" {
		if httpsPort != "" {
			go serveHttpRedirect(httpPort)
			serveHttps(httpsPort, httpsCertPath, httpsKeyPath)
		} else {
			serveHttp(httpPort)
		}
	} else {
		if httpsPort != "" {
			serveHttps(httpsPort, httpsCertPath, httpsKeyPath)
		} else {
			log.Fatal("Specify either HTTP_PORT or HTTPS_PORT env var")
		}
	}
}

func serveHttp(httpPort string) {
	log.Printf("Serving HTTP on port %s...", httpPort)

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", handleRequest)

	server := &http.Server{
		Addr:    ":" + httpPort,
		Handler: serveMux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error from ListenAndServe: %s", err)
	}
}

func serveHttpRedirect(httpPort string) {
	log.Printf("Serving HTTP redirect on port %s...", httpPort)

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", handleRequestWithRedirect)

	server := &http.Server{
		Addr:    ":" + httpPort,
		Handler: serveMux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error from ListenAndServe: %s", err)
	}
}

func serveHttps(httpsPort string, certPath string, keyPath string) {
	log.Printf("Serving HTTP with TLS on port %s...", httpsPort)

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", handleRequest)

	server := &http.Server{
		Addr:    ":" + httpsPort,
		Handler: serveMux,
	}

	err := server.ListenAndServeTLS(certPath, keyPath)
	if err != nil {
		log.Fatalf("Error from ListenAndServeTLS: %s", err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	if strings.HasPrefix(r.URL.Path, "/api/") {
		withGz.ServeHTTP(w, r)
	} else {
		fileServer.ServeHTTP(w, r)
	}
}

func handleRequestWithRedirect(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	target := "https://" + r.Host + r.URL.Path
	if len(r.URL.RawQuery) > 0 {
		target += "?" + r.URL.RawQuery
	}
	http.Redirect(w, r, target, http.StatusTemporaryRedirect)
}
