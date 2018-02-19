package learner

import (
	"boston/neural"
	"encoding/json"
	"errors"
	"io"
)

var (
	errorCreateNameRequired    = errors.New("Create 'name' cannot be blank")
	errorInputNeuronsRequired  = errors.New("Create 'input_neurons' must be greater than 0")
	errorHiddenNeuronsRequired = errors.New("Create 'hidden_neurons' must be greater than 0")
	errorOutputNeuronsRequired = errors.New("Create 'output_neurons' must be greater than 0")
	errorNumEpochsRequired     = errors.New("Create 'num_epochs' must be greater than 0")
	errorLearningRateRange     = errors.New("Create 'learning_rate' must be between 0.0 and 1.0")
)

// CreatePayload .
type CreatePayload struct {
	Name          string  `json:"name"`
	InputNeurons  int     `json:"input_neurons"`
	HiddenNeurons int     `json:"hidden_neurons"`
	OutputNeurons int     `json:"output_neurons"`
	NumEpochs     int     `json:"num_epochs"`
	LearningRate  float64 `json:"learning_rate"`
}

// NewCreatePayload .
func NewCreatePayload(name string, inputNeurons int, hiddenNeurons int, outputNeurons int, numEpochs int, learningRate float64) (CreatePayload, error) {
	payload := CreatePayload{
		Name:          name,
		InputNeurons:  inputNeurons,
		HiddenNeurons: hiddenNeurons,
		OutputNeurons: outputNeurons,
		NumEpochs:     numEpochs,
		LearningRate:  learningRate,
	}
	return payload, ValidateCreatePayload(payload)
}

// CreatePayloadFromJSON .
func CreatePayloadFromJSON(r io.Reader) (CreatePayload, error) {
	payload := CreatePayload{}
	err := json.NewDecoder(r).Decode(&payload)
	if err != nil {
		return payload, err
	}
	return payload, ValidateCreatePayload(payload)
}

// ValidateCreatePayload .
func ValidateCreatePayload(payload CreatePayload) error {
	if payload.Name == "" {
		return errorCreateNameRequired
	}
	if payload.InputNeurons == 0 {
		return errorInputNeuronsRequired
	}
	if payload.HiddenNeurons == 0 {
		return errorHiddenNeuronsRequired
	}
	if payload.OutputNeurons == 0 {
		return errorOutputNeuronsRequired
	}
	if payload.NumEpochs == 0 {
		return errorNumEpochsRequired
	}
	if payload.LearningRate <= 0.0 || payload.LearningRate >= 1.0 {
		return errorLearningRateRange
	}
	return nil
}

// CreatePayloadToNeuralNetworkConfig .
func CreatePayloadToNeuralNetworkConfig(payload CreatePayload) neural.NetworkConfig {
	return neural.NetworkConfig{
		LearningRate:  payload.LearningRate,
		InputNeurons:  payload.InputNeurons,
		HiddenNeurons: payload.HiddenNeurons,
		OutputNeurons: payload.OutputNeurons,
		NumEpochs:     payload.NumEpochs,
	}
}
