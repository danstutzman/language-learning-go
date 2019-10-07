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

	allL2, ok := r.URL.Query()["all_l2"]
	if ok && allL2[0] != "" {
		words := []string{}
		for _, word := range L2_WORD_REGEXP.FindAllString(allL2[0], -1) {
			words = append(words, strings.ToLower(word))
		}

		if len(words) > 0 {
			likes := []string{}
			for _, word := range words {
				like := db.Escape(word) + " LIKE (RTRIM(l2, '-') || '%')"
				likes = append(likes, like)
			}

			if where == "" {
				where += "WHERE "
			} else {
				where += " AND "
			}
			where += strings.Join(likes, " OR ")
		}
	}

	response := ListMorphemesResponse{
		Morphemes: db.FromMorphemes(api.db, where+" limit 20"),
	}
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

	w.WriteHeader(204) // send the headers with a 204 response code.
}
