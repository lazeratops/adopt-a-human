package human

type Body struct {
	Immunity *Immunity
	Organs   []*Organ
	HeightCM *Height
	weightKg *Weight
	maturity *Maturity
}

func generateBody() *Body {
	body := &Body{
		Immunity: generateImmunity(),
		HeightCM: generateHeight(),
		weightKg: generateWeight(),
		maturity: generateMaturity(),
	}
	organs := generateOrgans(body)
	body.Organs = organs
	return body
}

func (b *Body) tick() {
	for _, organ := range b.Organs {
		if b.maturity.Current < 100 {
			organ.grow()
		}
		organ.tickHealth()
	}
	b.weightKg.tick(b.maturity.Current, b.maturity.currentRate())
	b.HeightCM.tick(b.maturity.Current, b.maturity.currentRate())
	b.Immunity.tick(b.maturity.Current)
	b.maturity.tick()
}
