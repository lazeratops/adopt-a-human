package human

import "aah/pkg/events"

const (
	minImmunityAttrition = 0
	maxImmunityAttrition = 50
)

type body struct {
	immunity *immunity
	organs   []*organ
	height   *heightCM
	weightKg *weight
	maturity *maturity
}

type immunity struct {
	currentPercentage int
	maxPercentage     int
	attritionModifier int
	baseAttrition     int
}

func generateBody() *body {
	organs := generateOrgans()

	return &body{
		immunity: generateImmunity(),
		height:   generateHeight(),
		weightKg: generateWeight(),
		maturity: generateMaturity(),
		organs:   organs,
	}
}

func generateImmunity() *immunity {
	max := events.Roll(11, 100)
	current := events.Roll(10, max)

	baseAttrition := events.Roll(minImmunityAttrition, maxImmunityAttrition)
	return &immunity{
		currentPercentage: current,
		maxPercentage:     max,
		attritionModifier: 0,
		baseAttrition:     baseAttrition,
	}
}

func (b *body) tick() {
	cMaturity := b.maturity.currentPercent
	for _, organ := range b.organs {
		if cMaturity < 100 {
			_ = organ
		}
	}
	b.maturity.currentPercent += b.maturity.currentRate()

}
