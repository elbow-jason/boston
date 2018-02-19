package server

import (
	"boston/learner"
	"errors"
	"fmt"
	"net/http"
	"time"

	"gonum.org/v1/gonum/mat"
)

func handlePredictRequest(lmap *learner.Map) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("Predict Learner Started")
		payload, err := learner.PredictPayloadFromJSON(r.Body)
		if err != nil {
			apiError := NewAPIError("Predict Learner JSON Error", err.Error())
			respondAPIError(w, apiError)
			return
		}
		name := payload.Name
		config, err := lmap.GetConfig(name)
		if err != nil {
			apiError := NewAPIError("Predict Learner Request Error", err.Error())
			respondAPIError(w, apiError)
			return
		}
		payload.InputsSize = config.InputNeurons
		err = learner.ValidatePredictPayload(payload)
		if err != nil {
			apiError := NewAPIError("Predict Learner Request Error", err.Error())
			respondAPIError(w, apiError)
			return
		}
		action := learner.NewAction(name, learner.PredictSignal, payload)
		successChan := make(chan [][]float64)
		failureChan := make(chan error)
		action.SuccessCallback = func(result interface{}) {
			matrix := result.(*mat.Dense)
			outputs := learner.MatrixToOutputs(matrix)
			fmt.Println("Learner Predict Success Callback Called")
			fmt.Printf("Learner Predict Success Result %v\n", outputs)
			successChan <- outputs
		}
		action.FailureCallback = func(err error) {
			fmt.Println("Learner Predict Failure Callback Called - ", err.Error())
			failureChan <- err
		}
		err = lmap.SendAction(action)
		if err != nil {
			apiError := NewAPIError("Learner Predict Action Error", err.Error())
			respondAPIError(w, apiError)
			return
		}
		select {
		case err = <-failureChan:
			apiError := NewAPIError("Learner Predict Processing Error", err.Error())
			respondAPIError(w, apiError)
			return
		case predictions := <-successChan:
			responseMap := map[string]interface{}{
				"name":        name,
				"action":      "predict",
				"inputs":      payload.Inputs,
				"predictions": predictions,
			}
			respondWithJSON(w, responseMap)
			return
		case <-time.After(3 * time.Second):
			err = errors.New("Learner Prediction Timed Out")
			apiError := NewAPIError("Learner Predict Processing Error", err.Error())
			respondAPIError(w, apiError)
			return
		}
	}
}
