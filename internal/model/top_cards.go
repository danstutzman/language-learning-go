package model

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"strconv"
	"strings"
	"time"
)

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func (model *Model) GetTopCards() CardList {
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

	lastShownAtByMorphemeId := map[int]time.Time{}
	answerMorphemeRows := db.FromAnswerMorphemes(model.db, "")
	for _, answerMorpheme := range answerMorphemeRows {
		morphemeId := answerMorpheme.MorphemeId
		if answerMorpheme.ShownAt.After(lastShownAtByMorphemeId[morphemeId]) {
			lastShownAtByMorphemeId[morphemeId] = answerMorpheme.ShownAt
		}
	}

	cardRows := []db.CardRow{}
	for i := 0; i < 10; i += 1 {
		stalestCardId := 0
		stalestStaleness := time.Duration(0)
		for _, card := range allCards {
			cardStaleness := time.Duration(0)
			for _, morphemeId := range strings.Split(card.MorphemeIdsCsv, ",") {
				morphemeStaleness := now.Sub(
					lastShownAtByMorphemeId[mustAtoi(morphemeId)])
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
		cardRows = append(cardRows, topCard)

		// simulate morpheme staleness being updated so that each card in the
		// return challengeList has a non-overlapping set of morphemes.
		for _, morphemeId := range strings.Split(topCard.MorphemeIdsCsv, ",") {
			lastShownAtByMorphemeId[mustAtoi(morphemeId)] = now
		}
	}

	cards := model.cardRowsToCardsJoinMorphemes(cardRows)
	return CardList{Cards: cards}
}
