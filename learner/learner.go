package learner

import (
	"boston/neural"
	"fmt"

	"gonum.org/v1/gonum/mat"
)

// Learner is the Actor for applying actions to
// a neural network
type Learner struct {
	Name    string
	Mailbox chan Action
	Network *neural.Network
}

// NewLearner .
func NewLearner(name string, config neural.NetworkConfig) *Learner {
	return &Learner{
		Name:    name,
		Mailbox: make(chan Action),
		Network: neural.NewNetwork(config),
	}
}

// WaitForSignal is the function to call for recieving
// Action Signals
func (l *Learner) WaitForSignal() {
	for {
		switch action := <-l.Mailbox; action.Signal() {
		case DeleteSignal:
			handleDelete(l, action)
			// exit the loop by returning.
			// should exit the loop.
			return
		case ResetSignal:
			handleReset(l, action)
		case CreateSignal:
			handleCreate(l, action)
		case TrainSignal:
			handleTrain(l, action)
		case PredictSignal:
			handlePredict(l, action)
		}
	}
}

func inputsToMatrix(inputs [][]float64, inputsSize int) *mat.Dense {
	rows := len(inputs)
	var grouped []float64
	for _, input := range inputs {
		grouped = append(grouped, input...)
	}
	return mat.NewDense(rows, inputsSize, grouped)
}

// MatrixToOutputs transforms a matrix into a
// slice of a slice of float64s
func MatrixToOutputs(matrix *mat.Dense) [][]float64 {
	rows, cols := matrix.Dims()
	outputs := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		rowSlice := make([]float64, cols)
		for j := 0; j < cols; j++ {
			rowSlice[j] = matrix.At(i, j)
		}
		outputs[i] = rowSlice
	}
	return outputs
}

func handlePredict(l *Learner, action Action) {
	fmt.Println("handlePredict", l.Name)
	payload := action.Payload().(PredictPayload)
	matrix := inputsToMatrix(payload.Inputs, payload.InputsSize)
	predictions, err := l.Network.Predict(matrix)
	if err != nil {
		action.Failure(err)
		return
	}
	action.Success(predictions)
}

func handleTrain(l *Learner, action Action) {
	fmt.Println("handleTrain before", l.Name, l.Network.IsTrained)
	payload := action.Payload().(TrainPayload)
	if payload.TestSplit > 0.0 {
		trains, tests := applyTrainTestSplit(payload)
		trainInputs, trainLabels := DataEntriesToMatrices(trains)
		err := l.Network.Train(trainInputs, trainLabels)
		if err != nil {
			action.Failure(err)
			return
		}
		testInputs, testLabels := DataEntriesToMatrices(tests)
		accuracy := l.Network.TestAccuracy(testInputs, testLabels)
		result := TrainResult{
			Accuracy: accuracy,
			Tested:   true,
		}
		action.SuccessCallback(result)
	} else {
		trainInputs, trainLabels := DataEntriesToMatrices(payload.DataEntries)
		err := l.Network.Train(trainInputs, trainLabels)
		if err != nil {
			action.Failure(err)
			return
		}
		result := TrainResult{
			Accuracy: 0.0,
			Tested:   false,
		}
		action.SuccessCallback(result)
	}
}

func handleReset(l *Learner, action Action) {
	fmt.Println("handleReset", l.Name)
	l.Network.Reset()
	action.Success(nil)
}

func handleCreate(l *Learner, action Action) {
	fmt.Println("handleCreate", l.Name)
	action.Success(nil)
}

func handleDelete(l *Learner, action Action) {
	fmt.Println("handleDelete", l.Name)
	action.Success(nil)
}
