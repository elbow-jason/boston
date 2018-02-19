package main

import (
	"boston/learner"
	"encoding/json"
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

func respondWithJSON(w http.ResponseWriter, jsonMap map[string]interface{}) {
	resp, err := json.Marshal(jsonMap)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Write(resp)
}

func handleCreateLearner(lmap *LearnerMap) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		action, err := learner.CreateActionFromJSON(r.Body)
		if err != nil {
			apiError := NewAPIError("Create Learner JSON Error", err.Error())
			respondAPIError(w, apiError)
			return
		}
		err = lmap.StartLearner(action.Name, action.ToNeuralNetworkConfig())
		if err != nil {
			apiError := NewAPIError("Create Learner Request Error", err.Error())
			respondAPIError(w, apiError)
			return
		}
		err = lmap.SendAction(action)
		if err != nil {
			apiError := NewAPIError("Create Learner Request Error", err.Error())
			respondAPIError(w, apiError)
			return
		}
		responseMap := map[string]interface{}{
			"name":   action.Name,
			"action": "created",
		}
		respondWithJSON(w, responseMap)
	}
}

func handleListLearners(lmap *LearnerMap) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("List Learner Keys")
		keys := lmap.Keys()
		keysBytes, err := json.Marshal(keys)
		if err != nil {
			apiError := NewAPIError("List Learner Keys API Error", err.Error())
			respondAPIError(w, apiError)
			return
		}
		w.Write(keysBytes)
	}
}

func handleDeleteLearner(lmap *LearnerMap) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("Delete Learner Started")
		action, err := learner.DeleteActionFromJSON(r.Body)
		if err != nil {
			apiError := NewAPIError("Delete Learner JSON Error", err.Error())
			respondAPIError(w, apiError)
			return
		}
		action.SuccessCallback = func() {
			fmt.Println("Learner Delete Success Callback Called")
		}
		action.FailureCallback = func(err error) {
			fmt.Println("Learner Delete Failure Callback Called - ", err.Error())
		}
		responseMap := map[string]interface{}{
			"name":   action.Name,
			"action": "deleted",
		}
		err = lmap.SendAction(action)
		if err != nil {
			apiError := NewAPIError("Delete Learner Request Error", err.Error())
			respondAPIError(w, apiError)
			return
		}
		respondWithJSON(w, responseMap)
	}
}

func handleTrainLearner(lmap *LearnerMap) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("Train Learner Started")
		action, err := learner.UnvalidatedTrainActionFromJSON(r.Body)
		if err != nil {
			apiError := NewAPIError("Train Learner JSON Error", err.Error())
			respondAPIError(w, apiError)
			return
		}

		if err != nil {
			apiError := NewAPIError("Train Learner JSON Error", err.Error())
			respondAPIError(w, apiError)
			return
		}
		action.SuccessCallback = func() {
			fmt.Println("Learner Train Success Callback Called")
		}
		action.FailureCallback = func(err error) {
			fmt.Println("Learner Train Failure Callback Called - ", err.Error())
		}
		config, err := lmap.GetConfig(action.Name)
		if err != nil {
			apiError := NewAPIError("Train Learner Request Error", err.Error())
			respondAPIError(w, apiError)
			return
		}
		err = action.ValidateDimensions(config.InputNeurons, config.OutputNeurons)
		if err != nil {
			apiError := NewAPIError("Train Learner Request Error", err.Error())
			respondAPIError(w, apiError)
			return
		}
		err = lmap.SendAction(action)
		if err != nil {
			apiError := NewAPIError("Train Learner Request Error", err.Error())
			respondAPIError(w, apiError)
			return
		}
		responseMap := map[string]interface{}{
			"name":   action.Name,
			"action": "train",
		}
		respondWithJSON(w, responseMap)
	}
}

func handleResetLearner(lmap *LearnerMap) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("Reset Learner Started")
		action, err := learner.ResetActionFromJSON(r.Body)
		if err != nil {
			apiError := NewAPIError("Reset Learner JSON Error", err.Error())
			respondAPIError(w, apiError)
			return
		}
		action.SuccessCallback = func() {
			fmt.Println("Learner Reset Success Callback Called")
		}
		action.FailureCallback = func(err error) {
			fmt.Println("Learner Reset Failure Callback Called - ", err.Error())
		}
		err = lmap.SendAction(action)
		if err != nil {
			apiError := NewAPIError("Reset Learner Request Error", err.Error())
			respondAPIError(w, apiError)
			return
		}
		responseMap := map[string]interface{}{
			"name":   action.Name,
			"action": "reseted",
		}
		respondWithJSON(w, responseMap)
	}
}

func startHTTPServer(lmap *LearnerMap) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/learners/list", handleListLearners(lmap))
	router.HandleFunc("/learners/create", handleCreateLearner(lmap))
	router.HandleFunc("/learners/delete", handleDeleteLearner(lmap))
	router.HandleFunc("/learners/reset", handleResetLearner(lmap))
	router.HandleFunc("/learners/train", handleTrainLearner(lmap))
	log.Fatal(http.ListenAndServe(":10000", router))
}
