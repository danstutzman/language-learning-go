package api

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type ListMorphemesResponse struct {
	Morphemes []db.MorphemeRow `json:"morphemes"`
}

var L2_WORD_REGEXP = regexp.MustCompile(`(?i)[a-zñáéíóúü]+`)

func (api *Api) HandleListMorphemesRequest(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	var morphemes []db.MorphemeRow
	l2Phrase, ok := r.URL.Query()["l2_phrase"]
	if ok && l2Phrase[0] != "" {
		morphemes = []db.MorphemeRow{}
		for _, word := range L2_WORD_REGEXP.FindAllString(strings.ToLower(l2Phrase[0]), -1) {
			exactMatches := db.FromMorphemes(api.db, "WHERE l2 = "+db.Escape(word))
			if len(exactMatches) > 0 {
				morphemes = append(morphemes, exactMatches...)
			} else {
				// look for matches with two morphemes
				prefixMatches := db.FromMorphemes(api.db,
					"WHERE "+db.Escape(word)+" LIKE (RTRIM(l2, '-') || '%')")
				for _, prefixMatch := range prefixMatches {
					// subtract one to account for the prefix's hyphen
					suffix := "-" + word[(len(prefixMatch.L2)-1):]

					suffixMatches := db.FromMorphemes(api.db, "WHERE l2 = "+db.Escape(suffix))
					if len(suffixMatches) > 0 {
						morphemes = append(morphemes, prefixMatch)
						morphemes = append(morphemes, suffixMatches...)
					}
				}
			}
		}
	} else {
		where := ""
		l2Prefix, ok := r.URL.Query()["l2_prefix"]
		if ok && l2Prefix[0] != "" {
			if where == "" {
				where += "WHERE "
			} else {
				where += " AND "
			}
			where += "l2 LIKE " + db.Escape(l2Prefix[0]+"%")
		}

		morphemes = db.FromMorphemes(api.db, where+" limit 20")
	}

	response := ListMorphemesResponse{Morphemes: morphemes}
	bytes, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error from json.Marshal: %s", err)
	}
	w.Write(bytes)
}

func (api *Api) HandleShowMorphemeRequest(w http.ResponseWriter, r *http.Request,
	morphemeId int) {
	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	where := fmt.Sprintf("WHERE id = %d", morphemeId)
	morphemes := db.FromMorphemes(api.db, where)
	if len(morphemes) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
		return
	} else if len(morphemes) > 1 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		panic("Too many morphemes")
	}
	morpheme := morphemes[0]

	bytes, err := json.Marshal(morpheme)
	if err != nil {
		panic(err)
	}
	w.Write(bytes)
}

func (api *Api) HandleCreateMorphemeRequest(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var morpheme db.MorphemeRow
	err = json.Unmarshal(body, &morpheme)
	if err != nil {
		panic(err)
	}

	var savedMorpheme db.MorphemeRow
	existingMorphemes := db.FromMorphemes(api.db, "WHERE l2="+db.Escape(morpheme.L2))
	if len(existingMorphemes) == 0 {
		savedMorpheme = db.InsertMorpheme(api.db, morpheme)
	} else {
		savedMorpheme = existingMorphemes[0]
	}

	bytes, err := json.Marshal(savedMorpheme)
	if err != nil {
		panic(err)
	}
	w.Write(bytes)
}

func (api *Api) HandleUpdateMorphemeRequest(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var morpheme db.MorphemeRow
	err = json.Unmarshal(body, &morpheme)
	if err != nil {
		panic(err)
	}

	db.UpdateMorpheme(api.db, morpheme)

	bytes, err := json.Marshal(morpheme)
	if err != nil {
		panic(err)
	}
	w.Write(bytes)
}

func (api *Api) HandleDeleteMorphemeRequest(w http.ResponseWriter, r *http.Request, id int) {
	setCORSHeaders(w)

	where := fmt.Sprintf("WHERE id=%d", id)
	db.DeleteFromMorphemes(api.db, where)

	w.WriteHeader(http.StatusNoContent)
}
