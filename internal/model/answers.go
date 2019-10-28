package model

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"fmt"
	"gopkg.in/guregu/null.v3"
	"strconv"
	"strings"
	"time"
)

type AnswerList struct {
	Answers []Answer `json:"answers"`
}

type Answer struct {
	Id     int    `json:"id"`
	Type   string `json:"type"`
	CardId int    `json:"cardId"`
	Card   *Card  `json:"card"`

	Expectation string `json:"expectation"`

	ShownAt        time.Time   `json:"shownAt"`
	AnsweredL1     null.String `json:"answeredL1"`
	AnsweredL2     null.String `json:"answeredL2"`
	ShowedMnemonic bool        `json:"showedMnemonic"`
	FirstKeyMillis int         `json:"firstKeyMillis"`
	LastKeyMillis  int         `json:"lastKeyMillis"`

	Grade              null.String `json:"grade"`
	MisconnectedCardId null.Int    `json:"misconnectedCardId"`
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

		Grade:              answer.Grade,
		MisconnectedCardId: answer.MisconnectedCardId,
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

		Grade:              row.Grade,
		MisconnectedCardId: row.MisconnectedCardId,
	}
}

func (model *Model) ListAnswers() AnswerList {
	answerRows := db.FromAnswers(model.db, "ORDER BY card_id, id")

	answers := []Answer{}
	for _, answerRow := range answerRows {
		answers = append(answers, answerRowToAnswer(answerRow))
	}

	cardIds := []int{}
	for _, answer := range answers {
		cardIds = append(cardIds, answer.CardId)
	}

	cardRows := db.FromCards(model.db, "WHERE "+db.InIntList("id", cardIds))

	cardById := map[int]Card{}
	for _, row := range cardRows {
		cardById[row.Id] = model.cardRowToCardJoinMorphemes(row)
	}

	for i, _ := range answers {
		card := cardById[answers[i].CardId]
		answers[i].Card = &card
	}

	return AnswerList{Answers: answers}
}

func (model *Model) InsertAnswer(answer Answer) Answer {
	row := db.InsertAnswer(model.db, answerToAnswerRow(answer))
	return answerRowToAnswer(row)
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func (model *Model) GetTopChallenges(type_ string) CardList {
	allCards := db.FromCards(model.db, "")

	cardById := map[int]db.CardRow{}
	for _, card := range allCards {
		cardById[card.Id] = card
	}

	allMorphemes := db.FromMorphemes(model.db, "")

	morphemeById := map[int]db.MorphemeRow{}
	for _, morpheme := range allMorphemes {
		morphemeById[morpheme.Id] = morpheme
	}

	now := time.Now().UTC()

	cardRows := []db.CardRow{}
	for i := 0; i < 10; i += 1 {
		stalestCardId := 0
		stalestStaleness := time.Duration(0)
		for _, card := range allCards {
			cardStaleness := time.Duration(0)
			for _, morphemeId := range strings.Split(card.MorphemeIdsCsv, ",") {
				morpheme := morphemeById[mustAtoi(morphemeId)]
				morphemeStaleness := now.Sub(morpheme.LastSeenAt.Time)
				if morphemeStaleness > time.Duration(10000)*time.Hour {
					morphemeStaleness = time.Duration(10000) * time.Hour
				}
				cardStaleness += morphemeStaleness
			}

			if cardStaleness > stalestStaleness {
				stalestCardId = card.Id
				stalestStaleness = cardStaleness
			}
		}

		topCard := cardById[stalestCardId]
		if false {
			fmt.Printf("%-50s -> %s\n", topCard.L2, stalestStaleness)
			for _, morphemeId := range strings.Split(topCard.MorphemeIdsCsv, ",") {
				morpheme := morphemeById[mustAtoi(morphemeId)]
				fmt.Printf("  %s -> %s\n", morpheme.L2,
					now.Sub(morpheme.LastSeenAt.Time))
			}
		}

		cardRows = append(cardRows, topCard)

		// simulate morpheme staleness being updated so that each card in the
		// return challengeList has a non-overlapping set of morphemes.
		topCard = cardById[stalestCardId]
		for _, morphemeId := range strings.Split(topCard.MorphemeIdsCsv, ",") {
			morphemeById[mustAtoi(morphemeId)] = updateLastSeenAt(
				morphemeById[mustAtoi(morphemeId)], null.TimeFrom(now))
		}
	}

	cards := model.cardRowsToCardsJoinMorphemes(cardRows)
	return CardList{Cards: cards}
}

func updateLastSeenAt(morpheme db.MorphemeRow,
	lastSeenAt null.Time) db.MorphemeRow {
	morpheme.LastSeenAt = lastSeenAt
	return morpheme
}
