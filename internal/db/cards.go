package db

import (
	"database/sql"
	"fmt"
	"gopkg.in/guregu/null.v3"
	"log"
	"time"
)

type CardRow struct {
	Id             int
	L1             string
	L2             string
	LastAnsweredAt null.Time
	Mnemonic12     string
	Mnemonic21     string
	MorphemeIdsCsv string
	NounGender     string
	Type           string
}

func AssertCardsHasCorrectSchema(db *sql.DB) {
	query := `SELECT id, l1, l2, last_answered_at, mnemonic12, mnemonic21,
	  morpheme_ids_csv, noun_gender, type
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

	query := `SELECT id, l1, l2, last_answered_at, mnemonic12, mnemonic21,
	  morpheme_ids_csv, noun_gender, type FROM cards ` + whereLimit
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
			&row.L1,
			&row.L2,
			&row.LastAnsweredAt,
			&row.Mnemonic12,
			&row.Mnemonic21,
			&row.MorphemeIdsCsv,
			&row.NounGender,
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
	query := fmt.Sprintf(`INSERT INTO cards (l1, l2, last_answered_at,
	  mnemonic12, mnemonic21, morpheme_ids_csv, noun_gender, type)
		VALUES (%s, %s, %s, %s, %s, %s, %s, %s)`,
		Escape(card.L1),
		Escape(card.L2),
		EscapeNullTime(card.LastAnsweredAt),
		Escape(card.Mnemonic12),
		Escape(card.Mnemonic21),
		Escape(card.MorphemeIdsCsv),
		Escape(card.NounGender),
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
		`UPDATE cards SET l1=%s, l2=%s, last_answered_at=%s, mnemonic12=%s,
			mnemonic21=%s, morpheme_ids_csv=%s, noun_gender=%s, type=%s WHERE id=%d`,
		Escape(card.L1),
		Escape(card.L2),
		EscapeNullTime(card.LastAnsweredAt),
		Escape(card.Mnemonic12),
		Escape(card.Mnemonic21),
		Escape(card.MorphemeIdsCsv),
		Escape(card.NounGender),
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

func TouchCardLastAnsweredAt(db *sql.DB, cardId int) {
	query := fmt.Sprintf(
		`UPDATE cards SET last_answered_at=%s WHERE id=%d`,
		EscapeTime(time.Now().UTC()), cardId)
	if LOG {
		log.Println(query)
	}

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}
