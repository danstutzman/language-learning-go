package db

import (
	"database/sql"
	"fmt"
	"log"
)

type CardRow struct {
	Id             int
	IsSentence     bool
	L2             string
	MorphemeIdsCsv string
	Type           string
}

func AssertCardsHasCorrectSchema(db *sql.DB) {
	query := `SELECT id, is_sentence, l2, morpheme_ids_csv, type
		FROM cards LIMIT 1`
	if LOG {
		log.Println(query)
	}

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func FromCards(db *sql.DB, whereLimit string) []CardRow {
	rows := []CardRow{}

	query := `SELECT id, is_sentence, l2, morpheme_ids_csv, type FROM cards ` +
		whereLimit
	if LOG {
		log.Println(query)
	}
	rset, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rset.Close()

	for rset.Next() {
		var row CardRow
		err = rset.Scan(&row.Id,
			&row.IsSentence,
			&row.L2,
			&row.MorphemeIdsCsv,
			&row.Type)
		if err != nil {
			panic(err)
		}
		rows = append(rows, row)
	}

	err = rset.Err()
	if err != nil {
		panic(err)
	}

	return rows
}

func InsertCard(db *sql.DB, card CardRow) CardRow {
	query := fmt.Sprintf(`INSERT INTO cards
	  (is_sentence, l2, morpheme_ids_csv, type)
		VALUES (%s, %s, %s, %s)`,
		EscapeBool(card.IsSentence),
		Escape(card.L2),
		Escape(card.MorphemeIdsCsv),
		Escape(card.Type))
	if LOG {
		log.Println(query)
	}

	result, err := db.Exec(query)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	card.Id = int(id)

	return card
}

func UpdateCard(db *sql.DB, card *CardRow) {
	query := fmt.Sprintf(
		`UPDATE cards SET is_sentence=%s, l2=%s,
		morpheme_ids_csv=%s, type=%s WHERE id=%d`,
		EscapeBool(card.IsSentence),
		Escape(card.L2),
		Escape(card.MorphemeIdsCsv),
		Escape(card.Type),
		card.Id)
	if LOG {
		log.Println(query)
	}

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func DeleteFromCards(db *sql.DB, where string) {
	query := "DELETE FROM cards " + where
	if LOG {
		log.Println(query)
	}

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}
