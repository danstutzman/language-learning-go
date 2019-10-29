package mem_model

import (
	dbPkg "bitbucket.org/danstutzman/language-learning-go/internal/db"
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

type Card struct {
	Id         int
	Type       string
	IsSentence bool
	L1         string
	L2         string
	Morphemes  []CardMorpheme
}

func (memModel *MemModel) getNextCardId() int {
	id := memModel.nextCardId
	memModel.nextCardId += 1
	return id
}

func (memModel *MemModel) InsertCardIfNotExists(card Card) Card {
	if existingCard, exists := memModel.cardByL2[card.L2]; exists {
		return existingCard
	}

	card.Id = memModel.getNextCardId()
	memModel.cards = append(memModel.cards, card)
	memModel.cardByL2[card.L2] = card
	return card
}

func (memModel *MemModel) DumpCards() {
	for _, card := range memModel.cards {
		fmt.Printf("Card: %v\n", card)
	}
}

func (memModel *MemModel) SaveCardsToDb(db *sql.DB) {
	tx, err := db.BeginTx(context.TODO(), &sql.TxOptions{})
	if err != nil {
		panic(err)
	}

	query := fmt.Sprintf("DELETE FROM cards")
	_, err = tx.Exec(query)
	if err != nil {
		panic(err)
	}

	query = fmt.Sprintf("DELETE FROM cards_morphemes")
	_, err = tx.Exec(query)
	if err != nil {
		panic(err)
	}

	for _, card := range memModel.cards {
		morphemeIds := []string{}
		for _, morpheme := range card.Morphemes {
			morphemeIds = append(morphemeIds, strconv.Itoa(morpheme.Id))
		}

		query := fmt.Sprintf(`INSERT INTO cards
      (id, type, l1, l2, is_sentence, morpheme_ids_csv)
      VALUES (%d, %s, %s, %s, %s, '%s')`,
			card.Id,
			dbPkg.Escape(card.Type),
			dbPkg.Escape(card.L1),
			dbPkg.Escape(card.L2),
			dbPkg.EscapeBool(card.IsSentence),
			strings.Join(morphemeIds, ","))

		_, err := tx.Exec(query)
		if err != nil {
			panic(err)
		}

		for numMorpheme, cardMorpheme := range card.Morphemes {
			query = fmt.Sprintf(`INSERT INTO cards_morphemes
				(card_id, morpheme_id, num_morpheme, begin) VALUES (%d, %d, %d, %d)`,
				card.Id, cardMorpheme.Morpheme.Id, numMorpheme+1, cardMorpheme.Begin)

			_, err = tx.Exec(query)
			if err != nil {
				panic(err)
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
