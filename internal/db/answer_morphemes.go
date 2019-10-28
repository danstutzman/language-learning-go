package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type AnswerMorphemeRow struct {
	Id         int       `json:"id"`
	AnswerId   int       `json:"answerId"`
	MorphemeId int       `json:"morphemeId"`
	ShownAt    time.Time `json:"shownAt"`
	CorrectL2  string    `json:"correctL2"`
	AlignedL2  string    `json:"alignedL2"`
}

func AssertAnswerMorphemesHasCorrectSchema(db *sql.DB) {
	query := `SELECT id, answer_id, morpheme_id, shown_at, correct_l2, aligned_l2
	  FROM answer_morphemes
		LIMIT 1`
	if LOG {
		log.Println(query)
	}

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func FromAnswerMorphemes(db *sql.DB, where string) []AnswerMorphemeRow {
	rows := []AnswerMorphemeRow{}

	query := `SELECT id, answer_id, morpheme_id, shown_at, correct_l2, aligned_l2
	  FROM answer_morphemes ` + where
	if LOG {
		log.Println(query)
	}
	rset, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rset.Close()

	for rset.Next() {
		var row AnswerMorphemeRow
		err = rset.Scan(&row.Id, &row.AnswerId, &row.MorphemeId,
			&row.ShownAt, &row.CorrectL2, &row.AlignedL2)
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

func InsertAnswerMorpheme(db *sql.DB, row AnswerMorphemeRow) AnswerMorphemeRow {
	query := fmt.Sprintf(`INSERT INTO answer_morphemes
  	(answer_id, morpheme_id, shown_at, correct_l2, aligned_l2)
		VALUES (%d, %d, %s, %s, %s)`,
		row.AnswerId, row.MorphemeId, EscapeTime(row.ShownAt),
		Escape(row.CorrectL2), Escape(row.AlignedL2))

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
	row.Id = int(id)

	return row
}

func DeleteFromAnswerMorphemes(db *sql.DB, where string) {
	query := "DELETE FROM answer_morphemes " + where
	if LOG {
		log.Println(query)
	}

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}
