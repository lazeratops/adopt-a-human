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

type heightCM struct {
	current int
	max     int
}

type weight struct {
	current int
	ideal   int
}

type maturity struct {
	current      util.Percent
	baseRate     util.Percent
	rateModifier util.Percent
}

func (m *maturity) currentRate() util.Percent {
	return m.baseRate + m.rateModifier
}

type mind struct {
}

func generateHeight() *heightCM {
	maxHeight := util.Roll(minAdultHeight, maxAdultHeight+1)
	maxBabyHeight := maxHeight / 4
	if maxBabyHeight < minBabyHeight {
		maxBabyHeight = minBabyHeight
	}
	currentHeight := util.Roll(minBabyHeight, maxBabyHeight+1)
	return &heightCM{
		current: currentHeight,
		max:     maxHeight,
	}
}

func generateWeight() *weight {
	idealWeight := util.Roll(minAdultIdealWeight, maxAdultIdealWeight+1)
	maxBabyWeight := idealWeight / 4
	if maxBabyWeight < minBabyWeight {
		maxBabyWeight = minBabyWeight + 1
	}
	currentWeight := util.Roll(minBabyWeight, maxBabyWeight+1)
	return &weight{
		current: currentWeight,
		ideal:   idealWeight,
	}
}

func generateMaturity() *maturity {
	br := util.Roll(minMaturityBaseRate, maxMaturityBaseRate+1)
	return &maturity{
		current:      0,
		baseRate:     util.Percent(br),
		rateModifier: 0,
	}
}
