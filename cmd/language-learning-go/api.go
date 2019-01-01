package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Api struct {
	db *sql.DB
}

type Card struct {
	CardId int    `json:"cardId"`
	Es     string `json:"es"`
}

type CardsResponse struct {
	Cards []Card `json:"cards"`
}

func selectAllFromCards(db *sql.DB) []Card {
	cards := []Card{}

	rows, err := db.Query("select cardId, es from cards")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var card Card
		err = rows.Scan(&card.CardId, &card.Es)
		if err != nil {
			log.Fatal(err)
		}
		cards = append(cards, card)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return cards
}

func (api *Api) handleApiRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")
	if r.URL.Path == "/api/cards.json" {
		response := CardsResponse{Cards: selectAllFromCards(api.db)}
		bytes, err := json.Marshal(response)
		if err != nil {
			log.Fatalf("Error from json.Marshal: %s", err)
		}
		w.Write(bytes)

	} else {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	}
}
