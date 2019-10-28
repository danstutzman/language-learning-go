package model

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"fmt"
	"gopkg.in/guregu/null.v3"
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
