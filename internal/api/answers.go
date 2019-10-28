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

func (api *Api) HandleGetTopChallengesRequest(w http.ResponseWriter,
	r *http.Request) {

	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	type_ := r.URL.Query()["type"]
	if len(type_) != 1 || type_[0] == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	challengeList := api.model.GetTopChallenges(type_[0])

	bytes, err := json.Marshal(challengeList)
	if err != nil {
		log.Fatalf("Error from json.Marshal: %s", err)
	}
	w.Write(bytes)
}

func (api *Api) HandlePostAnswerRequest(w http.ResponseWriter,
	r *http.Request) {

	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var unsavedAnswer model.Answer
	err = json.Unmarshal(body, &unsavedAnswer)
	if err != nil {
		panic(err)
	}

	answer := api.model.InsertAnswer(unsavedAnswer)

	bytes, err := json.Marshal(answer)
	if err != nil {
		log.Fatalf("Error from json.Marshal: %s", err)
	}
	w.Write(bytes)
}
