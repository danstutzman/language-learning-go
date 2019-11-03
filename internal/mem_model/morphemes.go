package mem_model

import (
	dbPkg "bitbucket.org/danstutzman/language-learning-go/internal/db"
	"bitbucket.org/danstutzman/language-learning-go/internal/freeling"
	"bitbucket.org/danstutzman/language-learning-go/internal/spacy"
	"context"
	"database/sql"
	"fmt"
	"gopkg.in/guregu/null.v3"
	"strings"
)

type Morpheme struct {
	Id    int
	Type  string
	L2    string
	Lemma null.String
	Tag   null.String
}

type CardMorpheme struct {
	Morpheme
	Begin int
}

func (memModel *MemModel) getNextMorphemeId() int {
	id := memModel.nextMorphemeId
	memModel.nextMorphemeId += 1
	return id
}

func (memModel *MemModel) VerbTokenToCard(token spacy.Token) (Card, error) {
	conjugations := freeling.AnalyzeVerb(token.Lemma, token.VerbTag)
	if len(conjugations) == 0 {
		return Card{}, fmt.Errorf("Can't conjugate verb %v", token)
	}
	conjugation := conjugations[0]

	var cardMorphemes []CardMorpheme
	if conjugation.Suffix == "" {
		morpheme, exists := memModel.morphemeByL2Tag[conjugation.Stem+token.VerbTag]
		if !exists {
			morpheme = Morpheme{
				Id:    memModel.getNextMorphemeId(),
				Type:  "VERB_UNIQUE",
				L2:    conjugation.Stem,
				Lemma: null.StringFrom(token.Lemma),
				Tag:   null.StringFrom(token.VerbTag),
			}
			memModel.morphemes = append(memModel.morphemes, morpheme)
			memModel.morphemeByL2Tag[conjugation.Stem+token.VerbTag] = morpheme
		}
		cardMorphemes = []CardMorpheme{{
			Morpheme: morpheme,
			Begin:    token.Idx,
		}}
	} else {
		stemMorpheme, exists := memModel.morphemeByL2Tag[conjugation.Stem+""]
		if !exists {
			stemMorpheme = Morpheme{
				Id:    memModel.getNextMorphemeId(),
				Type:  "VERB_STEM",
				L2:    conjugation.Stem,
				Lemma: null.StringFrom(token.Lemma),
				Tag:   null.StringFrom(token.VerbTag),
			}
			memModel.morphemes = append(memModel.morphemes, stemMorpheme)
			memModel.morphemeByL2Tag[conjugation.Stem+""] = stemMorpheme
		}

		suffixMorpheme, exists :=
			memModel.morphemeByL2Tag[conjugation.Suffix+token.VerbTag]
		if !exists {
			suffixMorpheme = Morpheme{
				Id:    memModel.getNextMorphemeId(),
				Type:  "VERB_SUFFIX",
				L2:    conjugation.Suffix,
				Lemma: null.String{},
				Tag:   null.StringFrom(token.VerbTag),
			}
			memModel.morphemes = append(memModel.morphemes, suffixMorpheme)
			memModel.morphemeByL2Tag[conjugation.Suffix+token.VerbTag] = suffixMorpheme
		}
		cardMorphemes = []CardMorpheme{
			{Morpheme: stemMorpheme, Begin: token.Idx},
			{Morpheme: suffixMorpheme,
				Begin: token.Idx + len([]rune(conjugation.Stem))},
		}
	}

	card, exists := memModel.cardByL2[token.Text]
	if !exists {
		card = Card{
			Id:        memModel.getNextCardId(),
			Type:      "VERB",
			L2:        token.Text,
			Morphemes: cardMorphemes,
		}
		memModel.cards = append(memModel.cards, card)
		memModel.cardByL2[token.Text] = card
	}
	return card, nil
}

