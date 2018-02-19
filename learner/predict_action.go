package learner

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"gonum.org/v1/gonum/mat"
)

var (
	errorPredictNameRequired = errors.New("PredictAction 'name' cannot be blank")
	errorPredictEmptyInputs  = errors.New("PredictAction 'inputs' cannot be empty")
)

func makeErrorPredictInputsSize(size int) error {
	message := fmt.Sprintf("PredictAction all 'inputs' size must be %d", size)
	return errors.New(message)
}

// PredictAction .
type PredictAction struct {
	Name            string      `json:"name"`
	Inputs          [][]float64 `json:"inputs"`
	SuccessCallback func()
	FailureCallback func(error)
}

func inputsToMatrix(inputs [][]float64, size int) (*mat.Dense, error) {

}

// LearnerName .
func (a PredictAction) LearnerName() string {
	return a.Name
}

// Payload returns []DataEntry
func (a PredictAction) Payload() interface{} {
	return a.Inputs
}

// Validate a DeleteAction
func (a PredictAction) Validate() error {
	if a.Name == "" {
		return errorResetNameRequired
	}
	return nil
}

// ValidateInputs a PredictAction's Entries Dimensions
func (a PredictAction) ValidateInputs(inputsSize int) error {
	for _, inputGroup := range a.Payload().([][]float64) {
		if len(inputGroup) != inputsSize {
			return makeErrorPredictInputsSize(inputsSize)
		}
	}
	return nil
}

// Signal .
func (a PredictAction) Signal() int {
	return PredictSignal
}

// Success a callback upon success
func (a PredictAction) Success() {
	fmt.Println("PredictAction Success Called")
	if a.SuccessCallback != nil {
		a.SuccessCallback()
	}
}

// Failure a callback upon failure
func (a PredictAction) Failure(err error) {
	fmt.Println("PredictAction Failure Called")
	if a.FailureCallback != nil {
		a.FailureCallback(err)
	}
}

// UnvalidatedPredictActionFromJSON .
func UnvalidatedPredictActionFromJSON(r io.Reader) (PredictAction, error) {
	action := PredictAction{}
	err := json.NewDecoder(r).Decode(&action)
	if err != nil {
		return action, err
	}
	return action, nil
}
