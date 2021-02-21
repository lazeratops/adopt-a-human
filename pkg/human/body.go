package human

type Body struct {
	Immunity *Immunity
	Organs   []*Organ
	HeightCM *Height
	weightKg *Weight
	Maturity *Maturity
}

func generateBody() *Body {
	body := &Body{
		Immunity: generateImmunity(),
		HeightCM: generateHeight(),
		weightKg: generateWeight(),
		Maturity: generateMaturity(),
	}
	organs := generateOrgans(body)
	body.Organs = organs
	return body
}

func (b *Body) tick() {
	for _, organ := range b.Organs {
		if b.Maturity.Current < 100 {
			organ.grow()
		}
		organ.tickHealth()
	}
	b.weightKg.tick(b.Maturity.Current, b.Maturity.currentRate())
	b.HeightCM.tick(b.Maturity.Current, b.Maturity.currentRate())
	b.Immunity.tick(b.Maturity.Current)
	b.Maturity.tick()
}
