package human

import "aah/pkg/util"

const (
	minMaturityBaseRate = 1
	maxMaturityBaseRate = 5
)

type Maturity struct {
	Current      util.Percent
	BaseRate     util.Percent
	RateModifier util.Percent
}

func generateMaturity() *Maturity {
	br := util.Roll(minMaturityBaseRate, maxMaturityBaseRate+1)
	return &Maturity{
		Current:      0,
		BaseRate:     util.Percent(br),
		RateModifier: 0,
	}
}

func (m *Maturity) currentRate() util.Percent {
	return m.BaseRate + m.RateModifier
}

func (m *Maturity) tick() {
	m.Current += m.currentRate()
	if m.Current >= 100 && m.RateModifier > 0 {
		m.RateModifier = 0
	}
}

func (m *Maturity) Descriptor() string {
	if m.Current < 10 {
		return "extremely immature"
	}
	if m.Current < 25 {
		return "very immature"
	}
	if m.Current < 50 {
		return "pretty immature"
	}
	if m.Current < 75 {
		return "progressing to Maturity"
	}
	if m.Current < 100 {
		return "nearly mature"
	}
	return "completely mature"
}
