package db

import (
	"database/sql"
	"log"
)

type CardState struct {
	CardId          int   `json:"cardId"`
	CreatedAtMillis int64 `json:"createdAtMillis"`
}

func AssertCardStatesHasCorrectSchema(db *sql.DB) {
	stmt, err := db.Prepare(
		"select card_id, created_at_millis from card_states limit 1")
	if err != nil {
		log.Fatalf("Error from db.Prepare: %s", err)
	}
	defer stmt.Close()
}

func SelectAllFromCardStates(db *sql.DB) []CardState {
	cardStates := []CardState{}

	rows, err := db.Query("select card_id, created_at_millis from card_states")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var cardState CardState
		err = rows.Scan(&cardState.CardId, &cardState.CreatedAtMillis)
		if err != nil {
			log.Fatal(err)
		}
		cardStates = append(cardStates, cardState)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return cardStates
}

func InsertCardStates(cardStates []CardState, db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Error from db.Begin: %s", err)
	}

	stmt, err := tx.Prepare(
		"insert into card_states(card_id, created_at_millis) values(?,?)")
	if err != nil {
		log.Fatalf("Error from tx.Prepare: %s", err)
	}
	defer stmt.Close()

	for _, cardState := range cardStates {
		_, err = stmt.Exec(cardState.CardId, cardState.CreatedAtMillis)
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()
}
