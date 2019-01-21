package api

import (
	"net/http"
)

func (api *Api) HandleDownloadDictionaryRequest(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)
	w.Header().Set("Content-Type", "text/plain; charset=\"utf-8\"")

	http.ServeFile(w, r, api.dictionaryPath)
}
