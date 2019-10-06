package db

import (
	"database/sql"
	"fmt"
	"log"
)

type CardRow struct {
	Id int    `json:"id"`
	L1 string `json:"l1"`
	L2 string `json:"l2"`
}

func AssertCardsHasCorrectSchema(db *sql.DB) {
	query := "SELECT id, l1, l2 FROM cards LIMIT 1"
	log.Println(query)

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func FromCards(db *sql.DB, whereLimit string) []CardRow {
	rows := []CardRow{}

	query := "SELECT id, l1, l2 FROM cards " + whereLimit
	log.Println(query)
	rset, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rset.Close()

	for rset.Next() {
		var row CardRow
		err = rset.Scan(&row.Id,
			&row.L1,
			&row.L2)
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

func UpdateCard(db *sql.DB, card *CardRow) {
	query := fmt.Sprintf("UPDATE cards SET l1=%s, l2=%s WHERE id=%d",
		Escape(card.L1), Escape(card.L2), card.Id)
	log.Println(query)

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}
