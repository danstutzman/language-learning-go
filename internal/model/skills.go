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
	challenges := db.FromChallenges(model.db,
		"WHERE shown_at IS NOT NULL "+
			"ORDER BY id")

	lastChallengeByCardId := map[int]db.ChallengeRow{}
	for _, challenge := range challenges {
		lastChallengeByCardId[challenge.CardId] = challenge
	}

	numFailuresByCardId := map[int]int{}
	for _, challenge := range challenges {
		if challenge.Grade.Valid && GRADE_TO_IS_FAILURE[challenge.Grade.String] {
			numFailuresByCardId[challenge.CardId] += 1
		}
	}

	cardRows := db.FromCards(model.db, "")

	cards := model.cardRowsToCardsJoinMorphemes(cardRows)

	skills := []Skill{}
	for _, card := range cards {
		lastChallenge, hasLastChallenge := lastChallengeByCardId[card.Id]
		oneDayAgo := time.Now().UTC().AddDate(0, 0, -1)

		var state string
		if !hasLastChallenge {
			state = "UNTESTED"
		} else if lastChallenge.ShownAt.Valid && !lastChallenge.Grade.Valid {
			state = "NEEDS_GRADE"
		} else if lastChallenge.Grade.String == "RIGHT_WITHOUT_MNEMONIC" &&
			lastChallenge.ShownAt.Time.After(oneDayAgo) {
			state = "RETEST_IN_1D"
		} else if lastChallenge.Grade.String == "BLANK" &&
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
