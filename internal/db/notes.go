package db

import (
	"database/sql"
	"log"
)

type Note struct {
	Id              int    `json:"id"`
	Text            string `json:"text"`
	CreatedAtMillis int    `json:"createdAtMillis"`
	UpdatedAtMillis int    `json:"updatedAtMillis"`
}

func AssertNotesHasCorrectSchema(db *sql.DB) {
	stmt, err := db.Prepare(`
		select id, text, created_at_millis, updated_at_millis
		from notes
		limit 1
	`)
	if err != nil {
		log.Fatalf("Error from db.Prepare: %s", err)
	}
	defer stmt.Close()
}

func SelectAllFromNotes(db *sql.DB) []Note {
	notes := []Note{}

	rows, err := db.Query(
		"select id, text, created_at_millis, updated_at_millis from notes")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var note Note
		err = rows.Scan(&note.Id,
			&note.Text,
			&note.CreatedAtMillis,
			&note.UpdatedAtMillis)
		if err != nil {
			log.Fatal(err)
		}
		notes = append(notes, note)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return notes
}

func UpsertNote(note *Note, db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Error from db.Begin: %s", err)
	}

	stmt, err := tx.Prepare(
		`INSERT INTO notes(id, text, created_at_millis, updated_At_millis)
		VALUES(?, ?, ?, ?)
	  ON CONFLICT(id) DO UPDATE SET
			text=excluded.text,
			created_at_millis=excluded.created_at_millis,
			updated_at_millis=excluded.updated_at_millis`)
	if err != nil {
		log.Fatalf("Error from tx.Prepare: %s", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		note.Id,
		note.Text,
		note.CreatedAtMillis,
		note.UpdatedAtMillis)
	if err != nil {
		log.Fatal(err)
	}

	tx.Commit()
}
