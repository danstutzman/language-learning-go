package api

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type SyncResponse struct {
	Cards      []db.Card      `json:"cards"`
	CardStates []db.CardState `json:"cardStates"`
	Notes      []db.Note      `json:"notes"`
}

type SyncRequest struct {
	Uploads []Upload
	Notes   []db.Note
}

type Upload struct {
	UploadId        int    `json:"uploadId"`
	Type            string `json:"type"`
	CardId          int    `json:"cardId"`
	CreatedAtMillis int64  `json:"createdAtMillis"`
	LogJson         string `json:"logJson"`
	StateJson       string `json:"stateJson"`
}

func (api *Api) HandleSyncRequest(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Error from ioutil.ReadAll: %s", err)
	}
	defer r.Body.Close()

	var syncRequest SyncRequest
	err = json.Unmarshal(body, &syncRequest)
	if err != nil {
		log.Fatalf("Error from json.Unmarshal: %s", err)
	}

	cardStates := []db.CardState{}
	for _, upload := range syncRequest.Uploads {
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

	for _, note := range syncRequest.Notes {
		db.UpsertNote(&note, api.db)
	}

	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	response := SyncResponse{
		Cards:      db.SelectAllFromCards(api.db),
		CardStates: db.SelectAllFromCardStates(api.db),
		Notes:      db.SelectAllFromNotes(api.db),
	}
	bytes, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error from json.Marshal: %s", err)
	}
	w.Write(bytes)
}
