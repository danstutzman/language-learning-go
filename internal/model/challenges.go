package model

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"fmt"
	"gopkg.in/guregu/null.v3"
	"sort"
	"strconv"
	"strings"
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

	// Also break down which morphemes were involved, update their last_seen_at
	where := "WHERE id IN (" +
		"SELECT morpheme_id FROM cards_morphemes WHERE card_id = " +
		strconv.Itoa(update.CardId) + ")"
	db.TouchMorphemes(model.db, where)

	challengeRows := db.FromChallenges(model.db,
		"WHERE id="+strconv.Itoa(update.Id))
	return challengeRowToChallenge(challengeRows[0])
}

type CardStaleness struct {
	cardId    int
	staleness time.Duration
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func (model *Model) GetTopChallenge(type_ string) *Challenge {
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

	cardStalenesses := []CardStaleness{}
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
		cardStalenesses = append(cardStalenesses,
			CardStaleness{cardId: card.Id, staleness: cardStaleness})
	}

	sort.Slice(cardStalenesses, func(i, j int) bool {
		return cardStalenesses[i].staleness > cardStalenesses[j].staleness
	})

	if false {
		for _, cardStaleness := range cardStalenesses {
			fmt.Printf("%-50s -> %s\n",
				cardById[cardStaleness.cardId].L2, cardStaleness.staleness)
		}
	}

	topCard := cardById[cardStalenesses[0].cardId]

	newChallenge := model.challengeRowToChallengeJoinCard(
		db.InsertChallenge(model.db, db.ChallengeRow{
			Type:   type_,
			CardId: topCard.Id,
		}))

	return &newChallenge
}
