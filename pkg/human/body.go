package human

import "aah/pkg/util"

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
	body := &body{
		immunity: generateImmunity(),
		height:   generateHeight(),
		weightKg: generateWeight(),
		maturity: generateMaturity(),
	}
	organs := generateOrgans(body)
	body.organs = organs
	return body
}

func generateImmunity() *immunity {
	max := util.Roll(11, 100)
	current := util.Roll(10, max)

	baseAttrition := util.Roll(minImmunityAttrition, maxImmunityAttrition)
	return &immunity{
		currentPercentage: current,
		maxPercentage:     max,
		attritionModifier: 0,
		baseAttrition:     baseAttrition,
	}
}

func (b *body) tick() {
	cMaturity := b.maturity.current
	for _, organ := range b.organs {
		if cMaturity < 100 {
			organ.grow()
		}
	}
	b.maturity.current += b.maturity.currentRate()
	if b.maturity.current >= 100 {
		b.maturity.rateModifier = -b.maturity.baseRate
	}

}
