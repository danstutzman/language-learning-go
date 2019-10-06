package api

import (
	"database/sql"
	"log"
	"net/http"
	"regexp"
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

	if r.URL.Path == "/api/cards" {
		if r.Method == "GET" {
			api.HandleListCardsRequest(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	} else if match := regexp.MustCompile(
		`^/api/cards/(-?[0-9]+)$`).FindStringSubmatch(r.URL.Path); match != nil {
		api.HandleShowCardRequest(w, r, match[1])
	} else if r.URL.Path == "/api/morphemes" {
		if r.Method == "GET" {
			api.HandleListMorphemesRequest(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	} else if r.URL.Path == "/api/sync-cards" {
		if r.Method == "POST" {
			api.HandleSyncRequest(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	} else if r.URL.Path == "/api/download-dictionary" {
		if r.Method == "GET" {
			api.HandleDownloadDictionaryRequest(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	} else if match := regexp.MustCompile(
		`^/api/morphemes/(-?[0-9]+)$`).FindStringSubmatch(r.URL.Path); match != nil {
		api.HandleShowMorphemeRequest(w, r, match[1])
	} else {
		http.Error(w, "Not found", http.StatusNotFound)
	}
}
