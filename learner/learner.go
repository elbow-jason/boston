package learner

import (
	"boston/neural"
	"fmt"
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

func handlePredict(l *Learner, action Action) {
	fmt.Println("handlePredict", l.Name)
	inputs := action.Payload().([][]float64)
	outputs := [][]float64
}

func handleTrain(l *Learner, action Action) {
	fmt.Println("handleTrain before", l.Name, l.Network.IsTrained)
	entries := action.Payload().([]DataEntry)
	inputs, labels := DataEntriesToMatrices(entries)
	err := l.Network.Train(inputs, labels)
	if err != nil {
		action.Failure(err)
	}
	fmt.Println("handleTrain after", l.Name, l.Network.IsTrained)
}

func handleReset(l *Learner, action Action) {
	fmt.Println("handleReset", l.Name)
	l.Network.Reset()
	action.Success()
}

func handleCreate(l *Learner, action Action) {
	fmt.Println("handleCreate", l.Name)
	action.Success()
}

func handleDelete(l *Learner, action Action) {
	fmt.Println("handleDelete", l.Name)
	action.Success()
}
