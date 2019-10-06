package api

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ListMorphemesResponse struct {
	Morphemes []db.MorphemeRow `json:"morphemes"`
}

func (api *Api) HandleListMorphemesRequest(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	response := ListMorphemesResponse{
		Morphemes: db.FromMorphemes(api.db, "limit 20"),
	}
	bytes, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error from json.Marshal: %s", err)
	}
	w.Write(bytes)
}

func (api *Api) HandleShowMorphemeRequest(w http.ResponseWriter, r *http.Request,
	morphemeId int) {
	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	where := fmt.Sprintf("WHERE id = %d", morphemeId)
	morphemes := db.FromMorphemes(api.db, where)
	if len(morphemes) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
		return
	} else if len(morphemes) > 1 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		panic("Too many morphemes")
	}
	morpheme := morphemes[0]

	bytes, err := json.Marshal(morpheme)
	if err != nil {
		panic(err)
	}
	w.Write(bytes)
}
