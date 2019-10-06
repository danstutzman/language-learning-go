package api

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type ListMorphemesResponse struct {
	Morphemes []db.Morpheme `json:"morphemes"`
}

func (api *Api) HandleListMorphemesRequest(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	response := ListMorphemesResponse{
		Morphemes: db.SelectAllFromMorphemes(api.db),
	}
	bytes, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error from json.Marshal: %s", err)
	}
	w.Write(bytes)
}

func (api *Api) HandleShowMorphemeRequest(w http.ResponseWriter, r *http.Request,
	morphemeId string) {
	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	morpheme, err := db.FindMorphemeById(api.db, morphemeId)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		log.Fatalf("Internal Server Error: %s", err)
		return
	}

	bytes, err := json.Marshal(morpheme)
	if err != nil {
		log.Fatalf("Error from json.Marshal: %s", err)
	}
	w.Write(bytes)
}
