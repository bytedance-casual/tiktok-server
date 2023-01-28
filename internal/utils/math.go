package utils

import (
	"math"
	"math/rand"
)

func Sigmoid(x float64) float64 {
	return 1 / (1 + math.Pow(math.E, -x))
}

func RandomPick[T interface{}](arr []T) T {
	randIdx := rand.Intn(len(arr))
	return arr[randIdx]
}

func WeightedRandomPick[T interface{}](rates []float64, arr []T) T {
	idx := WeightedRandomIndex(rates)
	return arr[idx]
}

func WeightedRandomIndex(weights []float64) int {
	if len(weights) == 1 {
		return 0
	}
	var sum float64 = 0.0
	for _, w := range weights {
		sum += w
	}
	r := rand.Float64() * sum
	var t float64 = 0.0
	for i, w := range weights {
		t += w
		if t > r {
			return i
		}
	}
	return len(weights) - 1
}
