package db

import (
	"database/sql"
	"fmt"
	"gopkg.in/guregu/null.v3"
	"log"
)

type AnswerRow struct {
	Id         int
	Type       string
	CardId     int
	AnsweredL2 null.String
	AnsweredAt null.Time
}

func AssertAnswersHasCorrectSchema(db *sql.DB) {
	query := "SELECT id FROM answers LIMIT 1"
	if LOG {
		log.Println(query)
	}

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func InsertAnswer(db *sql.DB, answer AnswerRow) AnswerRow {
	query := fmt.Sprintf(`INSERT INTO answers
	(type, card_id, answered_l2, answered_at)
		VALUES (%s, %d, %s, %s)`, Escape(answer.Type), answer.CardId,
		EscapeNullString(answer.AnsweredL2), EscapeNullTime(answer.AnsweredAt))
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
	answer.Id = int(id)

	return answer
}

func GetTopGiven1Type2CardId(db *sql.DB) int {
	query := `SELECT cards.id
	  FROM cards
	  LEFT JOIN answers ON answers.card_id = cards.id
		GROUP BY cards.id
		ORDER BY MAX(answers.answered_at)
		LIMIT 1`
	if LOG {
		log.Println(query)
	}

	var cardId int
	rset := db.QueryRow(query)
	switch err := rset.Scan(&cardId); err {
	case sql.ErrNoRows:
		return 0
	case nil:
		return cardId
	default:
		panic(err)
	}
}
