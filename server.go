package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func respondAPIError(w http.ResponseWriter, apiError ...APIError) {
	w.WriteHeader(http.StatusBadRequest)
	errors := newAPIErrors(apiError...)
	w.Write(errors.marshalJSON())
}

func handleCreateNetwork(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	job, err := learingJobFromJSON(r.Body)
	if err != nil {
		apiError := NewAPIError("Learning Job JSON error", err.Error())
		respondAPIError(w, apiError)
		return
	}
	fmt.Printf("Learning Job Params Accepted %d\n", len(job.Entries))
	w.Write([]byte("[\"ok\"]"))
}

func startHTTPServer() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/network/create", handleCreateNetwork)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
