// Go connection Sample Code:
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"fmt"

	_ "github.com/microsoft/go-mssqldb"
	"github.com/rs/cors"
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


func main() {
	mux:= http.NewServeMux()


	//initializeDatabase()

	mux.HandleFunc("/predict", predictionHandler)
	handler := cors.Default().Handler(mux)

	log.Fatal(http.ListenAndServe(GetPort(), handler))

}

func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "8080"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
