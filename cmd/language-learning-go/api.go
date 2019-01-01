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

type Card struct {
	CardId int    `json:"cardId"`
	Es     string `json:"es"`
}

type Exposure struct {
	CardId    int     `json:"cardId"`
	CreatedAt float64 `json:"createdAt"`
}

type SyncResponse struct {
	Cards     []Card     `json:"cards"`
	Exposures []Exposure `json:"exposures"`
}

func selectAllFromCards(db *sql.DB) []Card {
	cards := []Card{}

	rows, err := db.Query("select cardId, es from cards")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var card Card
		err = rows.Scan(&card.CardId, &card.Es)
		if err != nil {
			log.Fatal(err)
		}
		cards = append(cards, card)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return cards
}

func selectAllFromExposures(db *sql.DB) []Exposure {
	exposures := []Exposure{}

	rows, err := db.Query("select cardId, createdAt from exposures")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var exposure Exposure
		err = rows.Scan(&exposure.CardId, &exposure.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		exposures = append(exposures, exposure)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return exposures
}

type UploadsRequest struct {
	Uploads []Upload
}

type Upload struct {
	UploadId  int     `json:"uploadId"`
	Type      string  `json:"type"`
	CardId    string  `json:"cardId"`
	CreatedAt float64 `json:"createdAt"`
}

func insertExposures(uploadsRequest UploadsRequest, db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Error from db.Begin: %s", err)
	}

	stmt, err := tx.Prepare(
		"insert into exposures(cardId, createdAt) values(?,?)")
	if err != nil {
		log.Fatalf("Error from tx.Prepare: %s", err)
	}
	defer stmt.Close()

	for _, upload := range uploadsRequest.Uploads {
		if upload.Type == "exposure" {
			_, err = stmt.Exec(upload.CardId, upload.CreatedAt)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	tx.Commit()
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
