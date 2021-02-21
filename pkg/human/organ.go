package human

import (
	"aah/pkg/util"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"math"
)

type organKind string

const (
	OrganHeart  organKind = "heart"
	OrganLung   organKind = "lung"
	OrganKidney organKind = "kidney"
	OrganBrain  organKind = "brain"
)

type Organ struct {
	body               *Body
	kind               organKind
	weightG            *Weight
	CurrentHealth      int
	maxHealth          int
	baseGrowthRate     int
	growthRateModifier int
}

func generateOrgans(body *Body) []*Organ {
	heart := Organ{
		kind: OrganHeart,
		body: body,
	}
	heart.generateAndSetHealths()
	heart.generateAndSetWeight(100, 500)
	heart.generateAndSetGrowthRate(body.maturity.BaseRate, heart.weightG.Ideal)

	brain := Organ{
		kind: OrganBrain,
		body: body,
	}
	brain.generateAndSetHealths()
	brain.generateAndSetWeight(500, 3000)
	brain.generateAndSetGrowthRate(body.maturity.BaseRate, brain.weightG.Ideal)

	kidney1 := Organ{
		kind: OrganKidney,
		body: body,
	}
	kidney1.generateAndSetHealths()
	kidney1.generateAndSetWeight(25, 400)
	kidney1.generateAndSetGrowthRate(body.maturity.BaseRate, kidney1.weightG.Ideal)

	kidney2 := Organ{
		kind: OrganKidney,
		body: body,
	}
	kidney2.generateAndSetHealths()
	kidney2.generateAndSetWeight(25, 400)
	kidney2.generateAndSetGrowthRate(body.maturity.BaseRate, kidney2.weightG.Ideal)

	lung1 := Organ{
		kind: OrganLung,
		body: body,
	}
	lung1.generateAndSetHealths()
	lung1.generateAndSetWeight(50, 400)
	lung1.generateAndSetGrowthRate(body.maturity.BaseRate, lung1.weightG.Ideal)

	lung2 := Organ{
		kind: OrganLung,
		body: body,
	}
	lung2.generateAndSetHealths()
	lung2.generateAndSetWeight(50, 400)
	lung2.generateAndSetGrowthRate(body.maturity.BaseRate, lung2.weightG.Ideal)

	return []*Organ{&heart, &brain, &kidney1, &kidney2, &lung1, &lung2}
}

func (o *Organ) generateAndSetHealths() {
	o.maxHealth = util.Roll(50, 250)
	maxBabyHealth := util.WhatIsPercentOf(25, o.maxHealth)
	o.CurrentHealth = util.Roll(5, maxBabyHealth+1)
}

func (o *Organ) generateAndSetWeight(minIdeal int, maxIdeal int) {
	idealWeight := util.Roll(minIdeal, maxIdeal)
	maxBabyWeight := idealWeight / 4
	if maxBabyWeight < minIdeal {
		maxBabyWeight = minIdeal
	}
	currentWeight := util.Roll(1, maxBabyWeight+1)
	o.weightG = &Weight{
		Current: currentWeight,
		Ideal:   idealWeight,
	}
}

func (o *Organ) generateAndSetGrowthRate(baseBodyMaturityRate util.Percent, idealWeightG int) error {
	// To start with, we generateMentalProperty a growth rate that will ensure even growth until Maturity
	bgr := util.WhatIsPercentOf(baseBodyMaturityRate, idealWeightG)
	if bgr == 0 {
		bgr = 1
	}
	o.baseGrowthRate = bgr
	// But when the human is at its youngest, it grows 50% faster than its Base rate
	o.growthRateModifier = util.WhatIsPercentOf(util.Percent(50), o.baseGrowthRate)
	return nil
}

func (o *Organ) grow() {
	o.weightG.Current += o.currentGrowthRate()

	// Decrease growth rate modifier by whatever the Current Maturity percentage is
	p := util.WhatIsPercentOf(o.body.maturity.Current, o.growthRateModifier)
	o.growthRateModifier -= p
}

func (o *Organ) tickHealth() {
	if o.CurrentHealth <= 0 {
		// This organ is already dead
		return
	}
	// If the body is not yet mature, the organ is still growing
	if o.CurrentHealth < o.maxHealth && o.body.maturity.Current < 100 {
		// See what percentage points from ideal size we are off by
		p := 100 - (math.Abs(float64(util.GetPercent(o.weightG.Current, o.weightG.Ideal) - 100)))
		inc := util.WhatIsPercentOf(util.Percent(p), o.maxHealth)
		o.AddHealth(inc)
	}
	// give them a break so they don't die right away...
	if o.body.maturity.Current < 10 {
		return
	}
	// but we still have a chance to get damaged
	damageRoll := util.Roll(0, 100)
	immunityPerc := util.GetPercent(o.body.Immunity.Current, o.body.Immunity.Max)
	if int(immunityPerc) < damageRoll {
		damage := util.WhatIsPercentOf(100 - immunityPerc, damageRoll)
		o.SubHealth(damage)
	}

	// And we also still have a chance to recover
	recoveryRoll := util.Roll(0, 100)
	if int(immunityPerc) > recoveryRoll {
		// based on immunity percentage of total current immunity
		// roll := util.Roll(0, o.body.Immunity.Current)
		recovery := util.WhatIsPercentOf(immunityPerc, recoveryRoll)
		o.AddHealth(recovery)
	}
}

func (o *Organ) currentGrowthRate() int {
	return o.baseGrowthRate + o.growthRateModifier
}

func (o *Organ) AddHealth(modifier int) {
	log.WithField("modifier", modifier).Debug("adding organ health")
	new := o.CurrentHealth + modifier
	if new > o.maxHealth {
		new = o.maxHealth
	}
	o.CurrentHealth = new
}

func (o *Organ) SubHealth(modifier int) {
	log.WithField("modifier", modifier).Debug("subtracting organ health")
	o.CurrentHealth -= modifier
}

func (o *Organ) Name() string {
	return string(o.kind)
}

func (o *Organ) Descriptor() string {
	healthPerc := util.GetPercent(o.CurrentHealth, o.maxHealth)
	if healthPerc < 10 {
		return "on the verge of collapse"
	}
	if healthPerc < 25 {
		return "doing very NOT WELL"
	}
	if healthPerc < 50 {
		return "severely damaged"
	}
	if healthPerc < 75 {
		return "pretty good"
	}
	if healthPerc < 90 {
		return "in good health"
	}
	if healthPerc < 100 {
		return "in excellent health"
	}
	return "in perfect condition"
}

func (o *Organ) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind               organKind
		WeightG            *Weight
		CurrentHealth      int
		MaxHealth          int
		BaseGrowthRate     int
		GrowthRateModifier int
	}{
		Kind:               o.kind,
		WeightG:            o.weightG,
		CurrentHealth:      o.CurrentHealth,
		MaxHealth:          o.maxHealth,
		BaseGrowthRate:     o.baseGrowthRate,
		GrowthRateModifier: o.growthRateModifier,
	})
}
