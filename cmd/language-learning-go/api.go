package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

type Api struct {
	db *sql.DB
}

func (api *Api) handleApiRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=\"utf-8\"")
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}
