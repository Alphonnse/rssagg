package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) { 
	if code > 499 { // errors in the 400 range is the client side errors
		log.Println("Responding with 5XX error:", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errResponse{
		Error: msg,
	})
}


func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}){ // it takes as input the same http response writer that http handles in go use, a status code, and the JSON structure
	data, err := json.Marshal(payload) // Marshal and Unmarshal convert a string into JSON and vice versa
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-type", "application/json") // adding the header for the json response
	w.WriteHeader(statusCode) // just a status code
	w.Write(data)
}
