package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Api struct {
	db *sql.DB
}

type SyncResponse struct {
	Cards     []Card     `json:"cards"`
	Exposures []Exposure `json:"exposures"`
}

type UploadsRequest struct {
	Uploads []Upload
}

type Upload struct {
	UploadId  int     `json:"uploadId"`
	Type      string  `json:"type"`
	CardId    int     `json:"cardId"`
	CreatedAt float64 `json:"createdAt"`
}

func (api *Api) handleApiRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var uploadsRequest UploadsRequest
	err := decoder.Decode(&uploadsRequest)
	if err != nil {
		log.Fatalf("Error from decoder.Decode: %s", err)
	}
	insertExposures(uploadsRequest, api.db)

	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.URL.Path == "/api/sync.json" {
		response := SyncResponse{
			Cards:     selectAllFromCards(api.db),
			Exposures: selectAllFromExposures(api.db),
		}
		bytes, err := json.Marshal(response)
		if err != nil {
			log.Fatalf("Error from json.Marshal: %s", err)
		}
		w.Write(bytes)

	} else {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	}
}
