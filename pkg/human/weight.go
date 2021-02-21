package human

import "aah/pkg/util"

const (
	minBabyWeight       = 1
	minAdultIdealWeight = 20
	maxAdultIdealWeight = 150
)

type Weight struct {
	Current int
	Ideal   int
}

func generateWeight() *Weight {
	idealWeight := util.Roll(minAdultIdealWeight, maxAdultIdealWeight+1)
	maxBabyWeight := idealWeight / 4
	if maxBabyWeight < minBabyWeight {
		maxBabyWeight = minBabyWeight + 1
	}
	currentWeight := util.Roll(minBabyWeight, maxBabyWeight+1)
	return &Weight{
		Current: currentWeight,
		Ideal:   idealWeight,
	}
}

func (w *Weight) tick(currentMaturity, rate util.Percent) {
	if currentMaturity < 100 {
		w.Current += util.WhatIsPercentOf(rate, w.Ideal)
	}
}
