package db

import (
	"database/sql"
	"log"
)

type Exposure struct {
	CardId          int   `json:"cardId"`
	CreatedAtMillis int64 `json:"createdAtMillis"`
}

func AssertExposuresHasCorrectSchema(db *sql.DB) {
	stmt, err := db.Prepare(
		"select card_id, created_at_millis from exposures limit 1")
	if err != nil {
		log.Fatalf("Error from db.Prepare: %s", err)
	}
	defer stmt.Close()
}

func SelectAllFromExposures(db *sql.DB) []Exposure {
	exposures := []Exposure{}

	rows, err := db.Query("select card_id, created_at_millis from exposures")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var exposure Exposure
		err = rows.Scan(&exposure.CardId, &exposure.CreatedAtMillis)
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
		"insert into exposures(card_id, created_at_millis) values(?,?)")
	if err != nil {
		log.Fatalf("Error from tx.Prepare: %s", err)
	}
	defer stmt.Close()

	for _, exposure := range exposures {
		_, err = stmt.Exec(exposure.CardId, exposure.CreatedAtMillis)
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()
}
