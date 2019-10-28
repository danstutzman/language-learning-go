package model

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"time"
)

var GRADE_TO_IS_FAILURE = map[string]bool{
	"RIGHT_WITH_MNEMONIC":           true,
	"MISCONNECTED_WITH_MNEMONIC":    true,
	"MISCONNECTED_WITHOUT_MNEMONIC": true,
	"BLANK":                         true,
	"WRONG_WITH_MNEMONIC":           true,
	"WRONG_WITHOUT_MNEMONIC":        true,
}

type SkillList struct {
	Skills []Skill `json:"skills"`
}

type Skill struct {
	Card        Card   `json:"card"`
	State       string `json:"state"`
	NumFailures int    `json:"numFailures"`
}

func (model *Model) ListSkills() SkillList {
	answers := db.FromAnswers(model.db,
		"WHERE shown_at IS NOT NULL "+
			"ORDER BY id")

	lastAnswerByCardId := map[int]db.AnswerRow{}
	for _, answer := range answers {
		lastAnswerByCardId[answer.CardId] = answer
	}

	numFailuresByCardId := map[int]int{}
	for _, answer := range answers {
		if answer.Grade.Valid && GRADE_TO_IS_FAILURE[answer.Grade.String] {
			numFailuresByCardId[answer.CardId] += 1
		}
	}

	cardRows := db.FromCards(model.db, "")

	cards := model.cardRowsToCardsJoinMorphemes(cardRows)

	skills := []Skill{}
	for _, card := range cards {
		lastAnswer, hasLastAnswer := lastAnswerByCardId[card.Id]
		oneDayAgo := time.Now().UTC().AddDate(0, 0, -1)

		var state string
		if !hasLastAnswer {
			state = "UNTESTED"
		} else if !lastAnswer.Grade.Valid {
			state = "NEEDS_GRADE"
		} else if lastAnswer.Grade.String == "RIGHT_WITHOUT_MNEMONIC" &&
			lastAnswer.ShownAt.After(oneDayAgo) {
			state = "RETEST_IN_1D"
		} else if lastAnswer.Grade.String == "BLANK" &&
			(card.Mnemonic21.String == "" || card.Mnemonic12.String == "") {
			state = "NEEDS_MNEMONIC"
		} else {
			state = "OKAY?"
		}

		skill := Skill{
			Card:        card,
			State:       state,
			NumFailures: numFailuresByCardId[card.Id],
		}
		skills = append(skills, skill)
	}

	return SkillList{Skills: skills}
}
