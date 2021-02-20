package human

import "aah/pkg/util"

const (
	minBabyHeight       = 30
	minAdultHeight      = 50
	maxAdultHeight      = 213
	minBabyWeight       = 1
	minAdultIdealWeight = 20
	maxAdultIdealWeight = 150
	minMaturityBaseRate = 1
	maxMaturityBaseRate = 5
)

type Height struct {
	Current int
	Max     int
}

type Weight struct {
	Current int
	Ideal   int
}

func generateHeight() *Height {
	maxHeight := util.Roll(minAdultHeight, maxAdultHeight+1)
	maxBabyHeight := maxHeight / 4
	if maxBabyHeight < minBabyHeight {
		maxBabyHeight = minBabyHeight
	}
	currentHeight := util.Roll(minBabyHeight, maxBabyHeight+1)
	return &Height{
		Current: currentHeight,
		Max:     maxHeight,
	}
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
