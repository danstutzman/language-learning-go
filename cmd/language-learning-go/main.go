package main

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/api"
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"bitbucket.org/danstutzman/language-learning-go/internal/model"
	"database/sql"
	"fmt"
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
	googleTranslateApiKey := os.Getenv("GOOGLE_TRANSLATE_API_KEY")
	languageModelPath := os.Getenv("LANGUAGE_MODEL_PATH")

	// Set mode=rw so it doesn't create database if file doesn't exist
	connString := fmt.Sprintf("file:%s?mode=rw", dbPath)
	dbConn, err := sql.Open("sqlite3", connString)
	if err != nil {
		log.Fatalf("Error from sql.Open: %s", err)
	}
	db.AssertCardsHasCorrectSchema(dbConn)
	db.AssertCardsMorphemesHasCorrectSchema(dbConn)
	db.AssertMorphemesHasCorrectSchema(dbConn)

	model := model.NewModel(dbConn, languageModelPath)
	api := api.NewApi(model, googleTranslateApiKey)
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
