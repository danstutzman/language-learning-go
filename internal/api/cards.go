package api

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type CardList struct {
	Cards []Card `json:"cards"`
}

type Card struct {
	Id        int              `json:"id"`
	L1        string           `json:"l1"`
	L2        string           `json:"l2"`
	Morphemes []db.MorphemeRow `json:"morphemes"`
}

func (api *Api) HandleListCardsRequest(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	cardRows := db.FromCards(api.db, "")

	cardIds := []int{}
	for _, cardRow := range cardRows {
		cardIds = append(cardIds, cardRow.Id)
	}

	cardsMorphemes := db.FromCardsMorphemes(api.db, "WHERE "+db.InIntList("card_id", cardIds))

	allMorphemeIds := []int{}
	morphemeIdsByCardId := map[int][]int{}
	for _, cardsMorphemes := range cardsMorphemes {
		cardId := cardsMorphemes.CardId
		morphemeId := cardsMorphemes.MorphemeId

		allMorphemeIds = append(allMorphemeIds, morphemeId)
		morphemeIdsByCardId[cardId] = append(morphemeIdsByCardId[cardId], morphemeId)
	}

	allMorphemes := db.FromMorphemes(api.db,
		"WHERE "+db.InIntList("id", allMorphemeIds))

	morphemeById := map[int]db.MorphemeRow{}
	for _, morpheme := range allMorphemes {
		morphemeById[morpheme.Id] = morpheme
	}

	cards := []Card{}
	for _, cardRow := range cardRows {
		morphemes := []db.MorphemeRow{}
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

	response := CardList{Cards: cards}
	bytes, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error from json.Marshal: %s", err)
	}
	w.Write(bytes)
}

func (api *Api) HandleShowCardRequest(w http.ResponseWriter, r *http.Request,
	cardId int) {
	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	cardRows := db.FromCards(api.db, fmt.Sprintf("WHERE id=%d", cardId))
	if len(cardRows) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
		return
	} else if len(cardRows) > 1 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		panic("Too many cards")
	}
	cardRow := cardRows[0]

	cardsMorphemes := db.FromCardsMorphemes(api.db, fmt.Sprintf("WHERE card_id=%d", cardRow.Id))

	morphemeIds := []int{}
	for _, cardsMorphemes := range cardsMorphemes {
		morphemeIds = append(morphemeIds, cardsMorphemes.MorphemeId)
	}

	morphemes := db.FromMorphemes(api.db, "WHERE "+db.InIntList("id", morphemeIds))

	card := Card{
		Id:        cardRow.Id,
		L1:        cardRow.L1,
		L2:        cardRow.L2,
		Morphemes: morphemes,
	}

	bytes, err := json.Marshal(card)
	if err != nil {
		log.Fatalf("Error from json.Marshal: %s", err)
	}
	w.Write(bytes)
}

func (api *Api) saveMorphemes(card Card) []db.MorphemeRow {
	db.DeleteFromCardsMorphemes(api.db, fmt.Sprintf("WHERE card_id=%d", card.Id))

	morphemeL2s := []string{}
	for _, morpheme := range card.Morphemes {
		if morpheme.L2 != "" {
			morphemeL2s = append(morphemeL2s, morpheme.L2)
		}
	}

	savedMorphemes := []db.MorphemeRow{}
	existingMorphemes := db.FromMorphemes(api.db, "WHERE "+db.InStringList("l2", morphemeL2s))
	for numMorpheme, desiredMorpheme := range card.Morphemes {
		if desiredMorpheme.L2 == "" && desiredMorpheme.Gloss == "" {
			continue
		}

		var savedMorpheme *db.MorphemeRow
		for _, existingMorpheme := range existingMorphemes {
			if existingMorpheme.L2 == desiredMorpheme.L2 &&
				existingMorpheme.Gloss == desiredMorpheme.Gloss {
				savedMorpheme = &existingMorpheme
				break
			}
		}

		if savedMorpheme == nil {
			insertedMorpheme := db.InsertMorpheme(api.db, desiredMorpheme)
			savedMorpheme = &insertedMorpheme
		}

		db.InsertCardsMorphemesRow(api.db, db.CardsMorphemesRow{
			CardId:      card.Id,
			MorphemeId:  savedMorpheme.Id,
			NumMorpheme: numMorpheme,
		})

		savedMorphemes = append(savedMorphemes, *savedMorpheme)
	}
	return savedMorphemes
}

func (api *Api) HandleCreateCardRequest(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var unsavedCard Card
	err = json.Unmarshal(body, &unsavedCard)
	if err != nil {
		panic(err)
	}

	savedCardRow := db.InsertCard(api.db, db.CardRow{
		L1: unsavedCard.L1,
		L2: unsavedCard.L2,
	})

	savedCard := Card{
		Id: savedCardRow.Id,
		L1: savedCardRow.L1,
		L2: savedCardRow.L2,
	}
	savedCard.Morphemes = api.saveMorphemes(savedCard)

	bytes, err := json.Marshal(savedCard)
	if err != nil {
		log.Fatalf("Error from json.Marshal: %s", err)
	}
	w.Write(bytes)
}

func (api *Api) HandleUpdateCardRequest(w http.ResponseWriter, r *http.Request,
	cardId int) {
	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var card Card
	err = json.Unmarshal(body, &card)
	if err != nil {
		panic(err)
	}

	db.UpdateCard(api.db, &db.CardRow{
		Id: card.Id,
		L1: card.L1,
		L2: card.L2,
	})

	card.Morphemes = api.saveMorphemes(card)

	bytes, err := json.Marshal(card)
	if err != nil {
		log.Fatalf("Error from json.Marshal: %s", err)
	}
	w.Write(bytes)
}

func (api *Api) HandleDeleteCardRequest(w http.ResponseWriter, r *http.Request, id int) {
	setCORSHeaders(w)

	where := fmt.Sprintf("WHERE id=%d", id)
	db.DeleteFromCards(api.db, where)

	w.WriteHeader(http.StatusNoContent)
}
