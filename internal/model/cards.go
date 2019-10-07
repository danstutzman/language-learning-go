package model

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"fmt"
)

type Card struct {
	Id        int        `json:"id"`
	L1        string     `json:"l1"`
	L2        string     `json:"l2"`
	Morphemes []Morpheme `json:"morphemes"`
}

type CardList struct {
	Cards []Card `json:"cards"`
}

func (model *Model) GetCard(id int) *Card {
	cardRows := db.FromCards(model.db, fmt.Sprintf("WHERE id=%d", id))
	if len(cardRows) == 0 {
		return nil
	} else if len(cardRows) > 1 {
		panic("Too many cards")
	}
	cardRow := cardRows[0]

	cardsMorphemes := db.FromCardsMorphemes(model.db, fmt.Sprintf("WHERE card_id=%d", cardRow.Id))

	morphemeIds := []int{}
	for _, cardsMorphemes := range cardsMorphemes {
		morphemeIds = append(morphemeIds, cardsMorphemes.MorphemeId)
	}

	morphemes := morphemeRowsToMorphemes(
		db.FromMorphemes(model.db, "WHERE "+db.InIntList("id", morphemeIds)))

	card := Card{
		Id:        cardRow.Id,
		L1:        cardRow.L1,
		L2:        cardRow.L2,
		Morphemes: morphemes,
	}
	return &card
}

func (model *Model) ListCards() CardList {
	cardRows := db.FromCards(model.db, "")

	cardIds := []int{}
	for _, cardRow := range cardRows {
		cardIds = append(cardIds, cardRow.Id)
	}

	cardsMorphemes := db.FromCardsMorphemes(model.db, "WHERE "+db.InIntList("card_id", cardIds))

	allMorphemeIds := []int{}
	morphemeIdsByCardId := map[int][]int{}
	for _, cardsMorphemes := range cardsMorphemes {
		cardId := cardsMorphemes.CardId
		morphemeId := cardsMorphemes.MorphemeId

		allMorphemeIds = append(allMorphemeIds, morphemeId)
		morphemeIdsByCardId[cardId] = append(morphemeIdsByCardId[cardId], morphemeId)
	}

	allMorphemeRows := db.FromMorphemes(model.db,
		"WHERE "+db.InIntList("id", allMorphemeIds))

	morphemeById := map[int]Morpheme{}
	for _, morpheme := range allMorphemeRows {
		morphemeById[morpheme.Id] = morphemeRowToMorpheme(morpheme)
	}

	cards := []Card{}
	for _, cardRow := range cardRows {
		morphemes := []Morpheme{}
		for _, morphemeId := range morphemeIdsByCardId[cardRow.Id] {
			morphemes = append(morphemes, morphemeById[morphemeId])
		}

		card := Card{
			Id:        cardRow.Id,
			L1:        cardRow.L1,
			L2:        cardRow.L2,
			Morphemes: morphemes,
		}
		cards = append(cards, card)
	}

	return CardList{Cards: cards}
}

func (model *Model) InsertCard(card Card) Card {
	savedCardRow := db.InsertCard(model.db, db.CardRow{
		L1: card.L1,
		L2: card.L2,
	})

	card.Id = savedCardRow.Id
	card.L1 = savedCardRow.L1
	card.L2 = savedCardRow.L2
	card.Morphemes = model.saveMorphemes(card)

	return card
}

func (model *Model) UpdateCard(card Card) Card {
	db.UpdateCard(model.db, &db.CardRow{
		Id: card.Id,
		L1: card.L1,
		L2: card.L2,
	})

	card.Morphemes = model.saveMorphemes(card)
	return card
}

func (model *Model) saveMorphemes(card Card) []Morpheme {
	db.DeleteFromCardsMorphemes(model.db, fmt.Sprintf("WHERE card_id=%d", card.Id))

	morphemeL2s := []string{}
	for _, morpheme := range card.Morphemes {
		if morpheme.L2 != "" {
			morphemeL2s = append(morphemeL2s, morpheme.L2)
		}
	}

	savedMorphemes := []Morpheme{}
	existingMorphemeRows := db.FromMorphemes(model.db, "WHERE "+db.InStringList("l2", morphemeL2s))
	for numMorpheme, desiredMorpheme := range card.Morphemes {
		if desiredMorpheme.L2 == "" && desiredMorpheme.Gloss == "" {
			continue
		}

		var savedMorpheme *Morpheme
		for _, existingMorphemeRow := range existingMorphemeRows {
			if existingMorphemeRow.L2 == desiredMorpheme.L2 &&
				existingMorphemeRow.Gloss == desiredMorpheme.Gloss {
				converted := morphemeRowToMorpheme(existingMorphemeRow)
				savedMorpheme = &converted
				break
			}
		}

		if savedMorpheme == nil {
			insertedMorpheme := morphemeRowToMorpheme(
				db.InsertMorpheme(model.db, morphemeToMorphemeRow(desiredMorpheme)))
			savedMorpheme = &insertedMorpheme
		}

		db.InsertCardsMorphemesRow(model.db, db.CardsMorphemesRow{
			CardId:      card.Id,
			MorphemeId:  savedMorpheme.Id,
			NumMorpheme: numMorpheme,
		})

		savedMorphemes = append(savedMorphemes, *savedMorpheme)
	}
	return savedMorphemes
}

func (model *Model) DeleteCardWithId(id int) {
	where := fmt.Sprintf("WHERE id=%d", id)
	db.DeleteFromCards(model.db, where)
}
