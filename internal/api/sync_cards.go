package api

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type SyncCardsResponse struct {
	Cards      []db.Card      `json:"cards"`
	CardStates []db.CardState `json:"cardStates"`
}

type UploadsRequest struct {
	Uploads []Upload
}

type Upload struct {
	UploadId        int    `json:"uploadId"`
	Type            string `json:"type"`
	CardId          int    `json:"cardId"`
	CreatedAtMillis int64  `json:"createdAtMillis"`
	LogJson         string `json:"logJson"`
	StateJson       string `json:"stateJson"`
}

func (api *Api) HandleSyncCardsRequest(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Error from ioutil.ReadAll: %s", err)
	}
	defer r.Body.Close()
	log.Printf("Request: %s", string(body))

	var uploadsRequest UploadsRequest
	err = json.Unmarshal(body, &uploadsRequest)
	if err != nil {
		log.Fatalf("Error from json.Unmarshal: %s", err)
	}

	cardStates := []db.CardState{}
	for _, upload := range uploadsRequest.Uploads {
		if upload.Type == "cardState" {
			cardStates = append(cardStates, db.CardState{
				CardId:    upload.CardId,
				StateJson: upload.StateJson,
			})
		} else if upload.Type == "log" {
			createdAt := time.Unix(
				upload.CreatedAtMillis/1000,
				upload.CreatedAtMillis%1000*1000000).Format(time.RFC3339Nano)
			log.Printf("Client log: %s %v", createdAt, upload.LogJson)
		}
	}
	if len(cardStates) > 0 {
		db.UpdateCardStates(cardStates, api.db)
		if err != nil {
			log.Fatal(err)
		}
	}

	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	response := SyncCardsResponse{
		Cards:      db.SelectAllFromCards(api.db),
		CardStates: db.SelectAllFromCardStates(api.db),
	}
	bytes, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error from json.Marshal: %s", err)
	}
	w.Write(bytes)
	log.Printf("Response: %s", string(bytes))
}
