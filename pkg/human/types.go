package human

import "aah/pkg/events"

const (
	minBabyHeight = 30
	minAdultHeight = 50
	maxAdultHeight = 213
	minBabyWeight = 1
	minAdultIdealWeight = 20
	maxAdultIdealWeight = 150
	minMaturityBaseRate = 1
	maxMaturityBaseRate = 5
)


type heightCM struct {
	current int
	max int
}

type weight struct {
	current int
	ideal int
}

type maturity struct {
	currentPercent int
	baseRate int
	modifier int
}

func (m *maturity) currentRate() int {
	return m.baseRate + m.modifier
}

type mind struct {

}

func generateHeight() *heightCM {
	maxHeight := events.Roll(minAdultHeight, maxAdultHeight)
	maxBabyHeight := maxHeight / 4
	if maxBabyHeight > minBabyHeight {
		maxBabyHeight = minBabyHeight
	}
	currentHeight := events.Roll(30, maxBabyHeight)
	return &heightCM{
		current: currentHeight,
		max:     maxHeight,
	}
}


func generateWeight() *weight {
	idealWeight := events.Roll(minAdultIdealWeight, maxAdultIdealWeight)
	maxBabyWeight := idealWeight / 4
	if maxBabyWeight > minBabyWeight {
		maxBabyWeight = minBabyWeight
	}
	currentWeight := events.Roll(1, maxBabyWeight)
	return &weight{
		current: currentWeight,
		ideal:     idealWeight,
	}
}

func generateMaturity() *maturity {
	baseRate := events.Roll(minMaturityBaseRate, maxMaturityBaseRate)
	return &maturity{
		currentPercent: 0,
		baseRate:       baseRate,
		modifier:    0,
	}
}