package model

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"fmt"
	"gopkg.in/guregu/null.v3"
	"regexp"
	"strings"
)

var L2_WORD_REGEXP = regexp.MustCompile(`(?i)[a-zñáéíóúü]+`)

type Morpheme struct {
	Id             int         `json:"id"`
	Type           string      `json:"type"`
	L2             string      `json:"l2"`
	Lemma          null.String `json:"lemma"`
	Tag            null.String `json:"tag"`
	NonsenseL2     null.String `json:"nonsenseL2"`
	NonsensePhones null.String `json:"nonsensePhones"`
}

func morphemeToMorphemeRow(morpheme Morpheme) db.MorphemeRow {
	return db.MorphemeRow{
		Id:             morpheme.Id,
		Type:           morpheme.Type,
		L2:             morpheme.L2,
		Lemma:          morpheme.Lemma,
		Tag:            morpheme.Tag,
		NonsenseL2:     morpheme.NonsenseL2,
		NonsensePhones: morpheme.NonsensePhones,
	}
}

func morphemeRowToMorpheme(row db.MorphemeRow) Morpheme {
	return Morpheme{
		Id:             row.Id,
		Type:           row.Type,
		L2:             row.L2,
		Lemma:          row.Lemma,
		Tag:            row.Tag,
		NonsenseL2:     row.NonsenseL2,
		NonsensePhones: row.NonsensePhones,
	}
}

func morphemeRowPtrToMorphemePtr(row *db.MorphemeRow) *Morpheme {
	if row == nil {
		return nil
	}

	morpheme := morphemeRowToMorpheme(*row)
	return &morpheme
}

func morphemeRowsToMorphemes(rows []db.MorphemeRow) []Morpheme {
	morphemes := []Morpheme{}
	for _, row := range rows {
		morphemes = append(morphemes, morphemeRowToMorpheme(row))
	}
	return morphemes
}

func (model *Model) SplitL2PhraseIntoWords(l2Phrase string) []string {
	return L2_WORD_REGEXP.FindAllString(strings.ToLower(l2Phrase), -1)
}

func (model *Model) ParseL2WordIntoMorphemes(word string) []Morpheme {
	morphemes := []Morpheme{}
	exactMatches := db.FromMorphemes(model.db, "WHERE l2 = "+db.Escape(word))
	if len(exactMatches) > 0 {
		for _, row := range exactMatches {
			morphemes = append(morphemes, morphemeRowToMorpheme(row))
		}
	} else {
		// look for matches with two morphemes
		prefixMatches := db.FromMorphemes(model.db,
			"WHERE "+db.Escape(word)+" LIKE (RTRIM(l2, '-') || '%')")
		for _, prefixMatch := range prefixMatches {

			var suffix string
			if strings.HasSuffix(prefixMatch.L2, "-") {
				suffix = "-" + word[(len(prefixMatch.L2)-1):]
			} else {
				suffix = "-" + word[len(prefixMatch.L2):]
			}

			suffixMatches := db.FromMorphemes(model.db, "WHERE l2 = "+db.Escape(suffix))
			if len(suffixMatches) > 0 {
				morphemes = append(morphemes, morphemeRowToMorpheme(prefixMatch))
				for _, row := range suffixMatches {
					morphemes = append(morphemes, morphemeRowToMorpheme(row))
				}
			}
		}
	}
	return morphemes
}

func (model *Model) ListMorphemesForPrefix(l2Prefix string, limit int) []Morpheme {
	where := ""
	if l2Prefix != "" {
		if where == "" {
			where += "WHERE "
		} else {
			where += " AND "
		}
		where += "l2 LIKE " + db.Escape(l2Prefix+"%")
	}

	rows := db.FromMorphemes(model.db, where+fmt.Sprintf(" LIMIT %d", limit))

	morphemes := []Morpheme{}
	for _, row := range rows {
		morphemes = append(morphemes, morphemeRowToMorpheme(row))
	}
	return morphemes
}

func (model *Model) MaybeFindMorphemeById(id int) *Morpheme {
	where := fmt.Sprintf("WHERE id = %d", id)
	morphemes := db.FromMorphemes(model.db, where)

	if len(morphemes) == 0 {
		return nil
	} else if len(morphemes) > 1 {
		panic("Too many morphemes")
	}
	morpheme := morphemeRowToMorpheme(morphemes[0])
	return &morpheme
}

func (model *Model) InsertMorpheme(morpheme Morpheme) Morpheme {
	return morphemeRowToMorpheme(db.InsertMorpheme(model.db, morphemeToMorphemeRow(morpheme)))
}

func (model *Model) UpsertMorpheme(morpheme Morpheme) Morpheme {
	existingMorphemes := db.FromMorphemes(model.db,
		"WHERE type="+db.Escape(morpheme.Type)+
			" AND l2="+db.Escape(morpheme.L2))
	if len(existingMorphemes) == 0 {
		return morphemeRowToMorpheme(db.InsertMorpheme(model.db, morphemeToMorphemeRow(morpheme)))
	} else {
		return morphemeRowToMorpheme(existingMorphemes[0])
	}
}

func (model *Model) UpdateMorpheme(morpheme Morpheme) {
	db.UpdateMorpheme(model.db, morphemeToMorphemeRow(morpheme))
}

func (model *Model) DeleteMorpheme(id int) {
	where := fmt.Sprintf("WHERE id=%d", id)
	db.DeleteFromMorphemes(model.db, where)
}
