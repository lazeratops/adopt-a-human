package human

import (
	"aah/pkg/util"
	log "github.com/sirupsen/logrus"
)

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
	i.subImmunity(i.currentAttrition())
	if i.AttritionModifier != 0 {
		i.AttritionModifierDuration--
		if i.AttritionModifierDuration == 0 {
			i.AttritionModifier = 0
		}
	}
}

func (i *Immunity) currentAttrition() int {
	return i.BaseAttrition + i.AttritionModifier
}

func (i *Immunity) subImmunity(amount int) {
	log.WithField("amount", amount).Debug("subtracting immunity")
	new := i.Current - amount
	if new < 0 {
		new = 0
	}
	i.Current = new
}

func (i *Immunity) AddToAttritionModifier(mod, duration int) {
	i.AttritionModifier += mod
	i.AttritionModifierDuration = duration
}
