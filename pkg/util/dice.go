package util

import (
	"log"
	"math/rand"
)

func Roll(min, max int) int {
	if max == -1 {
		max = 100
	}

	if max < min {
		log.Fatalf("max (%d) must be higher than min (%d)", max, min)
	}

	if max-min == 0 {
		log.Fatalf("range must be in [min, max) format")
	}
	return rand.Intn(max-min) + min
}
