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

	ShownAt        null.Time   `json:"shownAt"`
	AnsweredL1     null.String `json:"answeredL1"`
	AnsweredL2     null.String `json:"answeredL2"`
	ShowedMnemonic null.Bool   `json:"showedMnemonic"`
	FirstKeyMillis null.Int    `json:"firstKeyMillis"`
	LastKeyMillis  null.Int    `json:"lastKeyMillis"`

	Grade              null.String `json:"grade"`
	MisconnectedCardId null.Int    `json:"misconnectedCardId"`
}

func challengeToChallengeRow(challenge Challenge) db.ChallengeRow {
	return db.ChallengeRow{
		Id:     challenge.Id,
		Type:   challenge.Type,
		CardId: challenge.CardId,

		Expectation: challenge.Expectation,

		ShownAt:        challenge.ShownAt,
		AnsweredL1:     challenge.AnsweredL1,
		AnsweredL2:     challenge.AnsweredL2,
		ShowedMnemonic: challenge.ShowedMnemonic,
		FirstKeyMillis: challenge.FirstKeyMillis,
		LastKeyMillis:  challenge.LastKeyMillis,

		Grade:              challenge.Grade,
		MisconnectedCardId: challenge.MisconnectedCardId,
	}
}

func challengeRowToChallenge(row db.ChallengeRow) Challenge {
	return Challenge{
		Id:     row.Id,
		Type:   row.Type,
		CardId: row.CardId,

		Expectation: row.Expectation,

		ShownAt:        row.ShownAt,
		AnsweredL1:     row.AnsweredL1,
		AnsweredL2:     row.AnsweredL2,
		ShowedMnemonic: row.ShowedMnemonic,
		FirstKeyMillis: row.FirstKeyMillis,
		LastKeyMillis:  row.LastKeyMillis,

		Grade:              row.Grade,
		MisconnectedCardId: row.MisconnectedCardId,
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
	challengeRows := db.FromChallenges(model.db, "ORDER BY card_id, id")

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
		"WHERE type="+db.Escape(type_)+
			" AND shown_at IS NOT NULL "+
			"ORDER BY id")

	lastChallengeByCardId := map[int]db.ChallengeRow{}
	for _, challenge := range challenges {
		lastChallengeByCardId[challenge.CardId] = challenge
	}

	cardsUnfiltered := db.FromCards(model.db, "")

	cards := []db.CardRow{}
	for _, card := range cardsUnfiltered {
		lastChallenge, hasLastChallenge := lastChallengeByCardId[card.Id]
		oneDayAgo := time.Now().UTC().AddDate(0, 0, -1)

		if !hasLastChallenge {
			// Show card if it's never been shown
			cards = append(cards, card)
		} else if lastChallenge.ShownAt.Valid && !lastChallenge.Grade.Valid {
			// Waiting for manual grade, so suspend card for now
		} else if lastChallenge.Grade.String == "RIGHT_WITHOUT_MNEMONIC" &&
			lastChallenge.ShownAt.Time.After(oneDayAgo) {
			// Suspend card if correct within 24 hours
		} else if lastChallenge.Grade.String == "BLANK" &&
			((type_ == "Given1Type2" && card.Mnemonic21.String == "") ||
				(type_ == "Given2Type1" && card.Mnemonic12.String == "")) {
			// Suspend card if drew a blank but have no mnemonic
		} else {
			cards = append(cards, card)
		}
	}

	if len(cards) == 0 {
		return nil
	}

	sort.Slice(cards, func(i, j int) bool {
		return lastChallengeByCardId[cards[i].Id].ShownAt.Time.Before(
			lastChallengeByCardId[cards[j].Id].ShownAt.Time)
	})
	card := cards[0]

	newChallenge := model.challengeRowToChallengeJoinCard(
		db.InsertChallenge(model.db, db.ChallengeRow{
			Type:   type_,
			CardId: card.Id,
		}))

	return &newChallenge
}
