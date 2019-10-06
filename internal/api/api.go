package api

import (
	"database/sql"
	"log"
	"net/http"
	"regexp"
	"strconv"
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
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
}

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
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
		cardId := MustAtoi(match[1])
		if r.Method == "GET" {
			api.HandleShowCardRequest(w, r, cardId)
		} else if r.Method == "PUT" {
			api.HandleUpdateCardRequest(w, r, cardId)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	} else if r.URL.Path == "/api/morphemes" {
		if r.Method == "GET" {
			api.HandleListMorphemesRequest(w, r)
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
		morphemeId := MustAtoi(match[1])
		api.HandleShowMorphemeRequest(w, r, morphemeId)
	} else {
		http.Error(w, "Not found", http.StatusNotFound)
	}
}
