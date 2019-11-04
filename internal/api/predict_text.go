package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type PredictTextRequestBody struct {
	WordSoFar string `json:"wordSoFar"`
}

func (api *Api) HandlePredictTextRequest(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var requestBody PredictTextRequestBody
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		panic(err)
	}

	predictions := api.model.PredictText(requestBody.WordSoFar)

	bytes, err := json.Marshal(predictions)
	if err != nil {
		log.Fatalf("Error from json.Marshal: %s", err)
	}
	w.Write(bytes)
}
