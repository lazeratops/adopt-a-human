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
	heart.generateAndSetWeight(100, 500)

	brain := organ{
		kind:     OrganBrain,
		quantity: 1,
	}
	brain.generateAndSetHealths()
	brain.generateAndSetWeight(500, 3000)

	kidney := organ{
		kind:     OrganKidneys,
		quantity: 2,
	}
	kidney.generateAndSetHealths()
	kidney.generateAndSetWeight(25, 400)

	lungs := organ{
		kind:     OrganLungs,
		quantity: 2,
	}
	lungs.generateAndSetHealths()
	lungs.generateAndSetWeight(50, 400)
	return []*organ{&heart, &brain, &kidney, &lungs}
}

func (o *organ) generateAndSetHealths() {
	o.maxHealth = o.getRandomHealth(-1)
	o.currentHealth = o.getRandomHealth(o.maxHealth)
}

func (o *organ) generateAndSetWeight(minIdeal int, maxIdeal int) {
	idealWeight := events.Roll(minIdeal, maxIdeal)
	maxBabyWeight := idealWeight / 4
	if maxBabyWeight < minIdeal {
		maxBabyWeight = minIdeal
	}
	currentWeight := events.Roll(1, maxBabyWeight+1)
	o.weightG = &weight{
		current: currentWeight,
		ideal:   idealWeight,
	}
}

func (o *organ) getRandomHealth(max int) int {
	if max == -1 {
		max = 100
	}
	return events.Roll(0, max+1)
}

func (o *organ) grow() {

}
