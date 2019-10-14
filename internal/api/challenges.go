package api

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
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

func (api *Api) HandleGetTopChallengeRequest(w http.ResponseWriter,
	r *http.Request) {

	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	type_ := r.URL.Query()["type"]
	if len(type_) != 1 || type_[0] == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	challenge := api.model.GetTopChallenge(type_[0])
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

func (api *Api) HandleAnswerChallengeRequest(w http.ResponseWriter,
	r *http.Request, challengeId int) {

	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var update db.ChallengeUpdate
	err = json.Unmarshal(body, &update)
	if err != nil {
		panic(err)
	}

	challenge := api.model.UpdateChallenge(update)

	bytes, err := json.Marshal(challenge)
	if err != nil {
		log.Fatalf("Error from json.Marshal: %s", err)
	}
	w.Write(bytes)
}
