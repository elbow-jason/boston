package server

import (
	"boston/learner"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func handleTrainRequest(lmap *learner.Map) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("Train Learner Started")
		payload, err := learner.TrainPayloadFromJSON(r.Body)
		if err != nil {
			apiError := NewAPIError("Train Learner JSON Error", err.Error())
			respondAPIError(w, apiError)
			return
		}
		name := payload.Name
		config, err := lmap.GetConfig(name)
		if err != nil {
			apiError := NewAPIError("Train Learner Request Error", err.Error())
			respondAPIError(w, apiError)
			return
		}
		payload.InputsSize = config.InputNeurons
		payload.LabelsSize = config.OutputNeurons
		err = learner.ValidateTrainPayload(payload)
		if err != nil {
			apiError := NewAPIError("Train Learner Request Error", err.Error())
			respondAPIError(w, apiError)
			return
		}
		action := learner.NewAction(name, learner.TrainSignal, payload)
		successChan := make(chan learner.TrainResult)
		failureChan := make(chan error)
		action.SuccessCallback = func(result interface{}) {
			trainResult := result.(learner.TrainResult)
			fmt.Printf("Learner Train Success Callback Called with result %v\n", trainResult)
			successChan <- trainResult
		}
		action.FailureCallback = func(failure error) {
			fmt.Println("Learner Train Failure Callback Called - ", failure.Error())
			failureChan <- failure
		}
		err = lmap.SendAction(action)
		if err != nil {
			apiError := NewAPIError("Train Learner Action Error", err.Error())
			respondAPIError(w, apiError)
			return
		}
		select {
		case err = <-failureChan:
			apiError := NewAPIError("Learner Train Processing Error", err.Error())
			respondAPIError(w, apiError)
			return
		case trainResult := <-successChan:
			responseMap := map[string]interface{}{
				"name":     name,
				"action":   "train",
				"tested":   trainResult.Tested,
				"accuracy": trainResult.Accuracy,
			}
			respondWithJSON(w, responseMap)
			return
		case <-time.After(120 * time.Second):
			err = errors.New("Learner Train Timed Out")
			apiError := NewAPIError("Learner Train Processing Error", err.Error())
			respondAPIError(w, apiError)
			return
		}

	}
}
