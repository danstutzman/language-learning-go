package model

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"fmt"
	"gopkg.in/guregu/null.v3"
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

	Expectation string    `json:"expectation"`
	HideUntil   time.Time `json:"hideUntil"`

	AnsweredL1     null.String `json:"answeredL1"`
	AnsweredL2     null.String `json:"answeredL2"`
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
		HideUntil:   challenge.HideUntil,

		AnsweredL1:     challenge.AnsweredL1,
		AnsweredL2:     challenge.AnsweredL2,
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
		HideUntil:   row.HideUntil,

		AnsweredL1:     row.AnsweredL1,
		AnsweredL2:     row.AnsweredL2,
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
		"WHERE type='Given2Type1' ORDER BY card_id")

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

// Returns old (updated) challenge, not the new challenge
func (model *Model) UpdateChallengeAndCreateNew(
	update db.ChallengeUpdate) Challenge {

	db.UpdateChallenge(model.db, update)

	challengeRows := db.FromChallenges(model.db,
		"WHERE id="+strconv.Itoa(update.Id))
	oldChallenge := challengeRowToChallenge(challengeRows[0])

	if update.Grade.String == "RIGHT" {
		db.InsertChallenge(model.db, db.ChallengeRow{
			Type:   oldChallenge.Type,
			CardId: oldChallenge.CardId,

			Expectation: oldChallenge.Expectation,
			// HideUntil:   time.Now().UTC().AddDate(0, 0, 1),
			HideUntil: time.Now().UTC().Add(time.Minute * time.Duration(1)),
		})
	}

	return oldChallenge
}

func (model *Model) GetTopChallenge(type_ string) *Challenge {
	where := "WHERE type = " + db.Escape(type_) +
		" AND answered_at IS NULL" +
		" AND STRFTIME('%Y-%m-%dT%H:%M:%SZ', 'now') >= HIDE_UNTIL"
	challengeRows := db.FromChallenges(model.db, where)
	if len(challengeRows) == 0 {
		return nil
	}
	challenge := model.challengeRowToChallengeJoinCard(challengeRows[0])

	return &challenge
}
