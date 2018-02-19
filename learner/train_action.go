package learner

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

var (
	errorTrainNameRequired = errors.New("TrainAction 'name' cannot be blank")
	errorTrainEmptyEntries = errors.New("TrainAction 'entries' cannot be empty")
)

// TrainAction .
type TrainAction struct {
	Name            string      `json:"name"`
	DataEntries     []DataEntry `json:"entries"`
	SuccessCallback func()
	FailureCallback func(error)
}

// Entries to train a network with.
func (a TrainAction) Entries() []DataEntry {
	return a.DataEntries
}

// LearnerName .
func (a TrainAction) LearnerName() string {
	return a.Name
}

// Payload returns []DataEntry
func (a TrainAction) Payload() interface{} {
	return a.DataEntries
}

// Validate a DeleteAction
func (a TrainAction) Validate() error {
	if a.Name == "" {
		return errorResetNameRequired
	}
	return nil
}

// ValidateDimensions a TrainAction's Entries Dimensions
func (a TrainAction) ValidateDimensions(inputsSize, labelsSize int) error {
	if len(a.Entries()) == 0 {
		return errorTrainEmptyEntries
	}
	var err error
	for _, entry := range a.Entries() {
		err = entry.Validate(inputsSize, labelsSize)
		if err != nil {
			return err
		}
	}
	return nil
}

// Signal .
func (a TrainAction) Signal() int {
	return TrainSignal
}

// Success a callback upon success
func (a TrainAction) Success() {
	fmt.Println("TrainAction Success Called")
	if a.SuccessCallback != nil {
		a.SuccessCallback()
	}
}

// Failure a callback upon failure
func (a TrainAction) Failure(err error) {
	fmt.Println("TrainAction Failure Called")
	if a.FailureCallback != nil {
		a.FailureCallback(err)
	}
}

// UnvalidatedTrainActionFromJSON .
func UnvalidatedTrainActionFromJSON(r io.Reader) (TrainAction, error) {
	action := TrainAction{}
	err := json.NewDecoder(r).Decode(&action)
	if err != nil {
		return action, err
	}
	return action, nil
}
