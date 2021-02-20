package human

import "aah/pkg/util"

const (
	minMentalProperty = 10
	maxMentalProperty = 1000
)

type MentalProperty struct {
	Base    int
	Current int
}
type Mind struct {
	Resilience *MentalProperty
	Kindness   *MentalProperty
	Stress     *MentalProperty
	Maturity   *Maturity
}

func generateMind() *Mind {
	return &Mind{
		Resilience: generate(),
		Kindness:   generate(),
		Stress:     generate(),
		Maturity:   generateMaturity(),
	}
}

func generate() *MentalProperty {
	base := util.Roll(minMentalProperty, maxMentalProperty)

	maxBabyResilience := maxMentalProperty / 10
	if maxBabyResilience < minMentalProperty {
		maxBabyResilience = minMentalProperty
	}
	current := util.Roll(minMentalProperty, maxBabyResilience+1)
	return &MentalProperty{
		Base:    base,
		Current: current,
	}
}

func (m *Mind) tick() {
	if m.Maturity.Current < 100 {
		m.Resilience.Current += util.WhatIsPercentOf(m.Maturity.Current, m.Resilience.Base)
		m.Kindness.Current += util.WhatIsPercentOf(m.Maturity.Current, m.Kindness.Base)
		m.Stress.Current += util.WhatIsPercentOf(m.Maturity.Current, m.Stress.Base)
	}

	m.Maturity.tick()
}
