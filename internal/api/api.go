package api

import (
	"database/sql"
	"log"
	"net/http"
)

type Api struct {
	db             *sql.DB
	dictionaryPath string
}

func NewApi(db *sql.DB, dictionaryPath string) *Api {
	return &Api{db: db, dictionaryPath: dictionaryPath}
}

func setCORSHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "X-Client-Version")
}

func (api *Api) HandleApiRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("X-Client-Version: %s", r.Header.Get("X-Client-Version"))

	if r.Method == "OPTIONS" {
		setCORSHeaders(w)
		return
	}

	if r.URL.Path == "/api/sync-cards" {
		if r.Method == "POST" {
			api.HandleSyncCardsRequest(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	} else if r.URL.Path == "/api/download-dictionary" {
		if r.Method == "GET" {
			api.HandleDownloadDictionaryRequest(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	} else {
		http.Error(w, "Not found", http.StatusNotFound)
	}
}
