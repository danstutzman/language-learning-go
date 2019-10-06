package api

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ListMorphemesResponse struct {
	Morphemes []db.MorphemeRow `json:"morphemes"`
}

func (api *Api) HandleListMorphemesRequest(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	where := ""
	prefix, ok := r.URL.Query()["prefix"]
	if ok {
		where = "WHERE l2 LIKE " + db.Escape(prefix[0]+"%")
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

	savedMorpheme := db.InsertMorpheme(api.db, morpheme)

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
