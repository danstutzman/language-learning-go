package db

import (
	"database/sql"
	"fmt"
	"log"
)

type CardsMorphemesRow struct {
	CardId     int
	MorphemeId int
}

func AssertCardsMorphemesHasCorrectSchema(db *sql.DB) {
	query := "SELECT card_id, morpheme_id FROM cards_morphemes LIMIT 1"
	log.Println(query)

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func FromCardsMorphemes(db *sql.DB, where string) []CardsMorphemesRow {
	rows := []CardsMorphemesRow{}

	query := "SELECT card_id, morpheme_id FROM cards_morphemes " + where
	log.Println(query)

	rset, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rset.Close()

	for rset.Next() {
		var row CardsMorphemesRow
		err = rset.Scan(&row.CardId, &row.MorphemeId)
		if err != nil {
			panic(err)
		}
		rows = append(rows, row)
	}

	err = rset.Err()
	if err != nil {
		panic(err)
	}
	defer rset.Close()

	return rows
}

func DeleteFromCardsMorphemes(db *sql.DB, where string) {
	query := "DELETE FROM cards_morphemes " + where
	log.Println(query)

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func InsertCardsMorphemesRow(db *sql.DB, row CardsMorphemesRow) {
	query := fmt.Sprintf(`INSERT INTO cards_morphemes (card_id, morpheme_id) VALUES (%d, %d)`,
		row.CardId, row.MorphemeId)
	log.Println(query)

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}
