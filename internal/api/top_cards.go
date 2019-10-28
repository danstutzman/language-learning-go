package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func (api *Api) HandleGetTopCardsRequest(w http.ResponseWriter,
	r *http.Request) {

	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	cardList := api.model.GetTopCards()

	bytes, err := json.Marshal(cardList)
	if err != nil {
		log.Fatalf("Error from json.Marshal: %s", err)
	}
	w.Write(bytes)
}
