package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type Api struct {
	db *sql.DB
}

func NewApi(db *sql.DB) *Api {
	return &Api{db: db}
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
	} else {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	}
}
