package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PredictionRequest struct {
	Sequence string `json:"sequence"`
	ModelUrl string `json:"modelUrl"`
}

type Prediction struct {
	EndIndex   int    `json:"endIndex"`
	Prediction string `json:"prediction"`
	StartIndex int    `json:"startIndex"`
}

type PredictionResponse struct {
	Results []Prediction `json:"results"`
}

const (
	contentTypeJSON = "application/json"
)

func predictionHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method != http.MethodPost {
		writeErrorResponse(w, http.StatusBadRequest, fmt.Errorf("invalid HTTP method"))
		return
	}

	// Read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("error while reading raw UI request body"))
		return
	}

	// Unmarshal JSON body
	var req PredictionRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, fmt.Errorf("invalid JSON structure"))
		return
	}

	// Get the value of the "sequence" field
	sequence := req.Sequence
	modelServiceURL := req.ModelUrl

	payload := map[string]string{
		"sequence": sequence,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("error during model service request payload JSON marshalling"))
		return
	}

	resp, err := http.Post(modelServiceURL, contentTypeJSON, bytes.NewBuffer(payloadBytes))
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("error during POST request to model service"))
		return
	}
	defer resp.Body.Close()

	// Read response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("error while reading raw model service response body"))
		return
	}

	// Unmarshal JSON response body
	var res PredictionResponse
	err = json.Unmarshal(responseBody, &res)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("error during model service JSON response unmarshalling"))
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", contentTypeJSON)
	
	w.WriteHeader(http.StatusOK)

	// Marshal response object to JSON
	responseBytes, err := json.Marshal(res)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("error during response JSON marshalling"))
		return
	}

	// Write response body
	w.Write(responseBytes)
}
