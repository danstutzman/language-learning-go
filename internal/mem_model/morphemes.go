package mem_model

import (
	dbPkg "bitbucket.org/danstutzman/language-learning-go/internal/db"
	"bitbucket.org/danstutzman/language-learning-go/internal/freeling"
	"bitbucket.org/danstutzman/language-learning-go/internal/parsing"
	"context"
	"database/sql"
	"fmt"
	"gopkg.in/guregu/null.v3"
	"strconv"
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

func (memModel *MemModel) VerbTokenToCard(token parsing.Token) (Card, error) {
	conjugations := freeling.AnalyzeVerb(token.Lemma, token.Tag)
	if len(conjugations) == 0 {
		return Card{}, fmt.Errorf("Can't conjugate verb %v", token)
	}
	conjugation := conjugations[0]

	var cardMorphemes []CardMorpheme
	if conjugation.Suffix == "" {
		morpheme, exists := memModel.morphemeByL2Tag[token.Form+token.Tag]
		if !exists {
			morpheme = Morpheme{
				Id:    memModel.getNextMorphemeId(),
				Type:  "VERB_UNIQUE",
				L2:    conjugation.Stem,
				Lemma: null.StringFrom(token.Lemma),
				Tag:   null.StringFrom(token.Tag),
			}
			memModel.morphemes = append(memModel.morphemes, morpheme)
			memModel.morphemeByL2Tag[token.Form+token.Tag] = morpheme
		}
		cardMorphemes = []CardMorpheme{{
			Morpheme: morpheme,
			Begin:    mustAtoi(token.Begin),
		}}
	} else {
		stemMorpheme, exists := memModel.morphemeByL2Tag[conjugation.Stem+""]
		if !exists {
			stemMorpheme = Morpheme{
				Id:    memModel.getNextMorphemeId(),
				Type:  "VERB_STEM",
				L2:    conjugation.Stem,
				Lemma: null.StringFrom(token.Lemma),
				Tag:   null.StringFrom(token.Tag),
			}
			memModel.morphemes = append(memModel.morphemes, stemMorpheme)
			memModel.morphemeByL2Tag[conjugation.Stem+""] = stemMorpheme
		}

		suffixMorpheme, exists :=
			memModel.morphemeByL2Tag[conjugation.Suffix+token.Tag]
		if !exists {
			suffixMorpheme = Morpheme{
				Id:    memModel.getNextMorphemeId(),
				Type:  "VERB_SUFFIX",
				L2:    conjugation.Suffix,
				Lemma: null.String{},
				Tag:   null.StringFrom(token.Tag),
			}
			memModel.morphemes = append(memModel.morphemes, suffixMorpheme)
			memModel.morphemeByL2Tag[conjugation.Suffix+token.Tag] = suffixMorpheme
		}
		cardMorphemes = []CardMorpheme{
			{Morpheme: stemMorpheme, Begin: mustAtoi(token.Begin)},
			{Morpheme: suffixMorpheme,
				Begin: mustAtoi(token.Begin) + len([]rune(conjugation.Stem))},
		}
	}

	card, exists := memModel.cardByL2[token.Form]
	if !exists {
		card = Card{
			Id:        memModel.getNextCardId(),
			Type:      "VERB",
			L2:        token.Form,
			Morphemes: cardMorphemes,
		}
	}
	return card, nil
}

func (memModel *MemModel) NounOrAdjectiveTokenToCard(token parsing.Token) (Card,
	error) {
	var gender string
	var number string
	if token.Tag[0:1] == "A" { // adjective
		gender = token.Tag[3:4] // "F" or "M" or "C" (common)
		number = token.Tag[4:5] // "P" or "S"
	} else if token.Tag[0:1] == "N" { // noun
		gender = token.Tag[2:3] // "F" or "M" or "C" (common)
		number = token.Tag[3:4] // "P" or "S" or "N" (invariable)
	}

	form := token.Form

	var stem string
	var suffix string
	if gender == "M" {
		if number == "S" && strings.HasSuffix(form, "o") {
			stem = form[0 : len(form)-1]
			suffix = "o"
		} else if number == "P" && strings.HasSuffix(form, "os") {
			stem = form[0 : len(form)-2]
			suffix = "os"
		} else if number == "P" && strings.HasSuffix(form, "es") {
			stem = form[0 : len(form)-2]
			suffix = "es"
		}
	} else if gender == "F" {
		if number == "S" && strings.HasSuffix(form, "a") {
			stem = form[0 : len(form)-1]
			suffix = "a"
		} else if number == "P" && strings.HasSuffix(form, "as") {
			stem = form[0 : len(form)-2]
			suffix = "as"
		}
	} else if gender == "C" {
		if number == "P" && strings.HasSuffix(form, "os") {
			stem = form[0 : len(form)-2]
			suffix = "os"
		} else if number == "P" && strings.HasSuffix(form, "es") {
			stem = form[0 : len(form)-2]
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
			memModel.morphemeByL2Tag[suffix+token.Tag]
		if !exists {
			suffixMorpheme = Morpheme{
				Id:    memModel.getNextMorphemeId(),
				Type:  "NOUN_OR_ADJ_SUFFIX",
				L2:    suffix,
				Lemma: null.String{},
				Tag:   null.StringFrom(token.Tag),
			}
			memModel.morphemes = append(memModel.morphemes, suffixMorpheme)
			memModel.morphemeByL2Tag[suffix+token.Tag] = suffixMorpheme
		}

		cardMorphemes = []CardMorpheme{
			{Morpheme: stemMorpheme, Begin: mustAtoi(token.Begin)},
			{Morpheme: suffixMorpheme,
				Begin: mustAtoi(token.Begin) + len([]rune(stem))},
		}
	} else {
		morpheme, exists := memModel.morphemeByL2Tag[form+token.Tag]
		if !exists {
			morpheme = Morpheme{
				Id:    memModel.getNextMorphemeId(),
				Type:  "NOUN_OR_ADJECTIVE",
				L2:    form,
				Lemma: null.StringFrom(token.Lemma),
				Tag:   null.StringFrom(token.Tag),
			}
			memModel.morphemes = append(memModel.morphemes, morpheme)
			memModel.morphemeByL2Tag[form+token.Tag] = morpheme
		}
		cardMorphemes = []CardMorpheme{{
			Morpheme: morpheme, Begin: mustAtoi(token.Begin),
		}}
	}

	card, exists := memModel.cardByL2[form]
	if !exists {
		card = Card{
			Id:        memModel.getNextCardId(),
			Type:      "NOUN_OR_ADJECTIVE",
			L2:        token.Form,
			Morphemes: cardMorphemes,
		}
	}
	return card, nil
}

func (memModel *MemModel) TokenToCard(token parsing.Token) (Card, error) {
	if token.IsVerb() {
		return memModel.VerbTokenToCard(token)
	} else if token.IsNoun() || token.IsAdjective() {
		return memModel.NounOrAdjectiveTokenToCard(token)
	} else {
		var type_ string
		if token.IsAdverb() {
			type_ = "ADVERB"
		} else if token.IsConjunction() {
			type_ = "CONJUNCTION"
		} else if token.IsDate() {
			type_ = "DATE"
		} else if token.IsDeterminer() {
			type_ = "DETERMINER"
		} else if token.IsInterjection() {
			type_ = "INTERJECTION"
		} else if token.IsNumber() {
			type_ = "NUMBER"
		} else if token.IsPreposition() {
			type_ = "PREPOSITION"
		} else if token.IsPronoun() {
			type_ = "PRONOUN"
		} else if token.IsPunctuation() {
			type_ = "PUNCTUATION"
		} else {
			return Card{}, fmt.Errorf("Unknown token tag %s for form '%s'",
				token.Tag, token.Form)
		}

		morpheme, exists := memModel.morphemeByL2Tag[token.Form+token.Tag]
		if !exists {
			morpheme = Morpheme{
				Id:    memModel.getNextMorphemeId(),
				Type:  type_,
				L2:    token.Form,
				Lemma: null.StringFrom(token.Lemma),
				Tag:   null.StringFrom(token.Tag),
			}
			memModel.morphemes = append(memModel.morphemes, morpheme)
			memModel.morphemeByL2Tag[token.Form+token.Tag] = morpheme
		}

		card, exists := memModel.cardByL2[token.Form]
		if !exists {
			card = Card{
				Type: type_,
				L2:   token.Form,
				Morphemes: []CardMorpheme{{
					Morpheme: morpheme,
					Begin:    mustAtoi(token.Begin),
				}},
			}
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

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
