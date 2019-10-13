package api

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Answer struct {
	CardId   int    `json:"cardId"`
	L2Answer string `json:"l2Answer"`
}

func (api *Api) HandleListCardsRequest(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	cardList := api.model.ListCards("")

	bytes, err := json.Marshal(cardList)
	if err != nil {
		log.Fatalf("Error from json.Marshal: %s", err)
	}
	w.Write(bytes)
}

func (api *Api) HandleShowCardRequest(w http.ResponseWriter, r *http.Request,
	cardId int) {
	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	card := api.model.GetCard(cardId)
	if card == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
		return
	}

	bytes, err := json.Marshal(card)
	if err != nil {
		log.Fatalf("Error from json.Marshal: %s", err)
	}
	w.Write(bytes)
}

func (api *Api) HandleCreateCardRequest(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var unsavedCard model.Card
	err = json.Unmarshal(body, &unsavedCard)
	if err != nil {
		panic(err)
	}

	/*
		savedCard := api.model.InsertCard(unsavedCard)

		bytes, err := json.Marshal(savedCard)
		if err != nil {
			log.Fatalf("Error from json.Marshal: %s", err)
		}
		w.Write(bytes)
	*/
}

func (api *Api) HandleUpdateCardRequest(w http.ResponseWriter, r *http.Request,
	cardId int) {
	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var unsavedCard model.Card
	err = json.Unmarshal(body, &unsavedCard)
	if err != nil {
		panic(err)
	}

	savedCard := api.model.UpdateCard(unsavedCard)

	bytes, err := json.Marshal(savedCard)
	if err != nil {
		log.Fatalf("Error from json.Marshal: %s", err)
	}
	w.Write(bytes)
}

func (api *Api) HandleDeleteCardRequest(w http.ResponseWriter, r *http.Request, id int) {
	setCORSHeaders(w)

	api.model.DeleteCardWithId(id)

	w.WriteHeader(http.StatusNoContent)
}

func (api *Api) HandleGiven1Type2Request(w http.ResponseWriter,
	r *http.Request) {

	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	cardList := api.model.ListCards("ORDER BY last_answered_at LIMIT 1")
	if len(cardList.Cards) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
		return
	}
	topCard := cardList.Cards[0]

	bytes, err := json.Marshal(topCard)
	if err != nil {
		log.Fatalf("Error from json.Marshal: %s", err)
	}
	w.Write(bytes)
}

func (api *Api) HandleAnswerGiven1Type2Request(w http.ResponseWriter,
	r *http.Request) {

	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var answer Answer
	err = json.Unmarshal(body, &answer)
	if err != nil {
		panic(err)
	}

	api.model.TouchCardLastAnsweredAt(answer.CardId)

	bytes, err := json.Marshal(answer)
	if err != nil {
		log.Fatalf("Error from json.Marshal: %s", err)
	}
	w.Write(bytes)
}
