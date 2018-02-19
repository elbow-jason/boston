package neural

// NetworkConfig defines our neural network
// architecture and learning parameters.
type NetworkConfig struct {
	InputNeurons  int
	HiddenNeurons int
	OutputNeurons int
	NumEpochs     int
	LearningRate  float64
}
