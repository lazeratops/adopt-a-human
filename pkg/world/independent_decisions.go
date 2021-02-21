package world

import (
	"aah/pkg/human"
	"aah/pkg/util"
	"fmt"
	"github.com/gookit/color"
)

type independentDecision struct {
	minAge int
	maxAge int
	called bool
	description string
	choices choices
	decide func(h *human.Human)
}

var independentDecisions []*independentDecision


func initIndependentDecisions() {
	d1 := independentDecision{
		description: "Should I follow my passion or work a boring job for money?",
		minAge: 5,
		maxAge: -1,
		choices: choices{
			"I will follow my passion!": func(h *human.Human) {
				reportHumanResultHeading("I will follow my passion!")
				mind := h.Mind()
				subFromMentalPropertyHumanDecision(mind.Stress, mind.Maturity.Current, 5,26, "I joined a band.")
			},
			"I've gotta make that money.": func(h *human.Human) {
				reportHumanResultHeading("I've gotta make that money.")
				mind := h.Mind()
				addToMentalPropertyHumanDecision(mind.Stress, mind.Maturity.Current, 5,26, "I joined a bank.")
				addToMentalPropertyHumanDecision(mind.Resilience, mind.Maturity.Current, 0,16, "I need to stay strong...")
			},
		},
	}
	d1.decide = func(h *human.Human) {
		color.White.Println(d1.description)
		mind := h.Mind()
		if mind.Agreeableness.Current > mind.Stubbornness.Current {
			d1.choices["I've gotta make that money."](h)
			return
		}
		d1.choices["I will follow my passion!"](h)
	}

	independentDecisions = append(independentDecisions, &d1)
}

func subFromMentalPropertyHumanDecision(m *human.MentalProperty, maturityPerc util.Percent, minPerc, maxPerc int, context string) {
	reportHumanResult(context)
	perc := util.Percent(util.Roll(minPerc, maxPerc))
	m.SubCurrent(util.WhatIsPercentOf(perc, m.Current))
	if maturityPerc < 100 {
		m.SubBase(util.WhatIsPercentOf(perc, m.Base))
		reportHumanResult(fmt.Sprintf("Their current and base %s went down by %s", m.Name, perc))
		return
	}
	reportHumanResult(fmt.Sprintf("Their current %s went down by %s", m.Name, perc))
}

func addToMentalPropertyHumanDecision(m *human.MentalProperty, maturityPerc util.Percent, minPerc, maxPerc int, context string) {
	reportHumanResult(context)
	perc := util.Percent(util.Roll(minPerc, maxPerc))
	m.AddCurrent(util.WhatIsPercentOf(perc, m.Current))
	if maturityPerc < 100 {
		m.AddBase(util.WhatIsPercentOf(perc, m.Base))
		reportHumanResult(fmt.Sprintf("Their current and base %s went up by %s", m.Name, perc))
		return
	}
	reportHumanResult(fmt.Sprintf("Their current %s went up by %s", m.Name, perc))
}

func reportHumanResultHeading(result string) {
	color.Question.Println(result)
}

func reportHumanResultSubHeading(result string) {
	color.White.Println(result)
}

func reportHumanResult(result string) {
	color.Gray.Println(result)
}