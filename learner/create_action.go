package learner

import (
	"boston/neural"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

var (
	errorCreateNameRequired    = errors.New("CreateAction 'name' cannot be blank")
	errorInputNeuronsRequired  = errors.New("CreateAction 'input_neurons' must be greater than 0")
	errorHiddenNeuronsRequired = errors.New("CreateAction 'hidden_neurons' must be greater than 0")
	errorOutputNeuronsRequired = errors.New("CreateAction 'output_neurons' must be greater than 0")
	errorNumEpochsRequired     = errors.New("CreateAction 'num_epochs' must be greater than 0")
	errorLearningRateRange     = errors.New("CreateAction 'learning_rate' must be between 0.0 and 1.0")
)

// CreateAction .
type CreateAction struct {
	neural.NetworkConfig
	Name string `json:"name"`
}

// NewCreateAction .
func NewCreateAction(name string, inputNeurons int, hiddenNeurons int, outputNeurons int, numEpochs int, learningRate float64) (CreateAction, error) {
	config := neural.NetworkConfig{
		InputNeurons:  inputNeurons,
		HiddenNeurons: hiddenNeurons,
		OutputNeurons: outputNeurons,
		NumEpochs:     numEpochs,
		LearningRate:  learningRate,
	}
	newAction := CreateAction{
		Name:          name,
		NetworkConfig: config,
	}
	return newAction, newAction.Validate()
}

// CreateActionFromJSON .
func CreateActionFromJSON(r io.Reader) (CreateAction, error) {
	action := CreateAction{}
	err := json.NewDecoder(r).Decode(&action)
	if err != nil {
		return action, err
	}
	if err = action.Validate(); err != nil {
		return action, err
	}
	return action, nil
}

// Validate .
func (c CreateAction) Validate() error {
	if c.Name == "" {
		return errorCreateNameRequired
	}
	if c.InputNeurons == 0 {
		return errorInputNeuronsRequired
	}
	if c.HiddenNeurons == 0 {
		return errorHiddenNeuronsRequired
	}
	if c.OutputNeurons == 0 {
		return errorOutputNeuronsRequired
	}
	if c.NumEpochs == 0 {
		return errorNumEpochsRequired
	}
	if c.LearningRate <= 0.0 || c.LearningRate >= 1.0 {
		return errorLearningRateRange
	}
	return nil
}

// LearnerName .
func (c CreateAction) LearnerName() string {
	return c.Name
}

// Signal .
func (c CreateAction) Signal() int {
	return CreateSignal
}

// Success a callback upon success
func (c CreateAction) Success() {
	fmt.Println("CreateAction Success Called")
}

// Failure a callback upon failure
func (c CreateAction) Failure(err error) {
	fmt.Println("CreateAction Failure Called -", err.Error())
}

// Payload returns a network Config
func (c CreateAction) Payload() interface{} {
	return c.NetworkConfig
}

// ToNeuralNetworkConfig .
func (c CreateAction) ToNeuralNetworkConfig() neural.NetworkConfig {
	return neural.NetworkConfig{
		LearningRate:  c.LearningRate,
		InputNeurons:  c.InputNeurons,
		HiddenNeurons: c.HiddenNeurons,
		OutputNeurons: c.OutputNeurons,
		NumEpochs:     c.NumEpochs,
	}
}
