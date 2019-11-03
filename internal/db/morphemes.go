package db

import (
	"database/sql"
	"fmt"
	"gopkg.in/guregu/null.v3"
	"log"
)

type MorphemeRow struct {
	Id       int
	Type     string
	L2       string
	Lemma    null.String
	Tag      null.String
	Nonsense null.String
}

func AssertMorphemesHasCorrectSchema(db *sql.DB) {
	query := "SELECT id, type, l2, lemma, tag, nonsense " +
		"FROM morphemes LIMIT 1"
	if LOG {
		log.Println(query)
	}

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func OneFromMorphemes(db *sql.DB, where string) *MorphemeRow {
	query := "SELECT id, type, l2, lemma, ta, nonsense  " +
		"FROM morphemes " + where
	if LOG {
		log.Println(query)
	}

	var row MorphemeRow
	rset := db.QueryRow(query)
	switch err := rset.Scan(&row.Id,
		&row.Type,
		&row.L2,
		&row.Lemma,
		&row.Tag,
		&row.Nonsense); err {
	case sql.ErrNoRows:
		return nil
	case nil:
		return &row
	default:
		panic(err)
	}
}

func FromMorphemes(db *sql.DB, whereLimit string) []MorphemeRow {
	query := "SELECT id, type, l2, lemma, tag, nonsense " +
		"FROM morphemes " + whereLimit
	if LOG {
		log.Println(query)
	}

	rset, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rset.Close()

	rows := []MorphemeRow{}
	for rset.Next() {
		var row MorphemeRow
		err = rset.Scan(&row.Id, &row.Type, &row.L2, &row.Lemma, &row.Tag,
			&row.Nonsense)
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
	query := fmt.Sprintf(`INSERT INTO morphemes
  	(type, l2, lemma, tag)
		VALUES (%s, %s, %s, %s)`, Escape(morpheme.Type), Escape(morpheme.L2),
		EscapeNullString(morpheme.Lemma), EscapeNullString(morpheme.Tag))
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
	morpheme.Id = int(id)

	return morpheme
}

func UpdateMorpheme(db *sql.DB, morpheme MorphemeRow) {
	query := fmt.Sprintf(`UPDATE morphemes
		SET type=%s, l2=%s, lemma=%s, tag=%s WHERE id=%d`,
		Escape(morpheme.Type), Escape(morpheme.L2),
		EscapeNullString(morpheme.Lemma),
		EscapeNullString(morpheme.Tag),
		morpheme.Id)
	if LOG {
		log.Println(query)
	}

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func DeleteFromMorphemes(db *sql.DB, where string) {
	query := "DELETE FROM morphemes " + where
	if LOG {
		log.Println(query)
	}

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}
