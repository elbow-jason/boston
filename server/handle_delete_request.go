package server

import (
	"boston/learner"
	"fmt"
	"net/http"
)

func handleDeleteRequest(lmap *learner.Map) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("Delete Learner Started")
		payload, err := learner.DeletePayloadFromJSON(r.Body)
		if err != nil {
			apiError := NewAPIError("Delete Learner JSON Error", err.Error())
			respondAPIError(w, apiError)
			return
		}
		name := payload.Name
		action := learner.NewAction(name, learner.DeleteSignal, payload)
		action.SuccessCallback = func(interface{}) {
			fmt.Println("Learner Delete Success Callback Called")
		}
		action.FailureCallback = func(err error) {
			fmt.Println("Learner Delete Failure Callback Called - ", err.Error())
		}
		responseMap := map[string]interface{}{
			"name":   name,
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
