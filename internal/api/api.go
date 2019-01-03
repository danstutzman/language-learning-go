package api

import (
	"bitbucket.org/danstutzman/language-learning-go/internal/db"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Api struct {
	db *sql.DB
}

type SyncResponse struct {
	Cards     []db.Card     `json:"cards"`
	Exposures []db.Exposure `json:"exposures"`
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
}

func NewApi(db *sql.DB) *Api {
	return &Api{db: db}
}

func setCORSHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "X-Client-Version")
}

func (api *Api) HandleApiRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		setCORSHeaders(w)
		return
	} else if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	log.Printf("X-Client-Version: %s", r.Header.Get("X-Client-Version"))

	decoder := json.NewDecoder(r.Body)
	var uploadsRequest UploadsRequest
	err := decoder.Decode(&uploadsRequest)
	if err != nil {
		log.Fatalf("Error from decoder.Decode: %s", err)
	}

	exposures := []db.Exposure{}
	for _, upload := range uploadsRequest.Uploads {
		if upload.Type == "exposure" {
			exposures = append(exposures, db.Exposure{
				CardId:          upload.CardId,
				CreatedAtMillis: upload.CreatedAtMillis,
			})
		} else if upload.Type == "log" {
			createdAt := time.Unix(
				upload.CreatedAtMillis/1000,
				upload.CreatedAtMillis%1000*1000000).Format(time.RFC3339Nano)
			log.Printf("Client log: %s %v", createdAt, upload.LogJson)
		}
	}
	if len(exposures) > 0 {
		db.InsertExposures(exposures, api.db)
		if err != nil {
			log.Fatal(err)
		}
	}

	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	if r.URL.Path == "/api/sync.json" {
		response := SyncResponse{
			Cards:     db.SelectAllFromCards(api.db),
			Exposures: db.SelectAllFromExposures(api.db),
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