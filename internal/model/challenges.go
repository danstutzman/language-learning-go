package model

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"fmt"
	"gopkg.in/guregu/null.v3"
)

type ChallengeList struct {
	Challenges []Challenge `json:"challenges"`
}

type Challenge struct {
	Id             int         `json:"id"`
	Type           string      `json:"type"`
	CardId         int         `json:"cardId"`
	AnsweredL2     null.String `json:"answeredL2"`
	AnsweredAt     null.Time   `json:"answeredAt"`
	ShowedMnemonic bool        `json:"showedMnemonic"`
}

func challengeToChallengeRow(challenge Challenge) db.ChallengeRow {
	return db.ChallengeRow{
		Id:             challenge.Id,
		Type:           challenge.Type,
		CardId:         challenge.CardId,
		AnsweredL2:     challenge.AnsweredL2,
		AnsweredAt:     challenge.AnsweredAt,
		ShowedMnemonic: challenge.ShowedMnemonic,
	}
}

func challengeRowToChallenge(row db.ChallengeRow) Challenge {
	return Challenge{
		Id:             row.Id,
		Type:           row.Type,
		CardId:         row.CardId,
		AnsweredL2:     row.AnsweredL2,
		AnsweredAt:     row.AnsweredAt,
		ShowedMnemonic: row.ShowedMnemonic,
	}
}

func (model *Model) ListChallenges() ChallengeList {
	challengeRows := db.FromChallenges(model.db)

	challenges := []Challenge{}
	for _, challengeRow := range challengeRows {
		challenges = append(challenges, challengeRowToChallenge(challengeRow))
	}

	return ChallengeList{Challenges: challenges}
}

func (model *Model) ReplaceChallenge(challenge Challenge) {
	db.DeleteFromChallenges(model.db, fmt.Sprintf("WHERE card_id=%d AND type=%s",
		challenge.CardId, db.Escape(challenge.Type)))

	db.InsertChallenge(model.db, challengeToChallengeRow(challenge))
}

func (model *Model) GetTopGiven1Type2CardId() int {
	return db.GetTopGiven1Type2CardId(model.db)
}
