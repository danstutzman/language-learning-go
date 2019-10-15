package model

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"fmt"
	"gopkg.in/guregu/null.v3"
	"sort"
	"strconv"
	"time"
)

type ChallengeList struct {
	Challenges []Challenge `json:"challenges"`
}

type Challenge struct {
	Id     int    `json:"id"`
	Type   string `json:"type"`
	CardId int    `json:"cardId"`
	Card   *Card  `json:"card"`

	Expectation string `json:"expectation"`

	AnsweredL1     null.String `json:"answeredL1"`
	AnsweredAt     null.Time   `json:"answeredAt"`
	ShowedMnemonic null.Bool   `json:"showedMnemonic"`

	Grade null.String `json:"grade"`
}

func challengeToChallengeRow(challenge Challenge) db.ChallengeRow {
	return db.ChallengeRow{
		Id:     challenge.Id,
		Type:   challenge.Type,
		CardId: challenge.CardId,

		Expectation: challenge.Expectation,

		AnsweredL1:     challenge.AnsweredL1,
		AnsweredAt:     challenge.AnsweredAt,
		ShowedMnemonic: challenge.ShowedMnemonic,

		Grade: challenge.Grade,
	}
}

func challengeRowToChallenge(row db.ChallengeRow) Challenge {
	return Challenge{
		Id:     row.Id,
		Type:   row.Type,
		CardId: row.CardId,

		Expectation: row.Expectation,

		AnsweredL1:     row.AnsweredL1,
		AnsweredAt:     row.AnsweredAt,
		ShowedMnemonic: row.ShowedMnemonic,

		Grade: row.Grade,
	}
}

func (model *Model) challengeRowToChallengeJoinCard(
	row db.ChallengeRow) Challenge {

	challenge := challengeRowToChallenge(row)
	cards := db.FromCards(model.db, fmt.Sprintf("WHERE id=%d", row.CardId))
	card := model.cardRowToCard(cards[0])
	challenge.Card = &card
	return challenge
}

func (model *Model) ListChallenges() ChallengeList {
	challengeRows := db.FromChallenges(model.db,
		"WHERE type='Given2Type1' ORDER BY card_id, id")

	challenges := []Challenge{}
	for _, challengeRow := range challengeRows {
		challenges = append(challenges, challengeRowToChallenge(challengeRow))
	}

	cardIds := []int{}
	for _, challenge := range challenges {
		cardIds = append(cardIds, challenge.CardId)
	}

	cardRows := db.FromCards(model.db, "WHERE "+db.InIntList("id", cardIds))

	cardById := map[int]Card{}
	for _, row := range cardRows {
		cardById[row.Id] = model.cardRowToCardJoinMorphemes(row)
	}

	for i, _ := range challenges {
		card := cardById[challenges[i].CardId]
		challenges[i].Card = &card
	}

	return ChallengeList{Challenges: challenges}
}

func (model *Model) InsertChallenge(challenge Challenge) {
	db.InsertChallenge(model.db, challengeToChallengeRow(challenge))
}

func (model *Model) UpdateChallenge(update db.ChallengeUpdate) Challenge {
	db.UpdateChallenge(model.db, update)

	challengeRows := db.FromChallenges(model.db,
		"WHERE id="+strconv.Itoa(update.Id))
	return challengeRowToChallenge(challengeRows[0])
}

func (model *Model) GetTopChallenge(type_ string) *Challenge {
	challenges := db.FromChallenges(model.db,
		"WHERE type="+db.Escape(type_))

	lastAnsweredAtByCardId := map[int]time.Time{}
	for _, challenge := range challenges {
		cardId := challenge.CardId
		answeredAt := challenge.AnsweredAt.Time

		if answeredAt.After(lastAnsweredAtByCardId[cardId]) {
			lastAnsweredAtByCardId[cardId] = answeredAt
		}
	}

	cards := db.FromCards(model.db, "")
	sort.Slice(cards, func(i, j int) bool {
		return lastAnsweredAtByCardId[cards[i].Id].Before(
			lastAnsweredAtByCardId[cards[j].Id])
	})
	card := cards[0]

	newChallenge := model.challengeRowToChallengeJoinCard(
		db.InsertChallenge(model.db, db.ChallengeRow{
			Type:   type_,
			CardId: card.Id,
		}))

	return &newChallenge
}
