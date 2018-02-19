package server

import (
	"boston/learner"
	"fmt"
	"net/http"
)

func handleResetRequest(lmap *learner.Map) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("Reset Learner Started")
		payload, err := learner.ResetPayloadFromJSON(r.Body)
		if err != nil {
			apiError := NewAPIError("Reset Learner JSON Error", err.Error())
			respondAPIError(w, apiError)
			return
		}
		name := payload.Name
		action := learner.NewAction(name, learner.ResetSignal, payload)
		action.SuccessCallback = func(result interface{}) {
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
			"name":   name,
			"action": "reseted",
		}
		respondWithJSON(w, responseMap)
	}
}
