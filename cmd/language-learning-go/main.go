package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func main() {
	httpPort := os.Getenv("HTTP_PORT")
	httpsPort := os.Getenv("HTTPS_PORT")
	httpsCertPath := os.Getenv("HTTPS_CERT_PATH")
	httpsKeyPath := os.Getenv("HTTPS_KEY_PATH")
	dbPath := os.Getenv("DB_PATH")

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Error from sql.Open: %s", err)
	}
	api := &Api{db: db}
	handlerVars := InitHandlerVars(api)

	if httpPort != "" {
		if httpsPort != "" {
			go serveHttpRedirect(httpPort)
			serveHttps(handlerVars, httpsPort, httpsCertPath, httpsKeyPath)
		} else {
			serveHttp(handlerVars, httpPort)
		}
	} else {
		if httpsPort != "" {
			serveHttps(handlerVars, httpsPort, httpsCertPath, httpsKeyPath)
		} else {
			log.Fatal("Specify either HTTP_PORT or HTTPS_PORT env var")
		}
	}
}
