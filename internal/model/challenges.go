package model

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"fmt"
	"gopkg.in/guregu/null.v3"
)

type ChallengeList struct {
	Challenges []Challenge `json:"challenges"`
}

type Challenge struct {
	Id             int         `json:"id"`
	Type           string      `json:"type"`
	CardId         int         `json:"cardId"`
	AnsweredL2     null.String `json:"answeredL2"`
	AnsweredAt     null.Time   `json:"answeredAt"`
	ShowedMnemonic null.Bool   `json:"showedMnemonic"`
	Mnemonic       null.String `json:"mnemonic"`

	Card *Card `json:"card"`
}

func challengeToChallengeRow(challenge Challenge) db.ChallengeRow {
	return db.ChallengeRow{
		Id:             challenge.Id,
		Type:           challenge.Type,
		CardId:         challenge.CardId,
		AnsweredL2:     challenge.AnsweredL2,
		AnsweredAt:     challenge.AnsweredAt,
		ShowedMnemonic: challenge.ShowedMnemonic,
		Mnemonic:       challenge.Mnemonic,
	}
}

func challengeRowToChallenge(row db.ChallengeRow) Challenge {
	return Challenge{
		Id:             row.Id,
		Type:           row.Type,
		CardId:         row.CardId,
		AnsweredL2:     row.AnsweredL2,
		AnsweredAt:     row.AnsweredAt,
		ShowedMnemonic: row.ShowedMnemonic,
		Mnemonic:       row.Mnemonic,
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
	challengeRows := db.FromChallenges(model.db, "")

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

func (model *Model) ReplaceChallenge(challenge Challenge) {
	db.DeleteFromChallenges(model.db, fmt.Sprintf("WHERE card_id=%d AND type=%s",
		challenge.CardId, db.Escape(challenge.Type)))

	db.InsertChallenge(model.db, challengeToChallengeRow(challenge))
}

func (model *Model) GetTopGiven1Type2Challenge() *Challenge {
	challengeRows := db.FromChallenges(model.db, "ORDER BY answered_at")
	if len(challengeRows) == 0 {
		return nil
	}
	challenge := model.challengeRowToChallengeJoinCard(challengeRows[0])

	return &challenge
}
