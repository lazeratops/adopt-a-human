package human

import (
	"aah/pkg/events"
)

type organKind int

const (
	OrganHeart organKind = iota
	OrganLungs
	OrganKidneys
	OrganBrain
)

type organTemplate struct {
	name     string
	quantity int
}

type organ struct {
	kind          organKind
	quantity      int
	weightG       *weight
	currentHealth int
	maxHealth     int
}

func generateOrgans() []*organ {
	heart := organ{
		kind:     OrganHeart,
		quantity: 1,
	}
	heart.generateAndSetHealths()

	brain := organ{
		kind:     OrganBrain,
		quantity: 1,
	}
	brain.generateAndSetHealths()

	kidneys := organ{
		kind:     OrganKidneys,
		quantity: 2,
	}
	kidneys.generateAndSetHealths()

	lungs := organ{
		kind:     OrganLungs,
		quantity: 2,
	}
	lungs.generateAndSetHealths()
	return []*organ{&heart, &brain, &kidneys, &lungs}
}

func (o *organ) generateAndSetHealths() {
	o.maxHealth = o.getRandomHealth(-1)
	o.currentHealth = o.getRandomHealth(o.maxHealth)
}

func (o *organ) generateAndSetWeight(minIdeal int, maxIdeal int) {
	idealWeight := events.Roll(minIdeal, maxIdeal)
	maxBabyWeight := idealWeight / 4
	if maxBabyWeight > minIdeal {
		maxBabyWeight = minIdeal
	}
	currentWeight := events.Roll(1, maxBabyWeight)
	o.weightG = &weight{
		current: currentWeight,
		ideal:   idealWeight,
	}
}

func (o *organ) getRandomHealth(max int) int {
	if max == -1 {
		max = 100
	}
	return events.Roll(0, max)
}

func (o *organ) grow() {

}