func (memModel *MemModel) NounOrAdjectiveTokenToCard(token spacy.Token) (Card,
	error) {
	var gender string
	var number string
	if token.Pos == "ADJ" || token.Pos == "NOUN" {
		gender = token.Features["Gender"] // "Masc", "Fem", or ""
		number = token.Features["Number"] // "Sing" or "Plur" or ""
	}

	text := token.Text

	var stem string
	var suffix string
	if gender == "Masc" {
		if number == "Sing" && strings.HasSuffix(text, "o") {
			stem = text[0 : len(text)-1]
			suffix = "o"
		} else if number == "Plur" && strings.HasSuffix(text, "os") {
			stem = text[0 : len(text)-2]
			suffix = "os"
		} else if number == "Plur" && strings.HasSuffix(text, "es") {
			stem = text[0 : len(text)-2]
			suffix = "es"
		}
	} else if gender == "Fem" {
		if number == "Sing" && strings.HasSuffix(text, "a") {
			stem = text[0 : len(text)-1]
			suffix = "a"
		} else if number == "Plur" && strings.HasSuffix(text, "as") {
			stem = text[0 : len(text)-2]
			suffix = "as"
		}
	} else if gender == "" {
		if number == "Plur" && strings.HasSuffix(text, "os") {
			stem = text[0 : len(text)-2]
			suffix = "os"
		} else if number == "Plur" && strings.HasSuffix(text, "es") {
			stem = text[0 : len(text)-2]
			suffix = "es"
		}
	}

	var cardMorphemes []CardMorpheme
	if stem != "" {
		stemMorpheme, exists := memModel.morphemeByL2Tag[stem+""]
		if !exists {
			stemMorpheme = Morpheme{
				Id:    memModel.getNextMorphemeId(),
				Type:  "NOUN_OR_ADJ_STEM",
				L2:    stem,
				Lemma: null.StringFrom(token.Lemma),
				Tag:   null.String{},
			}
			memModel.morphemes = append(memModel.morphemes, stemMorpheme)
			memModel.morphemeByL2Tag[stem+""] = stemMorpheme
		}

		suffixMorpheme, exists :=
			memModel.morphemeByL2Tag[suffix+token.VerbTag]
		if !exists {
			suffixMorpheme = Morpheme{
				Id:    memModel.getNextMorphemeId(),
				Type:  "NOUN_OR_ADJ_SUFFIX",
				L2:    suffix,
				Lemma: null.String{},
				Tag:   null.StringFrom(token.VerbTag),
			}
			memModel.morphemes = append(memModel.morphemes, suffixMorpheme)
			memModel.morphemeByL2Tag[suffix+token.VerbTag] = suffixMorpheme
		}

		cardMorphemes = []CardMorpheme{
			{Morpheme: stemMorpheme, Begin: token.Idx},
			{Morpheme: suffixMorpheme, Begin: token.Idx + len([]rune(stem))},
		}
	} else {
		morpheme, exists := memModel.morphemeByL2Tag[text+token.VerbTag]
		if !exists {
			morpheme = Morpheme{
				Id:    memModel.getNextMorphemeId(),
				Type:  "NOUN_OR_ADJECTIVE",
				L2:    text,
				Lemma: null.StringFrom(token.Lemma),
				Tag:   null.StringFrom(token.VerbTag),
			}
			memModel.morphemes = append(memModel.morphemes, morpheme)
			memModel.morphemeByL2Tag[text+token.VerbTag] = morpheme
		}
		cardMorphemes = []CardMorpheme{{Morpheme: morpheme, Begin: token.Idx}}
	}

	card, exists := memModel.cardByL2[text]
	if !exists {
		card = Card{
			Id:        memModel.getNextCardId(),
			Type:      "NOUN_OR_ADJECTIVE",
			L2:        token.Text,
			Morphemes: cardMorphemes,
		}
		memModel.cards = append(memModel.cards, card)
		memModel.cardByL2[token.Text] = card
	}
	return card, nil
}

func (memModel *MemModel) TokenToCard(token spacy.Token) (Card, error) {
	if token.Pos == "VERB" {
		return memModel.VerbTokenToCard(token)
	} else if token.Pos == "NOUN" || token.Pos == "ADJ" {
		return memModel.NounOrAdjectiveTokenToCard(token)
	} else {
		morpheme, exists := memModel.morphemeByL2Tag[token.Text+token.VerbTag]
		if !exists {
			morpheme = Morpheme{
				Id:    memModel.getNextMorphemeId(),
				Type:  token.Pos,
				L2:    token.Text,
				Lemma: null.StringFrom(token.Lemma),
				Tag:   null.StringFrom(token.VerbTag),
			}
			memModel.morphemes = append(memModel.morphemes, morpheme)
			memModel.morphemeByL2Tag[token.Text+token.VerbTag] = morpheme
		}

		card, exists := memModel.cardByL2[token.Text]
		if !exists {
			card = Card{
				Id:        memModel.getNextCardId(),
				Type:      token.Pos,
				L2:        token.Text,
				Morphemes: []CardMorpheme{{Morpheme: morpheme, Begin: token.Idx}},
			}
			memModel.cards = append(memModel.cards, card)
			memModel.cardByL2[token.Text] = card
		}
		return card, nil
	}
}

func (memModel *MemModel) SaveMorphemesToDb(db *sql.DB) {
	tx, err := db.BeginTx(context.TODO(), &sql.TxOptions{})
	if err != nil {
		panic(err)
	}

	query := fmt.Sprintf("DELETE FROM morphemes")
	_, err = tx.Exec(query)
	if err != nil {
		panic(err)
	}

	for _, morpheme := range memModel.morphemes {
		query := fmt.Sprintf(`INSERT INTO morphemes
			(id, type, l2, lemma, tag)
			VALUES (%d, %s, %s, %s, %s)`, morpheme.Id,
			dbPkg.Escape(morpheme.Type),
			dbPkg.Escape(morpheme.L2),
			dbPkg.EscapeNullString(morpheme.Lemma),
			dbPkg.EscapeNullString(morpheme.Tag))

		_, err := tx.Exec(query)
		if err != nil {
			panic(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
