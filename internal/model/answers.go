package model

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"gopkg.in/guregu/null.v3"
)

type Answer struct {
	Id             int         `json:"id"`
	Type           string      `json:"type"`
	CardId         int         `json:"cardId"`
	AnsweredL2     null.String `json:"answeredL2"`
	AnsweredAt     null.Time   `json:"answeredAt"`
	ShowedMnemonic bool        `json:"showedMnemonic"`
}

func answerToAnswerRow(answer Answer) db.AnswerRow {
	return db.AnswerRow{
		Id:             answer.Id,
		Type:           answer.Type,
		CardId:         answer.CardId,
		AnsweredL2:     answer.AnsweredL2,
		AnsweredAt:     answer.AnsweredAt,
		ShowedMnemonic: answer.ShowedMnemonic,
	}
}

func (model *Model) InsertAnswer(answer Answer) {
	db.InsertAnswer(model.db, answerToAnswerRow(answer))
}

func (model *Model) GetTopGiven1Type2CardId() int {
	return db.GetTopGiven1Type2CardId(model.db)
}
