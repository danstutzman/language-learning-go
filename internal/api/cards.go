package api

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type CardList struct {
	Cards []db.Card `json:"cards"`
}

func (api *Api) HandleListCardsRequest(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	response := CardList{
		Cards: db.SelectAllFromCards(api.db),
	}
	bytes, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error from json.Marshal: %s", err)
	}
	w.Write(bytes)
}

func (api *Api) HandleShowCardRequest(w http.ResponseWriter, r *http.Request,
	cardId string) {
	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	card, err := db.FindCardById(api.db, cardId)
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

	bytes, err := json.Marshal(card)
	if err != nil {
		log.Fatalf("Error from json.Marshal: %s", err)
	}
	w.Write(bytes)
}
