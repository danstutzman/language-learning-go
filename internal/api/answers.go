package api

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (api *Api) HandleListAnswersRequest(w http.ResponseWriter,
	r *http.Request) {

	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	answerList := api.model.ListAnswers()

	bytes, err := json.Marshal(answerList)
	if err != nil {
		log.Fatalf("Error from json.Marshal: %s", err)
	}
	w.Write(bytes)
}

func (api *Api) HandleGiven1Type2Request(w http.ResponseWriter,
	r *http.Request) {

	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	cardId := api.model.GetTopGiven1Type2CardId()
	if cardId == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
		return
	}

	card := api.model.GetCard(cardId)

	bytes, err := json.Marshal(card)
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

	var answer model.Answer
	err = json.Unmarshal(body, &answer)
	if err != nil {
		panic(err)
	}

	api.model.ReplaceAnswer(answer)

	bytes, err := json.Marshal(answer)
	if err != nil {
		log.Fatalf("Error from json.Marshal: %s", err)
	}
	w.Write(bytes)
}
