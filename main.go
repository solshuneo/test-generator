package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

func generateIntNumber(min, max int) int {
	return rand.IntN(max-min) + min
}

func generateFloatNumber(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

func FloatWithPrecision(a float64, precision int) float64 {
	shift := math.Pow10(precision)
	return math.Round(a*shift) / shift
}

func main() {
	intNum := generateIntNumber(0, 100)
	floatNum := FloatWithPrecision(generateFloatNumber(0.0, 100.0), 3)
	fmt.Println(intNum, floatNum)
}
