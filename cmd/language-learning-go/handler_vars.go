package main

import (
	"github.com/nytimes/gziphandler"
	"github.com/shurcool/httpgzip"
	"log"
	"net/http"
	"strings"
)

type HandlerVars struct {
	api        *Api
	withGz     http.Handler
	fileServer http.Handler
}

func InitHandlerVars(api *Api) *HandlerVars {
	withoutGz := http.HandlerFunc(api.handleApiRequest)
	return &HandlerVars{
		api:    api,
		withGz: gziphandler.GzipHandler(withoutGz),
		fileServer: httpgzip.FileServer(
			http.Dir("web"),
			httpgzip.FileServerOptions{IndexHTML: true}),
	}
}

func (vars *HandlerVars) handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	if strings.HasPrefix(r.URL.Path, "/api/") {
		vars.withGz.ServeHTTP(w, r)
	} else {
		vars.fileServer.ServeHTTP(w, r)
	}
}

func serveHttp(handlerVars *HandlerVars, httpPort string) {
	log.Printf("Serving HTTP on port %s...", httpPort)

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", handlerVars.handleRequest)

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

func serveHttps(handlerVars *HandlerVars, httpsPort string, certPath string,
	keyPath string) {
	log.Printf("Serving HTTP with TLS on port %s...", httpsPort)

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", handlerVars.handleRequest)

	server := &http.Server{
		Addr:    ":" + httpsPort,
		Handler: serveMux,
	}

	err := server.ListenAndServeTLS(certPath, keyPath)
	if err != nil {
		log.Fatalf("Error from ListenAndServeTLS: %s", err)
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
