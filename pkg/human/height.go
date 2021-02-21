package human

import "aah/pkg/util"

const (
	minBabyHeight  = 30
	minAdultHeight = 50
	maxAdultHeight = 213
)

type Height struct {
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
		Ideal:   maxHeight,
	}
}

func (h *Height) tick(currentMaturity, rate util.Percent) {
	if currentMaturity < 100 {
		h.Current += util.WhatIsPercentOf(rate, h.Ideal)
	}
}
