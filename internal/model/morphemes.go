package model

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"database/sql"
	"fmt"
	"regexp"
	"strings"
)

var L2_WORD_REGEXP = regexp.MustCompile(`(?i)[a-zñáéíóúü]+`)

type Morpheme struct {
	Id          int     `json:"id"`
	Type        string  `json:"type"`
	L2          string  `json:"l2"`
	Lemma       *string `json:"lemma"`
	FreelingTag *string `json:"freeling_tag"`
}

func stringPtrToNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{Valid: false, String: ""}
	} else {
		return sql.NullString{Valid: true, String: *s}
	}
}

func morphemeToMorphemeRow(morpheme Morpheme) db.MorphemeRow {
	return db.MorphemeRow{
		Id:          morpheme.Id,
		Type:        morpheme.Type,
		L2:          morpheme.L2,
		Lemma:       stringPtrToNullString(morpheme.Lemma),
		FreelingTag: stringPtrToNullString(morpheme.FreelingTag),
	}
}

func nullStringToStringPtr(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	} else {
		return nil
	}
}

func morphemeRowToMorpheme(row db.MorphemeRow) Morpheme {
	return Morpheme{
		Id:          row.Id,
		Type:        row.Type,
		L2:          row.L2,
		Lemma:       nullStringToStringPtr(row.Lemma),
		FreelingTag: nullStringToStringPtr(row.FreelingTag),
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

func (model *Model) findVerbSuffix(l2, verbCategory, tag string) *Morpheme {
	where := fmt.Sprintf(`WHERE type='VERB_SUFFIX' 
		AND l2=%s AND verb_category=%s AND freeling_tag=%s`,
		db.Escape(l2), db.Escape(verbCategory), db.Escape(tag))
	return morphemeRowPtrToMorphemePtr(db.OneFromMorphemes(model.db, where))
}

func (model *Model) findVerbStemChange(lemma, tense string) *Morpheme {
	where := fmt.Sprintf(
		"WHERE type='VERB_STEM_CHANGE' AND lemma=%s AND tense=%s",
		db.Escape(lemma), db.Escape(tense))
	return morphemeRowPtrToMorphemePtr(db.OneFromMorphemes(model.db, where))
}

func (model *Model) findVerbUnique(l2, lemma, tag string) *Morpheme {
	where := fmt.Sprintf(
		"WHERE type='VERB_UNIQUE' AND l2=%s AND lemma=%s AND freeling_tag=%s",
		db.Escape(l2), db.Escape(lemma), db.Escape(tag))
	return morphemeRowPtrToMorphemePtr(db.OneFromMorphemes(model.db, where))
}

func (model *Model) verbToMorphemes(token Token) ([]Morpheme, error) {
	lemma := token.Lemma
	form := strings.ToLower(token.Form)
	tag := token.Tag

	unique := model.findVerbUnique(form, lemma, tag)
	if unique != nil {
		return []Morpheme{*unique}, nil
	}

	var category string
	if strings.HasSuffix(lemma, "ar") {
		category = "ar"
	} else if strings.HasSuffix(lemma, "er") {
		category = "er"
	} else if strings.HasSuffix(lemma, "ir") {
		category = "ir"
	} else {
		return []Morpheme{}, fmt.Errorf("Unknown category for lemma '%s'", lemma)
	}

	stemChangeMorpheme := model.findVerbStemChange(lemma, token.Tense)
	if stemChangeMorpheme != nil {
		suffix := "-" + form[len(stemChangeMorpheme.L2)-1:len(form)]

		category = "stempret"
		suffixMorpheme := model.findVerbSuffix(suffix, category, tag)
		if suffixMorpheme == nil {
			return []Morpheme{}, fmt.Errorf(
				"Can't find verb suffix '%s' with category=%s tag=%s",
				suffix, category, tag)
		}

		return []Morpheme{*stemChangeMorpheme, *suffixMorpheme}, nil
	} else { // If there is no stem change
		stem := lemma[0 : len(lemma)-len(category)]

		if !strings.HasPrefix(form, stem) {
			return []Morpheme{}, fmt.Errorf(
				"No stem change to explain why '%s' doesn't match lemma '%s'",
				form, lemma)
		}

		stemMorpheme := model.UpsertMorpheme(Morpheme{
			Type: "VERB_STEM",
			L2:   stem + "-",
		})

		// Warning: for verbs like 'tengo' the suffix could be weird like 'go'.
		// This should be caught by the unique verb look up earlier, but otherwise
		// it will just fail on the suffix look up.
		suffix := "-" + form[len(stem):len(form)]

		suffixMorpheme := model.findVerbSuffix(suffix, category, tag)
		if suffixMorpheme == nil {
			return []Morpheme{}, fmt.Errorf(
				"Can't find verb suffix '%s' with category=%s tag=%s",
				suffix, category, tag)
		}

		return []Morpheme{stemMorpheme, *suffixMorpheme}, nil
	}
}

func (model *Model) TokenToMorphemes(token Token) ([]Morpheme, error) {
	if token.IsVerb() {
		return model.verbToMorphemes(token)

	} else {
		var type_ string
		if token.IsAdjective() {
			type_ = "ADJECTIVE"
		} else if token.IsAdverb() {
			type_ = "ADVERB"
		} else if token.IsConjunction() {
			type_ = "CONJUNCTION"
		} else if token.IsInterjection() {
			type_ = "INTERJECTION"
		} else if token.IsNoun() {
			type_ = "NOUN"
		} else if token.IsPronoun() {
			type_ = "PRONOUN"
		} else if token.IsPunctuation() {
			type_ = "PUNCTUATION"
		} else {
			return []Morpheme{}, fmt.Errorf("Unknown token tag %s", token.Tag)
		}

		morpheme := model.UpsertMorpheme(Morpheme{
			Type:        type_,
			L2:          token.Form,
			Lemma:       &token.Lemma,
			FreelingTag: &token.Tag,
		})

		return []Morpheme{morpheme}, nil
	}
}
