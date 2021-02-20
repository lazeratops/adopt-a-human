package human

import (
	"aah/pkg/util"
)

type organKind int

const (
	OrganHeart organKind = iota
	OrganLung
	OrganKidney
	OrganBrain
)

type organ struct {
	body               *body
	kind               organKind
	weightG            *weight
	currentHealth      int
	maxHealth          int
	baseGrowthRate     int
	growthRateModifier int
}

func generateOrgans(body *body) []*organ {
	heart := organ{
		kind: OrganHeart,
		body: body,
	}
	heart.generateAndSetHealths()
	heart.generateAndSetWeight(100, 500)
	heart.generateAndSetGrowthRate(body.maturity.baseRate, heart.weightG.ideal)

	brain := organ{
		kind: OrganBrain,
		body: body,
	}
	brain.generateAndSetHealths()
	brain.generateAndSetWeight(500, 3000)
	brain.generateAndSetGrowthRate(body.maturity.baseRate, brain.weightG.ideal)

	kidney1 := organ{
		kind: OrganKidney,
		body: body,
	}
	kidney1.generateAndSetHealths()
	kidney1.generateAndSetWeight(25, 400)
	kidney1.generateAndSetGrowthRate(body.maturity.baseRate, kidney1.weightG.ideal)

	kidney2 := organ{
		kind: OrganKidney,
		body: body,
	}
	kidney2.generateAndSetHealths()
	kidney2.generateAndSetWeight(25, 400)

	lung1 := organ{
		kind: OrganLung,
		body: body,
	}
	lung1.generateAndSetHealths()
	lung1.generateAndSetWeight(50, 400)

	lung2 := organ{
		kind: OrganLung,
		body: body,
	}
	lung2.generateAndSetHealths()
	lung2.generateAndSetWeight(50, 400)

	return []*organ{&heart, &brain, &kidney1, &kidney2, &lung1, &lung2}
}

func (o *organ) generateAndSetHealths() {
	o.maxHealth = o.getRandomHealth(-1)
	o.currentHealth = o.getRandomHealth(o.maxHealth)
}

func (o *organ) getRandomHealth(max int) int {
	if max == -1 {
		max = 100
	}
	return util.Roll(0, max+1)
}

func (o *organ) generateAndSetWeight(minIdeal int, maxIdeal int) {
	idealWeight := util.Roll(minIdeal, maxIdeal)
	maxBabyWeight := idealWeight / 4
	if maxBabyWeight < minIdeal {
		maxBabyWeight = minIdeal
	}
	currentWeight := util.Roll(1, maxBabyWeight+1)
	o.weightG = &weight{
		current: currentWeight,
		ideal:   idealWeight,
	}
}

func (o *organ) generateAndSetGrowthRate(baseBodyMaturityRate util.Percent, idealWeightG int) error {
	// To start with, we generate a growth rate that will ensure even growth until maturity
	o.baseGrowthRate = util.WhatIsPercentOf(baseBodyMaturityRate, idealWeightG)

	// But when the human is at its youngest, it grows 50% faster than its base rate
	o.growthRateModifier = util.WhatIsPercentOf(util.Percent(50), o.baseGrowthRate)
	return nil
}

func (o *organ) grow() {
	o.weightG.current += o.currentGrowthRate()

	// Decrease growth rate modifier by whatever the current maturity percentage is
	// What is o.body.maturity.current percent of o.growthRateModifier?
	p := util.WhatIsPercentOf(o.body.maturity.current, o.growthRateModifier)
	o.growthRateModifier -= p
}

func (o *organ) currentGrowthRate() int {
	return o.baseGrowthRate + o.growthRateModifier
}
