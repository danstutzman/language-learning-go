package api

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (api *Api) HandleListChallengesRequest(w http.ResponseWriter,
	r *http.Request) {

	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	answerList := api.model.ListChallenges()

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

	challenge := api.model.GetTopGiven1Type2Challenge()
	if challenge == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
		return
	}

	bytes, err := json.Marshal(challenge)
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

	var challenge model.Challenge
	err = json.Unmarshal(body, &challenge)
	if err != nil {
		panic(err)
	}

	api.model.ReplaceChallenge(challenge)

	bytes, err := json.Marshal(challenge)
	if err != nil {
		log.Fatalf("Error from json.Marshal: %s", err)
	}
	w.Write(bytes)
}
