package model

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"fmt"
	"gopkg.in/guregu/null.v3"
	"strconv"
	"strings"
)

type Card struct {
	Id         int         `json:"id"`
	IsSentence bool        `json:"isSentence"`
	L2         string      `json:"l2"`
	L1         string      `json:"l1"`
	Mnemonic12 null.String `json:"mnemonic12"`
	Mnemonic21 null.String `json:"mnemonic21"`
	Type       string      `json:"type"`

	Morphemes []CardMorpheme `json:"morphemes"`
	State     string         `json:"state"`
}

type CardMorpheme struct {
	Morpheme
	Begin int `json:"begin"`
}

type CardList struct {
	Cards []Card `json:"cards"`
}

func (model *Model) cardRowToCard(row db.CardRow) Card {
	return Card{
		Id:         row.Id,
		IsSentence: row.IsSentence,
		L2:         row.L2,
		L1:         row.L1,
		Type:       row.Type,
		Mnemonic12: row.Mnemonic12,
		Mnemonic21: row.Mnemonic21,
		Morphemes:  []CardMorpheme{},
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

	morphemeRows := db.FromMorphemes(model.db,
		"WHERE "+db.InIntList("id", morphemeIds))

	morphemeRowById := map[int]db.MorphemeRow{}
	for _, morphemeRow := range morphemeRows {
		morphemeRowById[morphemeRow.Id] = morphemeRow
	}

	var cardMorphemes []CardMorpheme
	for _, cardMorpheme := range cardsMorphemes {
		morpheme := morphemeRowToMorpheme(morphemeRowById[cardMorpheme.MorphemeId])
		cardMorpheme := CardMorpheme{Morpheme: morpheme, Begin: cardMorpheme.Begin}
		cardMorphemes = append(cardMorphemes, cardMorpheme)
	}
	card.Morphemes = cardMorphemes

	return card
}

func cardToCardRow(card Card) db.CardRow {
	return db.CardRow{
		IsSentence: card.IsSentence,
		L2:         card.L2,
		L1:         card.L1,
		Mnemonic12: card.Mnemonic12,
		Mnemonic21: card.Mnemonic21,
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
	cardsMorphemesByCardId := map[int][]db.CardsMorphemesRow{}
	for _, cardsMorphemes := range cardsMorphemes {
		allMorphemeIds = append(allMorphemeIds, cardsMorphemes.MorphemeId)
		cardsMorphemesByCardId[cardsMorphemes.CardId] =
			append(cardsMorphemesByCardId[cardsMorphemes.CardId], cardsMorphemes)
	}

	allMorphemeRows := db.FromMorphemes(model.db,
		"WHERE "+db.InIntList("id", allMorphemeIds))

	morphemeById := map[int]Morpheme{}
	for _, morpheme := range allMorphemeRows {
		morphemeById[morpheme.Id] = morphemeRowToMorpheme(morpheme)
	}

	cards := []Card{}
	for _, cardRow := range cardRows {
		cardMorphemes := []CardMorpheme{}
		for _, cardsMorphemes := range cardsMorphemesByCardId[cardRow.Id] {
			cardMorpheme := CardMorpheme{
				Morpheme: morphemeById[cardsMorphemes.MorphemeId],
				Begin:    cardsMorphemes.Begin,
			}
			cardMorphemes = append(cardMorphemes, cardMorpheme)
		}

		card := Card{
			Id:         cardRow.Id,
			IsSentence: cardRow.IsSentence,
			L2:         cardRow.L2,
			L1:         cardRow.L1,
			Mnemonic12: cardRow.Mnemonic12,
			Mnemonic21: cardRow.Mnemonic21,
			Type:       cardRow.Type,

			Morphemes: cardMorphemes,
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
		return *model.GetCardJoinMorphemes(cardRows[0].Id)
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
