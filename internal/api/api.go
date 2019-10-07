package api

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/model"
	"net/http"
	"regexp"
	"strconv"
)

type Api struct {
	model                 *model.Model
	googleTranslateApiKey string
}

func NewApi(model *model.Model, googleTranslateApiKey string) *Api {
	return &Api{
		model:                 model,
		googleTranslateApiKey: googleTranslateApiKey,
	}
}

func setCORSHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
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
	if r.Method == "OPTIONS" {
		setCORSHeaders(w)
		return
	}

	if r.URL.Path == "/api/cards" {
		if r.Method == "GET" {
			api.HandleListCardsRequest(w, r)
		} else if r.Method == "POST" {
			api.HandleCreateCardRequest(w, r)
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
		} else if r.Method == "DELETE" {
			api.HandleDeleteCardRequest(w, r, cardId)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	} else if r.URL.Path == "/api/morphemes" {
		if r.Method == "GET" {
			api.HandleListMorphemesRequest(w, r)
		} else if r.Method == "POST" {
			api.HandleCreateMorphemeRequest(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	} else if match := regexp.MustCompile(
		`^/api/morphemes/(-?[0-9]+)$`).FindStringSubmatch(r.URL.Path); match != nil {
		morphemeId := MustAtoi(match[1])
		if r.Method == "GET" {
			api.HandleShowMorphemeRequest(w, r, morphemeId)
		} else if r.Method == "PUT" {
			api.HandleUpdateMorphemeRequest(w, r)
		} else if r.Method == "DELETE" {
			api.HandleDeleteMorphemeRequest(w, r, morphemeId)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	} else if r.URL.Path == "/api/translate" {
		if r.Method == "GET" {
			api.HandleTranslateRequest(w, r)
		} else if r.Method == "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	} else {
		http.Error(w, "Not found", http.StatusNotFound)
	}
}
