package human

import "aah/pkg/util"

const (
	minImmunityAttrition = 0
	maxImmunityAttrition = 50
	maxImmunity          = 250
)

type Immunity struct {
	Current                   int
	Max                       int
	AttritionModifier         int
	AttritionModifierDuration int
	BaseAttrition             int
}

func generateImmunity() *Immunity {
	max := util.Roll(11, maxImmunity)
	current := util.Roll(10, max)

	baseAttrition := util.Roll(minImmunityAttrition, maxImmunityAttrition)
	return &Immunity{
		Current:           current,
		Max:               max,
		AttritionModifier: 0,
		BaseAttrition:     baseAttrition,
	}
}

func (i *Immunity) tick(currentMaturity util.Percent) {
	// While the human is growing, the immunity is stabilizing.
	if currentMaturity < 100 {
		off := util.WhatIsPercentOf(currentMaturity, i.Max)
		i.Current = off
	}
	i.Current -= i.AttritionModifier
	if i.AttritionModifier != 0 {
		i.AttritionModifierDuration--
		if i.AttritionModifierDuration == 0 {
			i.AttritionModifier = 0
		}
	}
}
