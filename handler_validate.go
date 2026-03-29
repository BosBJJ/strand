package main

import (
	"encoding/json"
	"net/http"
)

func (cfg *apiConfig) handlerValidate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	resp := jsonResponse{}
	err := decoder.Decode(&resp)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}
	if len(resp.Body) > 140 {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return
	}
	respondWithJSON(w, 200, jsonValid{Validity: true})
}