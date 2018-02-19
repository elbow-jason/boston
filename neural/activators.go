package neural

import "math"

// Sigmoid implements the sigmoid function
// for use in activation functions.
func Sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

// SigmoidPrime implements the derivative
// of the sigmoid function for backpropagation.
func SigmoidPrime(x float64) float64 {
	return x * (1.0 - x)
}
