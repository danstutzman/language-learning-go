package db

import (
	"database/sql"
	"fmt"
	"gopkg.in/guregu/null.v3"
	"log"
)

type ChallengeRow struct {
	Id             int
	Type           string
	CardId         int
	AnsweredL2     null.String
	AnsweredAt     null.Time
	ShowedMnemonic bool
}

func AssertChallengesHasCorrectSchema(db *sql.DB) {
	query := `SELECT id, card_id, answered_l2, answered_at, showed_mnemonic 
	  FROM challenges LIMIT 1`
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

	query := `SELECT id, card_id, answered_l2, answered_at, showed_mnemonic
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
		err = rset.Scan(&row.Id,
			&row.CardId,
			&row.AnsweredL2,
			&row.AnsweredAt,
			&row.ShowedMnemonic)
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
	(type, card_id, answered_l2, answered_at, showed_mnemonic)
		VALUES (%s, %d, %s, %s, %s)`, Escape(challenge.Type), challenge.CardId,
		EscapeNullString(challenge.AnsweredL2),
		EscapeNullTime(challenge.AnsweredAt),
		EscapeBool(challenge.ShowedMnemonic))
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

func GetTopGiven1Type2CardId(db *sql.DB) int {
	query := `SELECT card_id
		FROM challenges
    WHERE challenges.type = 'Given1Type2'
		ORDER BY answered_at
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
