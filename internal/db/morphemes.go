package db

import (
	"database/sql"
	"log"
)

type Morpheme struct {
	Id              int     `json:"id"`
	L2              string  `json:"l2"`
	Gloss           string  `json:"gloss"`
	CreatedAtMillis float64 `json:"created_at_millis"`
	UpdatedAtMillis float64 `json:"updated_at_millis"`
}

func AssertMorphemesHasCorrectSchema(db *sql.DB) {
	stmt, err := db.Prepare(`
		select id, l2, gloss, created_at_millis, updated_at_millis
		from morphemes
		limit 1
	`)
	if err != nil {
		log.Fatalf("Error from db.Prepare in AssertMorphemesHasCorrectSchema: %s", err)
	}
	defer stmt.Close()
}

func SelectAllFromMorphemes(db *sql.DB) []Morpheme {
	morphemes := []Morpheme{}

	rows, err := db.Query(
		"select id, l2, gloss, created_at_millis, updated_at_millis from morphemes limit 20")
	if err != nil {
		log.Fatalf("Error from db.Query in SelectAllFromMorphemes: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var morpheme Morpheme
		err = rows.Scan(&morpheme.Id,
			&morpheme.L2,
			&morpheme.Gloss,
			&morpheme.CreatedAtMillis,
			&morpheme.UpdatedAtMillis)
		if err != nil {
			log.Fatalf("Error from rows.Scan in SelectAllFromMorphemes: %s", err)
		}
		morphemes = append(morphemes, morpheme)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("Error from rows.Err in SelectAllFromMorphemes: %s", err)
	}

	return morphemes
}

func UpsertMorpheme(morpheme *Morpheme, db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Error from db.Begin in UpsertMorpheme: %s", err)
	}

	stmt, err := tx.Prepare(
		`INSERT INTO morphemes(id, l2, gloss, created_at_millis, updated_at_millis)
		VALUES(?, ?, ?, ?, ?)
	  ON CONFLICT(id) DO UPDATE SET
			l2=excluded.l2,
			gloss=excluded.gloss,
			created_at_millis=excluded.created_at_millis,
			updated_at_millis=excluded.updated_at_millis`)
	if err != nil {
		log.Fatalf("Error from tx.Prepare in UpsertMorpheme: %s", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		morpheme.Id,
		morpheme.L2,
		morpheme.Gloss,
		morpheme.CreatedAtMillis,
		morpheme.UpdatedAtMillis)
	if err != nil {
		log.Fatalf("Error from stmt.Exec in UpsertMorpheme: %s", err)
	}

	tx.Commit()
}

func FindMorphemeById(db *sql.DB, id string) (Morpheme, error) {
	row := db.QueryRow(
		"select id, l2, gloss, created_at_millis, updated_at_millis from morphemes where id=$1", id)

	morpheme := Morpheme{}
	err := row.Scan(
		&morpheme.Id,
		&morpheme.L2,
		&morpheme.Gloss,
		&morpheme.CreatedAtMillis,
		&morpheme.UpdatedAtMillis)

	return morpheme, err
}
