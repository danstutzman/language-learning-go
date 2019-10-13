package model

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"fmt"
	"strconv"
	"strings"
)

type Card struct {
	Id         int    `json:"id"`
	L1         string `json:"l1"`
	L2         string `json:"l2"`
	NounGender string `json:"nounGender"`
	Type       string `json:"type"`

	Morphemes []Morpheme `json:"morphemes"`
}

type CardList struct {
	Cards []Card `json:"cards"`
}

func (model *Model) cardRowToCard(row db.CardRow) Card {
	return Card{
		Id:         row.Id,
		L1:         row.L1,
		L2:         row.L2,
		NounGender: row.NounGender,
		Type:       row.Type,
		Morphemes:  []Morpheme{},
	}
}

func (model *Model) cardRowToCardJoinMorphemes(row db.CardRow) Card {
	card := model.cardRowToCard(row)

	cardsMorphemes := db.FromCardsMorphemes(model.db,
		fmt.Sprintf("WHERE card_id=%d", row.Id))

	morphemeIds := []int{}
	for _, cardsMorphemes := range cardsMorphemes {
		morphemeIds = append(morphemeIds, cardsMorphemes.MorphemeId)
	}

	card.Morphemes = morphemeRowsToMorphemes(
		db.FromMorphemes(model.db, "WHERE "+db.InIntList("id", morphemeIds)))

	return card
}

func cardToCardRow(card Card) db.CardRow {
	return db.CardRow{
		L1:         card.L1,
		L2:         card.L2,
		NounGender: card.NounGender,
		Type:       card.Type,

		MorphemeIdsCsv: joinMorphemeIdsCsv(card),
	}
}

func (model *Model) GetCardJoinMorphemes(id int) *Card {
	cardRows := db.FromCards(model.db, fmt.Sprintf("WHERE id=%d", id))
	if len(cardRows) == 0 {
		return nil
	} else if len(cardRows) > 1 {
		panic("Too many cards")
	}
	cardRow := cardRows[0]
	card := model.cardRowToCardJoinMorphemes(cardRow)

	return &card
}

func (model *Model) cardRowsToCardsJoinMorphemes(cardRows []db.CardRow) []Card {
	cardIds := []int{}
	for _, cardRow := range cardRows {
		cardIds = append(cardIds, cardRow.Id)
	}

	cardsMorphemes := db.FromCardsMorphemes(model.db,
		"WHERE "+db.InIntList("card_id", cardIds))

	allMorphemeIds := []int{}
	morphemeIdsByCardId := map[int][]int{}
	for _, cardsMorphemes := range cardsMorphemes {
		cardId := cardsMorphemes.CardId
		morphemeId := cardsMorphemes.MorphemeId

		allMorphemeIds = append(allMorphemeIds, morphemeId)
		morphemeIdsByCardId[cardId] =
			append(morphemeIdsByCardId[cardId], morphemeId)
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
			Id:         cardRow.Id,
			L1:         cardRow.L1,
			L2:         cardRow.L2,
			NounGender: cardRow.NounGender,
			Type:       cardRow.Type,

			Morphemes: morphemes,
		}
		cards = append(cards, card)
	}
	return cards
}

func (model *Model) ListCardsJoinMorphemes(whereOrderLimit string) CardList {
	cardRows := db.FromCards(model.db, whereOrderLimit)

	cards := model.cardRowsToCardsJoinMorphemes(cardRows)

	return CardList{Cards: cards}
}

func (model *Model) InsertCardIfNotExists(card Card) Card {
	cardRows := db.FromCards(model.db,
		"WHERE morpheme_ids_csv="+db.Escape(joinMorphemeIdsCsv(card)))
	if len(cardRows) == 1 {
		return model.cardRowToCard(cardRows[0])
	} else {
		return model.InsertCard(card)
	}
}

func joinMorphemeIdsCsv(card Card) string {
	morphemeIds := []string{}
	for _, morpheme := range card.Morphemes {
		morphemeIds = append(morphemeIds, strconv.Itoa(morpheme.Id))
	}

	return strings.Join(morphemeIds, ",")
}

func (model *Model) InsertCard(card Card) Card {
	savedCardRow := db.InsertCard(model.db, cardToCardRow(card))

	card.Id = savedCardRow.Id

	model.saveCardsMorphemes(card)

	return card
}

func (model *Model) UpdateCard(card Card) Card {
	row := cardToCardRow(card)
	db.UpdateCard(model.db, &row)

	model.saveCardsMorphemes(card)

	return card
}

func (model *Model) saveCardsMorphemes(card Card) {
	db.DeleteFromCardsMorphemes(model.db,
		fmt.Sprintf("WHERE card_id=%d", card.Id))

	for numMorpheme, morpheme := range card.Morphemes {
		db.InsertCardsMorphemesRow(model.db, db.CardsMorphemesRow{
			CardId:      card.Id,
			MorphemeId:  morpheme.Id,
			NumMorpheme: numMorpheme,
		})
	}
}

func (model *Model) DeleteCardWithId(id int) {
	where := fmt.Sprintf("WHERE id=%d", id)
	db.DeleteFromCards(model.db, where)
}
