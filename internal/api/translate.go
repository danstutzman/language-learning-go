package api

import (
	"cloud.google.com/go/translate"
	"context"
	"encoding/json"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
	"net/http"
)

func (api *Api) HandleTranslateRequest(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	l1s, ok := r.URL.Query()["l1"]
	if !ok || l1s[0] == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	l1 := l1s[0]

	client, err := translate.NewClient(
		context.TODO(),
		option.WithAPIKey(api.googleTranslateApiKey))
	if err != nil {
		panic(err)
	}

	translations, err := client.Translate(
		context.TODO(),
		[]string{l1},
		language.Spanish,
		&translate.Options{Source: language.English})
	if err != nil {
		panic(err)
	}
	l2 := translations[0].Text

	bytes, err := json.Marshal(l2)
	if err != nil {
		panic(err)
	}
	w.Write(bytes)
}
