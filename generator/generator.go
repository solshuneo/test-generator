package generator

import (
	"fmt"
	"math/rand/v2"
)

type Variable struct {
	Min int `yaml:"min"`
	Max int `yaml:"max"`
}

func GenerateIntNumber(min, max int) (int, error) {
	if min > max {
		return 0, fmt.Errorf("min must be less than or equal to max")
	}
	if min == max {
		return min, nil
	}
	return rand.IntN(max-min) + min, nil
}
