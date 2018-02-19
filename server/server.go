package server

import (
	"boston/learner"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func respondAPIError(w http.ResponseWriter, apiError ...APIError) {
	w.WriteHeader(http.StatusBadRequest)
	errors := newAPIErrors(apiError...)
	w.Write(errors.marshalJSON())
}

func respondWithJSON(w http.ResponseWriter, jsonMap map[string]interface{}) {
	resp, err := json.Marshal(jsonMap)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Write(resp)
}

// StartLearnerHTTPServer starts the router associated with manipulating
// a learner.Map and it's member learners
func StartLearnerHTTPServer(lmap *learner.Map) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/learners/list", handleListRequest(lmap))
	router.HandleFunc("/learners/create", handleCreateRequest(lmap))
	router.HandleFunc("/learners/delete", handleDeleteRequest(lmap))
	router.HandleFunc("/learners/reset", handleResetRequest(lmap))
	router.HandleFunc("/learners/train", handleTrainRequest(lmap))
	router.HandleFunc("/learners/predict", handlePredictRequest(lmap))
	log.Fatal(http.ListenAndServe(":4351", router))
}
