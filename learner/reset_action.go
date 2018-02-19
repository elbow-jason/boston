package learner

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

var (
	errorResetNameRequired = errors.New("ResetAction 'name' cannot be blank")
)

// ResetAction .
type ResetAction struct {
	Name            string `json:"name"`
	SuccessCallback func()
	FailureCallback func(error)
}

// LearnerName .
func (a ResetAction) LearnerName() string {
	return a.Name
}

// Validate a DeleteAction
func (a ResetAction) Validate() error {
	if a.Name == "" {
		return errorResetNameRequired
	}
	return nil
}

// Signal .
func (a ResetAction) Signal() int {
	return ResetSignal
}

// Payload returns nil
func (a ResetAction) Payload() interface{} {
	return nil
}

// Success a callback upon success
func (a ResetAction) Success() {
	fmt.Println("ResetAction Success Called")
	if a.SuccessCallback != nil {
		a.SuccessCallback()
	}
}

// Failure a callback upon failure
func (a ResetAction) Failure(err error) {
	fmt.Println("ResetAction Failure Called - ", err.Error())
	if a.FailureCallback != nil {
		a.FailureCallback(err)
	}
}

// ResetActionFromJSON .
func ResetActionFromJSON(r io.Reader) (ResetAction, error) {
	action := ResetAction{}
	err := json.NewDecoder(r).Decode(&action)
	if err != nil {
		return action, err
	}
	if err = action.Validate(); err != nil {
		return action, err
	}
	return action, nil
}
