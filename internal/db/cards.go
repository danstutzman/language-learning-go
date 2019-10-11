package db

import (
	"database/sql"
	"fmt"
	"log"
)

type CardRow struct {
	Id             int
	Type           string
	L1             string
	L2             string
	MorphemeIdsCsv string
}

func AssertCardsHasCorrectSchema(db *sql.DB) {
	query := "SELECT id, type, l1, l2, morpheme_ids_csv FROM cards LIMIT 1"
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

	query := "SELECT id, type, l1, l2, morpheme_ids_csv FROM cards " + whereLimit
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
			&row.Type,
			&row.L1,
			&row.L2,
			&row.MorphemeIdsCsv)
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
	query := fmt.Sprintf(`INSERT INTO cards (type, l1, l2, morpheme_ids_csv)
		VALUES (%s, %s, %s, %s)`,
		Escape(card.Type), Escape(card.L1), Escape(card.L2),
		Escape(card.MorphemeIdsCsv))
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
		"UPDATE cards SET type=%s, l1=%s, l2=%s, morpheme_ids_csv=%s WHERE id=%d",
		Escape(card.Type), Escape(card.L1), Escape(card.L2),
		Escape(card.MorphemeIdsCsv), card.Id)
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
