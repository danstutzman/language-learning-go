package model

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"gopkg.in/guregu/null.v3"
	"time"
)

type AnswerList struct {
	Answers []Answer `json:"answers"`
}

type Answer struct {
	Id        int                    `json:"id"`
	Type      string                 `json:"type"`
	CardId    int                    `json:"cardId"`
	Card      *Card                  `json:"card"`
	Morphemes []db.AnswerMorphemeRow `json:"morphemes"`

	Expectation string `json:"expectation"`

	ShownAt        time.Time   `json:"shownAt"`
	AnsweredL1     null.String `json:"answeredL1"`
	AnsweredL2     null.String `json:"answeredL2"`
	ShowedMnemonic bool        `json:"showedMnemonic"`
	FirstKeyMillis int         `json:"firstKeyMillis"`
	LastKeyMillis  int         `json:"lastKeyMillis"`
}

func answerToAnswerRow(answer Answer) db.AnswerRow {
	return db.AnswerRow{
		Id:     answer.Id,
		Type:   answer.Type,
		CardId: answer.CardId,

		Expectation: answer.Expectation,

		ShownAt:        answer.ShownAt,
		AnsweredL1:     answer.AnsweredL1,
		AnsweredL2:     answer.AnsweredL2,
		ShowedMnemonic: answer.ShowedMnemonic,
		FirstKeyMillis: answer.FirstKeyMillis,
		LastKeyMillis:  answer.LastKeyMillis,
	}
}

func answerRowToAnswer(row db.AnswerRow) Answer {
	return Answer{
		Id:     row.Id,
		Type:   row.Type,
		CardId: row.CardId,

		Expectation: row.Expectation,

		ShownAt:        row.ShownAt,
		AnsweredL1:     row.AnsweredL1,
		AnsweredL2:     row.AnsweredL2,
		ShowedMnemonic: row.ShowedMnemonic,
		FirstKeyMillis: row.FirstKeyMillis,
		LastKeyMillis:  row.LastKeyMillis,
	}
}

func (model *Model) ListAnswers() AnswerList {
	answerRows := db.FromAnswers(model.db, "ORDER BY card_id, id")

	answers := []Answer{}
	for _, answerRow := range answerRows {
		answers = append(answers, answerRowToAnswer(answerRow))
	}

	cardIds := []int{}
	answerIds := []int{}
	for _, answer := range answers {
		cardIds = append(cardIds, answer.CardId)
		answerIds = append(answerIds, answer.Id)
	}

	cardRows := db.FromCards(model.db, "WHERE "+db.InIntList("id", cardIds))

	cardById := map[int]Card{}
	for _, row := range cardRows {
		cardById[row.Id] = model.cardRowToCardJoinMorphemes(row)
	}

	answerMorphemes := db.FromAnswerMorphemes(model.db,
		"WHERE "+db.InIntList("answer_id", answerIds))

	answerMorphemesById := map[int][]db.AnswerMorphemeRow{}
	for _, answerMorpheme := range answerMorphemes {
		answerMorphemesById[answerMorpheme.AnswerId] =
			append(answerMorphemesById[answerMorpheme.AnswerId], answerMorpheme)
	}

	for i, answer := range answers {
		card := cardById[answers[i].CardId]
		answers[i].Card = &card
		answers[i].Morphemes = answerMorphemesById[answer.Id]
	}

	return AnswerList{Answers: answers}
}

func gatherAnsweredL2(alignments []Alignment, answeredL2Runes []rune,
	begin int, end int) []rune {
	out := []rune{}
	for _, alignment := range alignments {
		if alignment.X >= begin && alignment.X < end {
			if alignment.Y != -1 {
				out = append(out, answeredL2Runes[alignment.Y])
			}
		}
	}
	return out
}

func (model *Model) InsertAnswer(answer Answer) {
	answerRow := db.InsertAnswer(model.db, answerToAnswerRow(answer))

	card := model.GetCardJoinMorphemes(answer.CardId)
	answeredL2Runes := []rune(answer.AnsweredL2.String)
	alignments := AlignRuneArrays([]rune(card.L2), answeredL2Runes)

	for _, cardMorpheme := range card.Morphemes {
		db.InsertAnswerMorpheme(model.db, db.AnswerMorphemeRow{
			AnswerId:   answerRow.Id,
			MorphemeId: cardMorpheme.Morpheme.Id,
			ShownAt:    answerRow.ShownAt,
			CorrectL2:  cardMorpheme.L2,
			AlignedL2: string(gatherAnsweredL2(alignments, answeredL2Runes,
				cardMorpheme.Begin, cardMorpheme.Begin+len([]rune(cardMorpheme.L2)))),
		})
	}
}

func updateLastSeenAt(morpheme db.MorphemeRow,
	lastSeenAt null.Time) db.MorphemeRow {
	morpheme.LastSeenAt = lastSeenAt
	return morpheme
}
