package model

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"fmt"
	"gopkg.in/guregu/null.v3"
)

type AnswerList struct {
	Answers []Answer `json:"answers"`
}

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

func answerRowToAnswer(row db.AnswerRow) Answer {
	return Answer{
		Id:             row.Id,
		Type:           row.Type,
		CardId:         row.CardId,
		AnsweredL2:     row.AnsweredL2,
		AnsweredAt:     row.AnsweredAt,
		ShowedMnemonic: row.ShowedMnemonic,
	}
}

func (model *Model) ListAnswers() AnswerList {
	answerRows := db.FromAnswers(model.db)

	answers := []Answer{}
	for _, answerRow := range answerRows {
		answers = append(answers, answerRowToAnswer(answerRow))
	}

	return AnswerList{Answers: answers}
}

func (model *Model) ReplaceAnswer(answer Answer) {
	db.DeleteFromAnswers(model.db, fmt.Sprintf("WHERE card_id=%d AND type=%s",
		answer.CardId, db.Escape(answer.Type)))

	db.InsertAnswer(model.db, answerToAnswerRow(answer))
}

func (model *Model) GetTopGiven1Type2CardId() int {
	return db.GetTopGiven1Type2CardId(model.db)
}
