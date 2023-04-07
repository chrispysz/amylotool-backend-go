// Go connection Sample Code:
package main

import (
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/microsoft/go-mssqldb"
)

type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func writeErrorResponse(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", contentTypeJSON)
	w.WriteHeader(code)
	response := ErrorResponse{
		Error: err.Error(),
		Code:  code,
	}
	responseBytes, _ := json.Marshal(response)
	w.Write(responseBytes)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	}

func main() {
	

	//initializeDatabase()

	http.HandleFunc("/predict", predictionHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
