package learner

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

var (
	errorPredictNameRequired          = errors.New("Predict 'name' cannot be blank")
	errorPredictEmptyInputs           = errors.New("Predict 'inputs' cannot be empty")
	errorPredictPayloadInputsSizeZero = errors.New("Predict inputsSize cannot be 0")
)

func makeErrorPredictInputsSize(size int) error {
	message := fmt.Sprintf("Predict all 'inputs' size must be %d", size)
	return errors.New(message)
}

// PredictPayload .
type PredictPayload struct {
	Name       string      `json:"name"`
	Inputs     [][]float64 `json:"inputs"`
	InputsSize int
}

// ValidatePredictPayload .
func ValidatePredictPayload(payload PredictPayload) error {
	if payload.Name == "" {
		return errorResetNameRequired
	}
	size := payload.InputsSize
	if size == 0 {
		return errorPredictPayloadInputsSizeZero
	}
	for _, inputGroup := range payload.Inputs {
		if len(inputGroup) != size {
			return makeErrorPredictInputsSize(size)
		}
	}
	return nil
}

// PredictPayloadFromJSON .
func PredictPayloadFromJSON(r io.Reader) (PredictPayload, error) {
	payload := PredictPayload{}
	err := json.NewDecoder(r).Decode(&payload)
	if err != nil {
		return payload, err
	}
	return payload, nil

}
