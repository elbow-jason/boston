package learner

const (
	// CreateSignal is the signal to create a network
	CreateSignal = iota
	// DeleteSignal is the signal to delete a network
	DeleteSignal
	// ResetSignal is the signal to reset a network
	ResetSignal
	// TrainSignal is the signal to train a network
	TrainSignal
	// PredictSignal is the signal to use a network to predict inputs
	PredictSignal
)
