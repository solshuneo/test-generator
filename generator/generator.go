package generator

import (
	"fmt"
	"math"
	"math/rand/v2"
)

func GenerateIntNumber(min, max int) (int, error) {
	if min > max {
		return 0, fmt.Errorf("min must be less than or equal to max")
	}
	return rand.IntN(max-min) + min, nil
}

func GenerateFloatNumber(min, max float64) (float64, error) {
	if min > max {
		return 0, fmt.Errorf("min must be less than or equal to max")
	}
	return rand.Float64()*(max-min) + min, nil
}

func FloatWithPrecision(a float64, precision int) (float64, error) {
	if precision < 0 {
		return 0, fmt.Errorf("precision must be non-negative")
	}
	shift := math.Pow10(precision)
	return math.Round(a*shift) / shift, nil
}
