package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func (api *Api) HandleListSkillsRequest(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")

	skillList := api.model.ListSkills()

	bytes, err := json.Marshal(skillList)
	if err != nil {
		log.Fatalf("Error from json.Marshal: %s", err)
	}
	w.Write(bytes)
}
