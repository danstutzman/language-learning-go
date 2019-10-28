package mem_model

import (
	dbPkg "bitbucket.org/danstutzman/language-learning-go/internal/db"
	"bitbucket.org/danstutzman/language-learning-go/internal/freeling"
	"bitbucket.org/danstutzman/language-learning-go/internal/parsing"
	"context"
	"database/sql"
	"fmt"
	"gopkg.in/guregu/null.v3"
)

type Morpheme struct {
	Id    int
	Type  string
	L2    string
	Lemma null.String
	Tag   null.String
}

func (memModel *MemModel) getNextMorphemeId() int {
	id := memModel.nextMorphemeId
	memModel.nextMorphemeId += 1
	return id
}

func (memModel *MemModel) TokenToCard(token parsing.Token) (Card, error) {
	if token.IsVerb() {
		conjugations := freeling.AnalyzeVerb(token.Lemma, token.Tag)
		if len(conjugations) == 0 {
			return Card{}, fmt.Errorf("Can't conjugate verb %v", token)
		}
		conjugation := conjugations[0]

		var morphemes []Morpheme
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
			morphemes = []Morpheme{morpheme}
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
					Lemma: null.StringFrom(token.Lemma),
					Tag:   null.StringFrom(token.Tag),
				}
				memModel.morphemes = append(memModel.morphemes, suffixMorpheme)
				memModel.morphemeByL2Tag[conjugation.Suffix+token.Tag] = suffixMorpheme
			}
			morphemes = []Morpheme{stemMorpheme, suffixMorpheme}
		}

		card, exists := memModel.cardByL2[token.Form]
		if !exists {
			card = Card{
				Id:        memModel.getNextCardId(),
				Type:      "VERB",
				L2:        token.Form,
				Morphemes: morphemes,
			}
		}
		return card, nil
	} else {
		var type_ string
		if token.IsAdjective() {
			type_ = "ADJECTIVE"
		} else if token.IsAdverb() {
			type_ = "ADVERB"
		} else if token.IsConjunction() {
			type_ = "CONJUNCTION"
		} else if token.IsDate() {
			type_ = "DATE"
		} else if token.IsDeterminer() {
			type_ = "DETERMINER"
		} else if token.IsInterjection() {
			type_ = "INTERJECTION"
		} else if token.IsNoun() {
			type_ = "NOUN"
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
				Type:      type_,
				L2:        token.Form,
				Morphemes: []Morpheme{morpheme},
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
