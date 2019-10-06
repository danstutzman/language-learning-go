package db

import (
	"database/sql"
	"log"
)

type Card struct {
	Id              int     `json:"id"`
	L1              string  `json:"l1"`
	L2              string  `json:"l2"`
	MorphemeIdsJson string  `json:"morpheme_ids_json"`
	CreatedAtMillis float64 `json:"created_at_millis"`
	UpdatedAtMillis float64 `json:"updated_at_millis"`
}

func AssertCardsHasCorrectSchema(db *sql.DB) {
	stmt, err := db.Prepare(`
		select id, l1, l2, morpheme_ids_json, created_at_millis, updated_at_millis
		from cards
		limit 1
	`)
	if err != nil {
		log.Fatalf("Error from db.Prepare in AssertCardsHasCorrectSchema: %s", err)
	}
	defer stmt.Close()
}

func SelectAllFromCards(db *sql.DB) []Card {
	cards := []Card{}

	rows, err := db.Query(
		"select id, l1, l2, morpheme_ids_json, created_at_millis, updated_at_millis from cards")
	if err != nil {
		log.Fatalf("Error from db.Query in SelectAllFromCards: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var card Card
		err = rows.Scan(&card.Id,
			&card.L1,
			&card.L2,
			&card.MorphemeIdsJson,
			&card.CreatedAtMillis,
			&card.UpdatedAtMillis)
		if err != nil {
			log.Fatalf("Error from rows.Scan in SelectAllFromCards: %s", err)
		}
		cards = append(cards, card)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("Error from rows.Err in SelectAllFromCards: %s", err)
	}

	return cards
}

func UpsertCard(card *Card, db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Error from db.Begin in UpsertCard: %s", err)
	}

	stmt, err := tx.Prepare(
		`INSERT INTO cards(id, l1, l2, morpheme_ids_json, created_at_millis, updated_at_millis)
		VALUES(?, ?, ?, ?, ?, ?)
	  ON CONFLICT(id) DO UPDATE SET
			l1=excluded.l1,
			l2=excluded.l2,
			morpheme_ids_json=excluded.morpheme_ids_json,
			created_at_millis=excluded.created_at_millis,
			updated_at_millis=excluded.updated_at_millis`)
	if err != nil {
		log.Fatalf("Error from tx.Prepare in UpsertCard: %s", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		card.Id,
		card.L1,
		card.L2,
		card.MorphemeIdsJson,
		card.CreatedAtMillis,
		card.UpdatedAtMillis)
	if err != nil {
		log.Fatalf("Error from stmt.Exec in UpsertCard: %s", err)
	}

	tx.Commit()
}

func FindCardById(db *sql.DB, id string) (Card, error) {
	row := db.QueryRow(
		`select id, l1, l2, morpheme_ids_json, created_at_millis, updated_at_millis
		from cards
		where id=$1`, id)

	card := Card{}
	err := row.Scan(
		&card.Id,
		&card.L1,
		&card.L2,
		&card.MorphemeIdsJson,
		&card.CreatedAtMillis,
		&card.UpdatedAtMillis)

	return card, err
}
