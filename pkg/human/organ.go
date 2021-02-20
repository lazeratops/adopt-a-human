package human

import (
	"aah/pkg/util"
	"encoding/json"
	"fmt"
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
	currentHealth      int
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
	o.currentHealth = util.Roll(5, maxBabyHealth + 1)
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
	// To start with, we generate a growth rate that will ensure even growth until Maturity
	bgr := util.WhatIsPercentOf(baseBodyMaturityRate, idealWeightG)
	if bgr == 0 {
		bgr = 1
	}
	o.baseGrowthRate = bgr
	// But when the human is at its youngest, it grows 50% faster than its base rate
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
	if o.currentHealth <= 0 {
		// This organ is already dead
		return
	}
	if o.currentHealth < o.maxHealth && o.body.maturity.Current < 100 {
		// Increase health by Current Maturity percentage of Max health
		o.currentHealth += util.WhatIsPercentOf(o.body.maturity.Current, o.maxHealth)
	}
	// but we still have a chance to get damaged
	roll := util.Roll(0, 100)
	if int(o.body.Immunity.Current) > roll {
		// Check how serious the damage will be
		damageSeriousness := util.Roll(0, 5)
		var damageRoll int
		// Very magical damage seriousness calculations
		switch damageSeriousness {
		case 0:
			damageRoll = util.Roll(0, 5)
		case 1:
			damageRoll = util.Roll(0, 10)
		case 2:
			damageRoll = util.Roll(0, 20)
		case 3:
			damageRoll = util.Roll(0, 25)
		case 4:
			damageRoll = util.Roll(0, 35)
		case 5:
			damageRoll = util.Roll(0, 50)
		}
		o.currentHealth -= damageRoll
	}
	fmt.Printf("\n%v health: %d", o.Name(), o.currentHealth)
}

func (o *Organ) currentGrowthRate() int {
	return o.baseGrowthRate + o.growthRateModifier
}

func (o *Organ) AddHealth(modifier int) {
	o.currentHealth += modifier
}

func (o *Organ) SubHealth(modifier int) {
	o.currentHealth -= modifier
}

func (o *Organ) Name() string {
	return string(o.kind)
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
		CurrentHealth:      o.currentHealth,
		MaxHealth:          o.maxHealth,
		BaseGrowthRate:     o.baseGrowthRate,
		GrowthRateModifier: o.growthRateModifier,
	})
}
