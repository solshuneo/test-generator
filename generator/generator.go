package generator

import (
	"math"
	"math/rand/v2"
)

func GenerateIntNumber(min, max int) int {
	return rand.IntN(max-min) + min
}

func GenerateFloatNumber(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

func FloatWithPrecision(a float64, precision int) float64 {
	shift := math.Pow10(precision)
	return math.Round(a*shift) / shift
}
