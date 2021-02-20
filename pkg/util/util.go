package util

import (
	"fmt"
	"math"
)

type Percent int

//  if `total` == 100%, what percent is `part`?
func GetPercent(part int, total int) Percent {
	return Percent(math.Round(100 * float64(part) / float64(total)))
}

// What is p% of n?
func WhatIsPercentOf(p Percent, total int) int {
	return int(math.Round(float64(total) * float64(p) / 100))
}

func (p Percent) String() string {
	return fmt.Sprintf("%d%%", p)
}
