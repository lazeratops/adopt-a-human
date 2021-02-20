package events

import (
	"log"
	"math"
	"math/rand"
)

func Roll(min, max int) int {
	if max == -1 {
		max = 100
	}
	if max < min {
		log.Fatalf("max (%d) cannot be lower than min (%d)", max, min)
	}
	if max <= 0 {
		// if max <= 0, make max and min positive and then flip
		min = int(math.Abs(float64(max)))
		max = int(math.Abs(float64(min)))
	}
	return rand.Intn(max - min) + min
}
