package server

import (
	"boston/learner"
	"net/http"
)

func handleCreateRequest(lmap *learner.Map) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		payload, err := learner.CreatePayloadFromJSON(r.Body)
		if err != nil {
			apiError := NewAPIError("Create Learner JSON Error", err.Error())
			respondAPIError(w, apiError)
			return
		}
		name := payload.Name
		networkConfig := learner.CreatePayloadToNeuralNetworkConfig(payload)
		err = lmap.StartLearner(name, networkConfig)
		if err != nil {
			apiError := NewAPIError("Create Learner Request Error", err.Error())
			respondAPIError(w, apiError)
			return
		}
		action := learner.NewAction(name, learner.CreateSignal, payload)
		err = lmap.SendAction(action)
		if err != nil {
			apiError := NewAPIError("Create Learner Request Error", err.Error())
			respondAPIError(w, apiError)
			return
		}
		responseMap := map[string]interface{}{
			"name":   name,
			"action": "created",
		}
		respondWithJSON(w, responseMap)
	}
}
