package api

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type ListMorphemesResponse struct {
	Morphemes []model.Morpheme `json:"morphemes"`
}

func (api *Api) HandleListMorphemesRequest(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	var morphemes []model.Morpheme
	l2Phrase := r.URL.Query()["l2_phrase"]
	l2Prefix := r.URL.Query()["l2_prefix"]
	if len(l2Phrase) == 1 && l2Phrase[0] != "" {
		morphemes = []model.Morpheme{}
		words := api.model.SplitL2PhraseIntoWords(l2Phrase[0])
		for _, word := range words {
			morphemes = append(morphemes, api.model.ParseL2WordIntoMorphemes(word)...)
		}
	} else if len(l2Prefix) == 1 && l2Prefix[0] != "" {
		morphemes = api.model.ListMorphemesForPrefix(l2Prefix[0], 20)
	} else {
		morphemes = api.model.ListMorphemesForPrefix("", 20)
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

	morpheme := api.model.MaybeFindMorphemeById(morphemeId)
	if morpheme == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
		return
	}

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

	var morpheme model.Morpheme
	err = json.Unmarshal(body, &morpheme)
	if err != nil {
		panic(err)
	}

	savedMorpheme := api.model.UpsertMorpheme(morpheme)

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

	var morpheme model.Morpheme
	err = json.Unmarshal(body, &morpheme)
	if err != nil {
		panic(err)
	}

	bytes, err := json.Marshal(morpheme)
	if err != nil {
		panic(err)
	}
	w.Write(bytes)
}

func (api *Api) HandleDeleteMorphemeRequest(w http.ResponseWriter, r *http.Request, id int) {
	setCORSHeaders(w)

	api.model.DeleteMorpheme(id)

	w.WriteHeader(http.StatusNoContent)
}
