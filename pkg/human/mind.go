package human

import (
	"aah/pkg/util"
)

const (
	minMentalProperty = 100
	maxMentalProperty = 10000
)

type MentalProperty struct {
	Base    int
	Current int
}
type Mind struct {
	Resilience *MentalProperty
	Kindness   *MentalProperty
	Stress     *MentalProperty
	Agreeableness *MentalProperty
	Stubbornness *MentalProperty
	Maturity   *Maturity
}

func generateMind() *Mind {
	return &Mind{
		Resilience:    generateMentalProperty(),
		Kindness:      generateMentalProperty(),
		Stress:        generateMentalProperty(),
		Agreeableness: generateMentalProperty(),
		Stubbornness:  generateMentalProperty(),
		Maturity:      generateMaturity(),
	}
}

func generateMentalProperty() *MentalProperty {
	base := util.Roll(minMentalProperty, maxMentalProperty)

	maxBabyMentalProperty := maxMentalProperty / 10
	if maxBabyMentalProperty < minMentalProperty {
		maxBabyMentalProperty = minMentalProperty
	}
	current := util.Roll(minMentalProperty, maxBabyMentalProperty+1)
	return &MentalProperty{
		Base:    base,
		Current: current,
	}
}

func (m *Mind) tick() {
	m.Resilience.tick(m.Maturity.Current, m.Maturity.currentRate())
	m.Kindness.tick(m.Maturity.Current, m.Maturity.currentRate())
	m.Stress.tick(m.Maturity.Current, m.Maturity.currentRate())
	m.Agreeableness.tick(m.Maturity.Current, m.Maturity.currentRate())
	m.Stubbornness.tick(m.Maturity.Current, m.Maturity.currentRate())
	m.Maturity.tick()
}

func (m *Mind) StateReport() string {
	var state string
	if m.Agreeableness.Current > m.Stubbornness.Current {
		state += "They are more agreeable than they are stubborn."
	} else if m.Agreeableness.Current < m.Stubbornness.Current {
		state += "They are more stubborn than they are agreeable."
	} else {
		state += "They are as agreeable as they are stubborn."
	}
	if m.Resilience.Current > m.Stress.Current {
		state += " They are more resilient than they are stressed."
	} else if m.Resilience.Current < m.Stress.Current {
		state += " They are more stressed than they are resilient."
	} else {
		state += " They are as resilient as they are stressed."
	}

	return state
}

func (m *MentalProperty) AddBase(modifier int) {
	m.Base += modifier
}

func (m *MentalProperty) AddCurrent(modifier int) {
	m.Current += modifier
}


func (m *MentalProperty) SubBase(modifier int) {
	new := m.Base - modifier
	if new < 0 {
		new = 0
	}
	m.Base = new
}

func (m *MentalProperty) SubCurrent(modifier int) {
	new := m.Current - modifier
	if new < 0 {
		new = 0
	}
	m.Current = new
}

func (m *MentalProperty) tick(maturityCurrent, rate util.Percent) {
	if maturityCurrent < 100 {
		m.Current += util.WhatIsPercentOf(rate, m.Base)
	}
	if m.Current > m.Base {
		m.Current--
	}
}