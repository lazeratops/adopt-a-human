package human

import "aah/pkg/util"

const (
	minImmunityAttrition = 0
	maxImmunityAttrition = 50
)

type Immunity struct {
	Current           util.Percent
	Max               util.Percent
	AttritionModifier int
	BaseAttrition     int
}

func generateImmunity() *Immunity {
	max := util.Roll(11, 100)
	current := util.Roll(10, max)

	baseAttrition := util.Roll(minImmunityAttrition, maxImmunityAttrition)
	return &Immunity{
		Current:           util.Percent(current),
		Max:               util.Percent(max),
		AttritionModifier: 0,
		BaseAttrition:     baseAttrition,
	}
}

func (i *Immunity) tick(currentMaturity util.Percent) {
	// While the human is growing, the immunity is stabilizing.
	if currentMaturity < 100 {

	}

}
