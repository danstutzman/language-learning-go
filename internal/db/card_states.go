package db

import (
	"database/sql"
	"log"
)

type CardState struct {
	CardId    int    `json:"cardId"`
	StateJson string `json:"stateJson"`
}

func AssertCardStatesHasCorrectSchema(db *sql.DB) {
	stmt, err := db.Prepare(
		"select card_id, state_json from card_states limit 1")
	if err != nil {
		log.Fatalf("Error from db.Prepare: %s", err)
	}
	defer stmt.Close()
}

func SelectAllFromCardStates(db *sql.DB) []CardState {
	cardStates := []CardState{}

	rows, err := db.Query("select card_id, state_json from card_states")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var cardState CardState
		err = rows.Scan(&cardState.CardId, &cardState.StateJson)
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

func UpdateCardStates(cardStates []CardState, db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Error from db.Begin: %s", err)
	}

	stmt, err := tx.Prepare(
		"update card_states set state_json=? where card_id=?")
	if err != nil {
		log.Fatalf("Error from tx.Prepare: %s", err)
	}
	defer stmt.Close()

	for _, cardState := range cardStates {
		_, err = stmt.Exec(cardState.StateJson, cardState.CardId)
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()
}
