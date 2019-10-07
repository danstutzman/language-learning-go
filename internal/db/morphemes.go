package db

import (
	"database/sql"
	"fmt"
	"log"
)

type MorphemeRow struct {
	Id    int    `json:"id"`
	L2    string `json:"l2"`
	Gloss string `json:"gloss"`
}

func AssertMorphemesHasCorrectSchema(db *sql.DB) {
	query := "SELECT id, l2, gloss FROM morphemes LIMIT 1"
	log.Println(query)

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func FromMorphemes(db *sql.DB, whereLimit string) []MorphemeRow {
	query := "SELECT id, l2, gloss FROM morphemes " + whereLimit
	log.Println(query)

	rset, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rset.Close()

	rows := []MorphemeRow{}
	for rset.Next() {
		var row MorphemeRow
		err = rset.Scan(&row.Id,
			&row.L2,
			&row.Gloss)
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

func InsertMorpheme(db *sql.DB, morpheme MorphemeRow) MorphemeRow {
	query := fmt.Sprintf(`INSERT INTO morphemes (l2, gloss) VALUES (%s, %s)`,
		Escape(morpheme.L2), Escape(morpheme.Gloss))
	log.Println(query)

	result, err := db.Exec(query)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	morpheme.Id = int(id)

	return morpheme
}

func UpdateMorpheme(db *sql.DB, morpheme MorphemeRow) {
	query := fmt.Sprintf("UPDATE morphemes SET l2=%s, gloss=%s WHERE id=%d",
		Escape(morpheme.L2), Escape(morpheme.Gloss), morpheme.Id)
	log.Println(query)

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func DeleteFromMorphemes(db *sql.DB, where string) {
	query := "DELETE FROM morphemes " + where
	log.Println(query)

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}
