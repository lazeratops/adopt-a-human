package human

import "aah/pkg/events"

const (
	minImmunityAttrition = -50
	maxImmunityAttrition = 50
)
type body struct {
	immunity *immunity
	organs []*organ
	height *heightCM
	weightKg *weight
	maturity *maturity
}

type immunity struct {
	currentPercentage int
	maxPercentage               int
	attritionModifier int
	baseAttrition     int
}

func generateBody() *body {
	organs := generateOrgans()

	return &body{
		immunity: generateImmunity(),
		height: generateHeight(),
		weightKg: generateWeight(),
		maturity: generateMaturity(),
		organs:          organs,
	}
}

func generateImmunity() *immunity {
	max := events.Roll(10, 100)
	current := events.Roll(10, max)

	minAttrition := events.Roll(minImmunityAttrition, 0)
	maxAttrition := events.Roll(0, maxImmunityAttrition)

	realMaxAttrition := events.Roll(minAttrition, maxAttrition)
	baseAttrition := events.Roll(minAttrition, realMaxAttrition)
	return &immunity{
		currentPercentage: current,
		maxPercentage:               max,
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

