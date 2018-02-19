package learner

import "fmt"

// Action holds all data necessary to
// make changes to or introspect a learner.
type Action struct {
	ActionSignal    int
	ActionName      string
	ActionPayload   interface{}
	SuccessCallback func(interface{})
	FailureCallback func(error)
}

// NewAction returns an Action without callbacks attached
func NewAction(name string, signal int, payload interface{}) Action {
	return Action{
		ActionName:    name,
		ActionSignal:  signal,
		ActionPayload: payload,
	}
}

// Name .
func (a Action) Name() string {
	return a.ActionName
}

// Validate a DeleteAction
func (a Action) Validate() error {
	if a.ActionName == "" {
		return errorDeleteNameRequired
	}
	return nil
}

// Signal .
func (a Action) Signal() int {
	return a.ActionSignal
}

// Payload returns nil
func (a Action) Payload() interface{} {
	return a.ActionPayload
}

// Success a callback upon success
func (a Action) Success(result interface{}) {
	fmt.Println("Action Success Called")
	if a.SuccessCallback != nil {
		a.SuccessCallback(result)
	}
}

// Failure a callback upon failure
func (a Action) Failure(err error) {
	fmt.Println("Action Failure Called - ", err.Error())
	if a.FailureCallback != nil {
		a.FailureCallback(err)
	}
}
