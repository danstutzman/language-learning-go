package db

import (
	"database/sql"
	"fmt"
	"gopkg.in/guregu/null.v3"
	"log"
	"strconv"
	"strings"
)

type ChallengeRow struct {
	Id     int
	Type   string
	CardId int

	Expectation string

	ShownAt        null.Time
	AnsweredL1     null.String
	AnsweredL2     null.String
	ShowedMnemonic null.Bool
	FirstKeyMillis null.Int
	LastKeyMillis  null.Int

	Grade              null.String
	MisconnectedCardId null.Int
}

type ChallengeUpdate struct {
	Id int

	ShownAt        null.Time
	AnsweredL1     null.String
	AnsweredL2     null.String
	ShowedMnemonic null.Bool
	FirstKeyMillis null.Int
	LastKeyMillis  null.Int

	Grade              null.String
	MisconnectedCardId null.Int
}

func AssertChallengesHasCorrectSchema(db *sql.DB) {
	query := `SELECT id, type, card_id,
  		expectation,
	    shown_at, answered_l1, showed_mnemonic,
			first_key_millis, last_key_millis,
		  grade, misconnected_card_id
	  FROM challenges
		LIMIT 1`
	if LOG {
		log.Println(query)
	}

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func FromChallenges(db *sql.DB, where string) []ChallengeRow {
	rows := []ChallengeRow{}

	query := `SELECT id, type, card_id,
  	expectation,
	  shown_at, answered_l1, answered_l2, showed_mnemonic,
		first_key_millis, last_key_millis,
	  grade, misconnected_card_id
	  FROM challenges ` + where
	if LOG {
		log.Println(query)
	}
	rset, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rset.Close()

	for rset.Next() {
		var row ChallengeRow
		err = rset.Scan(&row.Id, &row.Type, &row.CardId,
			&row.Expectation,
			&row.ShownAt, &row.AnsweredL1, &row.AnsweredL2, &row.ShowedMnemonic,
			&row.FirstKeyMillis, &row.LastKeyMillis,
			&row.Grade, &row.MisconnectedCardId)
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

func InsertChallenge(db *sql.DB, challenge ChallengeRow) ChallengeRow {
	query := fmt.Sprintf(`INSERT INTO challenges
	(type, card_id, grade,
		expectation,
	  shown_at, answered_l1, answered_l2, showed_mnemonic,
		first_key_millis, last_key_millis,
	  grade, misconnected_card_id)
		VALUES (%s, %d, %s,
		  %s,
		  %s, %s, %s, %s,
			%s, %s,
			%s, %s)`,
		Escape(challenge.Type),
		challenge.CardId,
		EscapeNullString(challenge.Grade),

		Escape(challenge.Expectation),

		EscapeNullTime(challenge.ShownAt),
		EscapeNullString(challenge.AnsweredL1),
		EscapeNullString(challenge.AnsweredL2),
		EscapeNullBool(challenge.ShowedMnemonic),
		EscapeNullInt(challenge.FirstKeyMillis),
		EscapeNullInt(challenge.LastKeyMillis),

		EscapeNullString(challenge.Grade),
		EscapeNullInt(challenge.MisconnectedCardId))

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
	challenge.Id = int(id)

	return challenge
}

func UpdateChallenge(db *sql.DB, update ChallengeUpdate) {
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
	if update.Grade.Valid {
		pairs = append(pairs, "grade="+EscapeNullString(update.Grade))
	}
	if update.ShowedMnemonic.Valid {
		pairs = append(pairs, "showed_mnemonic="+
			EscapeNullBool(update.ShowedMnemonic))
	}
	if update.MisconnectedCardId.Valid {
		pairs = append(pairs, "misconnected_card_id="+
			EscapeNullInt(update.MisconnectedCardId))
	}

	if len(pairs) == 0 {
		return
	}

	query := "UPDATE challenges SET " +
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

func DeleteFromChallenges(db *sql.DB, where string) {
	query := "DELETE FROM challenges " + where
	if LOG {
		log.Println(query)
	}

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}
