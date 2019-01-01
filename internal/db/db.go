package db

import (
	"database/sql"
	"log"
)

type Card struct {
	CardId int    `json:"cardId"`
	Es     string `json:"es"`
}

type Exposure struct {
	CardId    int     `json:"cardId"`
	CreatedAt float64 `json:"createdAt"`
}

func AssertCardsHasCorrectSchema(db *sql.DB) {
	stmt, err := db.Prepare("select cardId, es from cards limit 1")
	if err != nil {
		log.Fatalf("Error from db.Prepare: %s", err)
	}
	defer stmt.Close()

	var card Card
	err = stmt.QueryRow().Scan(&card.CardId, &card.Es)
	if err != nil {
		log.Fatalf("Error from stmt.QueryRow: %s", err)
	}
}

func AssertExposuresHasCorrectSchema(db *sql.DB) {
	stmt, err := db.Prepare("select cardId, createdAt from exposures limit 1")
	if err != nil {
		log.Fatalf("Error from db.Prepare: %s", err)
	}
	defer stmt.Close()

	var exposure Exposure
	err = stmt.QueryRow().Scan(&exposure.CardId, &exposure.CreatedAt)
	if err != nil {
		log.Fatalf("Error from stmt.QueryRow: %s", err)
	}
}

func SelectAllFromCards(db *sql.DB) []Card {
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

func SelectAllFromExposures(db *sql.DB) []Exposure {
	exposures := []Exposure{}

	rows, err := db.Query("select cardId, createdAt from exposures")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var exposure Exposure
		err = rows.Scan(&exposure.CardId, &exposure.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		exposures = append(exposures, exposure)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return exposures
}

func InsertExposures(exposures []Exposure, db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Error from db.Begin: %s", err)
	}

	stmt, err := tx.Prepare(
		"insert into exposures(cardId, createdAt) values(?,?)")
	if err != nil {
		log.Fatalf("Error from tx.Prepare: %s", err)
	}
	defer stmt.Close()

	for _, exposure := range exposures {
		_, err = stmt.Exec(exposure.CardId, exposure.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()
}
