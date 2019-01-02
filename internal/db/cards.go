package db

import (
	"database/sql"
	"log"
)

type Card struct {
	CardId int    `json:"cardId"`
	EsText string `json:"esText"`
}

func AssertCardsHasCorrectSchema(db *sql.DB) {
	stmt, err := db.Prepare("select card_id, es_text from cards limit 1")
	if err != nil {
		log.Fatalf("Error from db.Prepare: %s", err)
	}
	defer stmt.Close()
}

func SelectAllFromCards(db *sql.DB) []Card {
	cards := []Card{}

	rows, err := db.Query("select card_id, es_text from cards")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var card Card
		err = rows.Scan(&card.CardId, &card.EsText)
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
