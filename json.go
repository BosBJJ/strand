package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type jsonValid struct {
	Validity bool 	`json:"valid"`
}

type jsonInvalid struct {
	Error string 	`json:"error"`
}

type jsonResponse struct {
	Body string 	`json:"body"`
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	errorMsg := jsonInvalid{
		Error: msg,
	}
	respondWithJSON(w, code, errorMsg)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	resp, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Error Marshaling JSON: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}

