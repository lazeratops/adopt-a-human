package human

import "aah/pkg/util"

const (
	minResilience = 10
	maxResilience = 1000
)

type Mind struct {
	BaseResilience    int
	CurrentResilience int
	Maturity          *Maturity
}

func generateMind() *Mind {
	baseResilience, currentResilience := generateResilience()
	return &Mind{
		BaseResilience:    baseResilience,
		CurrentResilience: currentResilience,
		Maturity:          generateMaturity(),
	}
}

func generateResilience() (int, int) {
	base := util.Roll(minResilience, maxResilience)

	maxBabyResilience := maxResilience / 10
	if maxBabyResilience < minResilience {
		maxBabyResilience = minResilience
	}
	current := util.Roll(minResilience, maxBabyResilience+1)
	return base, current
}

func (m *Mind) tick() {
	if m.Maturity.Current < 100 {
		// What is m.Maturity.Current percentage of BaselineResilience?
		inc := util.WhatIsPercentOf(m.Maturity.Current, m.BaseResilience)
		m.CurrentResilience += inc
	}

	m.Maturity.tick()
}
