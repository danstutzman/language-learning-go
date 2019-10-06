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
	Cards     []db.Card     `json:"cards"`
	Morphemes []db.Morpheme `json:"morphemes"`
}

type SyncRequest struct {
	Uploads []Upload
	Cards   []db.Card
}

type Upload struct {
	UploadId        int    `json:"uploadId"`
	Type            string `json:"type"`
	CreatedAtMillis int64  `json:"createdAtMillis"`
	LogJson         string `json:"logJson"`
}

func (api *Api) HandleSyncRequest(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Error from ioutil.ReadAll: %s", err)
	}
	defer r.Body.Close()
	log.Printf("HandleSyncRequest body: %s", body)

	var syncRequest SyncRequest
	err = json.Unmarshal(body, &syncRequest)
	if err != nil {
		log.Fatalf("Error from json.Unmarshal: %s", err)
	}

	for _, upload := range syncRequest.Uploads {
		if upload.Type == "log" {
			createdAt := time.Unix(
				upload.CreatedAtMillis/1000,
				upload.CreatedAtMillis%1000*1000000).Format(time.RFC3339Nano)
			log.Printf("Client log: %s %v", createdAt, upload.LogJson)
		}
	}

	for _, card := range syncRequest.Cards {
		db.UpsertCard(&card, api.db)
	}

	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	response := SyncResponse{
		Cards:     db.SelectAllFromCards(api.db),
		Morphemes: db.SelectAllFromMorphemes(api.db),
	}
	bytes, err := json.Marshal(response)
	log.Printf("HandleSyncRequest response: %s", bytes)
	if err != nil {
		log.Fatalf("Error from json.Marshal: %s", err)
	}
	w.Write(bytes)
}
