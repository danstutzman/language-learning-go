package db

import (
	"database/sql"
	"fmt"
	"gopkg.in/guregu/null.v3"
	"log"
	"strconv"
	"strings"
	"time"
)

type AnswerRow struct {
	Id     int
	Type   string
	CardId int

	Expectation string

	ShownAt        time.Time
	AnsweredL1     null.String
	AnsweredL2     null.String
	ShowedMnemonic bool
	FirstKeyMillis int
	LastKeyMillis  int
}

type AnswerUpdate struct {
	Id     int
	CardId int

	ShownAt        null.Time
	AnsweredL1     null.String
	AnsweredL2     null.String
	ShowedMnemonic null.Bool
	FirstKeyMillis null.Int
	LastKeyMillis  null.Int
}

func AssertAnswersHasCorrectSchema(db *sql.DB) {
	query := `SELECT id, type, card_id,
  		expectation,
	    shown_at, answered_l1, showed_mnemonic,
			first_key_millis, last_key_millis
	  FROM answers
		LIMIT 1`
	if LOG {
		log.Println(query)
	}

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func FromAnswers(db *sql.DB, where string) []AnswerRow {
	rows := []AnswerRow{}

	query := `SELECT id, type, card_id,
  	expectation,
	  shown_at, answered_l1, answered_l2, showed_mnemonic,
		first_key_millis, last_key_millis
	  FROM answers ` + where
	if LOG {
		log.Println(query)
	}
	rset, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rset.Close()

	for rset.Next() {
		var row AnswerRow
		err = rset.Scan(&row.Id, &row.Type, &row.CardId,
			&row.Expectation,
			&row.ShownAt, &row.AnsweredL1, &row.AnsweredL2, &row.ShowedMnemonic,
			&row.FirstKeyMillis, &row.LastKeyMillis)
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

func InsertAnswer(db *sql.DB, answer AnswerRow) AnswerRow {
	query := fmt.Sprintf(`INSERT INTO answers
	(type, card_id,
		expectation,
	  shown_at, answered_l1, answered_l2, showed_mnemonic,
		first_key_millis, last_key_millis)
		VALUES (%s, %d,
		  %s,
		  %s, %s, %s, %s,
			%d, %d)`,
		Escape(answer.Type),
		answer.CardId,

		Escape(answer.Expectation),

		EscapeTime(answer.ShownAt),
		EscapeNullString(answer.AnsweredL1),
		EscapeNullString(answer.AnsweredL2),
		EscapeBool(answer.ShowedMnemonic),
		answer.FirstKeyMillis,
		answer.LastKeyMillis)

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

func UpdateAnswer(db *sql.DB, update AnswerUpdate) {
	pairs := []string{}
	if update.ShownAt.Valid {
		pairs = append(pairs, "shown_at="+EscapeNullTime(update.ShownAt))
	}
	if update.AnsweredL1.Valid {
		pairs = append(pairs, "answered_l1="+EscapeNullString(update.AnsweredL1))
	}
	if update.AnsweredL2.Valid {
		pairs = append(pairs, "answered_l2="+EscapeNullString(update.AnsweredL2))
	}
	if update.FirstKeyMillis.Valid {
		pairs = append(pairs,
			"first_key_millis="+EscapeNullInt(update.FirstKeyMillis))
	}
	if update.LastKeyMillis.Valid {
		pairs = append(pairs, "last_key_millis="+EscapeNullInt(update.LastKeyMillis))
	}
	if update.ShowedMnemonic.Valid {
		pairs = append(pairs, "showed_mnemonic="+
			EscapeNullBool(update.ShowedMnemonic))
	}

	if len(pairs) == 0 {
		return
	}

	query := "UPDATE answers SET " +
		strings.Join(pairs, ", ") +
		" WHERE id=" + strconv.Itoa(update.Id)
	if LOG {
		log.Println(query)
	}

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func DeleteFromAnswers(db *sql.DB, where string) {
	query := "DELETE FROM answers " + where
	if LOG {
		log.Println(query)
	}

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}
