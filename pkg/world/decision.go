package world

import (
	"aah/pkg/human"
	"aah/pkg/util"
	"fmt"
	"github.com/gookit/color"
	"github.com/smartystreets/goconvey/convey/reporting"
)

// maybe decisions should be based on current state of the human...
type decision struct {
	minAge  int
	maxAge  int
	query   string
	choices choices
}

type choices map[string]func(h *human.Human)

const resultStr = "As a result..."

var decisions = []decision{
	{
		minAge: -1,
		maxAge: -1,
		query:  "Your human has crawled into a hole. Do you help them get out?",
		choices: choices{
			"Yes! Help them out of the hole.": func(h *human.Human) {
				reportResultHeading(fmt.Sprintf("You helped your human out of the hole. %s", resultStr))
				// Baseline resilience goes down by 1-5%, but no physical damage
				reportResult(fmt.Sprintf("They suffered no immediate physical damage."))
				body := h.Body()

				recoverAllOrgans(body, 1, 11)

				downPerc := util.Roll(1, 6)
				mind := h.Mind()
				down := util.WhatIsPercentOf(util.Percent(downPerc), mind.Resilience.Base)
				mind.Resilience.Base -= down

				reportResult(fmt.Sprintf("However, their base resilience went down by %d", down))
			},

			"No, crawling out on their own will teach them mental toughness.": func(h *human.Human) {
				reportResultHeading(fmt.Sprintf("You left your human in the hole. %s", resultStr))

				body := h.Body()
				// Roll an extra damage check
				damageCheck := util.Roll(0, 100)
				if damageCheck > 60 {
					hitRoll := util.Roll(0, 5)
					// Pick a random organ
					organIdx := util.Roll(0, len(body.Organs))
					organ := body.Organs[organIdx]
					organ.SubHealth(hitRoll)
					reportResult(fmt.Sprintf("They suffered %d damage to their %s", hitRoll, organ.Name()))
				}

				// Ideal immunity goes up a little, maybe
				immunityInc := util.Roll(0, 10)
				body.Immunity.Max += immunityInc
				body.Immunity.Current += immunityInc
				reportResult(fmt.Sprintf("Their immunity increased by %v", immunityInc))

				// But baseline resilience goes up by 5-10%
				upPerc := util.Roll(5, 11)
				mind := h.Mind()
				up := util.WhatIsPercentOf(util.Percent(upPerc), mind.Resilience.Base)
				mind.Resilience.Base += up
				reportResult(fmt.Sprintf("Their baseline mental resilience increased by %v", up))
			},
		},
	},
	{
		minAge: -1,
		maxAge: -1,
		query:  "Your human kicked a puppy. How do you react?",
		choices: choices{
			"With a stern lecture about how to treat other animals.": func(h *human.Human) {
				reportResultHeading(fmt.Sprintf("You gave your human a stern lecture about how to treat other animals. %s", resultStr))
				// Baseline kindenss goes up by 1-10%
				perc := util.Roll(1, 11)
				mind := h.Mind()
				res := util.WhatIsPercentOf(util.Percent(perc), mind.Kindness.Base)
				mind.Kindness.Base += res
				reportResult(fmt.Sprintf("Their base kindness increased by %v", util.Percent(perc)))
			},

			"With capital punishment.": func(h *human.Human) {
				reportResultHeading(fmt.Sprintf("You physically punished your human. There's no animal beating allowed around here! %s", resultStr))
				mind := h.Mind()

				// Baseline kindenss goes up by 1-5%
				perc := util.Roll(1, 6)
				res := util.WhatIsPercentOf(util.Percent(perc), mind.Kindness.Base)
				mind.Kindness.Base += res
				reportResult(fmt.Sprintf("Their base kindness increased by %v", util.Percent(perc)))

				// Current stress goes up by 0-10%
				perc = util.Roll(0, 11)
				res = util.WhatIsPercentOf(util.Percent(perc), mind.Stress.Base)
				mind.Stress.Current += res
				reportResult(fmt.Sprintf("Their stress increased by %v", util.Percent(perc)))

				// If their resilience is lower then their stress, there are more consequences
				if mind.Resilience.Current < mind.Stress.Current {
					body := h.Body()

					// Base immunity goes down by 0-5% (chronic stress effect)
					perc = util.Roll(0, 5)
					body.Immunity.Max -= util.WhatIsPercentOf(util.Percent(perc), body.Immunity.Max)
					if body.Immunity.Max < body.Immunity.Current {
						body.Immunity.Current = body.Immunity.Max
					}
					reportResult(fmt.Sprintf("The stress of the punishment was too much for them. Their immunity decreased by %v", util.Percent(perc)))
				}
			},

			"You kick the puppy with them.": func(h *human.Human) {
				reportResultHeading(fmt.Sprintf("You decided to join the fun and kick the puppy as well. You're both pretty messed up... %s", resultStr))
				mind := h.Mind()

				// Baseline and current kindenss goes down by 5-10%
				perc := util.Roll(5, 11)
				res := util.WhatIsPercentOf(util.Percent(perc), mind.Kindness.Base)
				mind.Kindness.SubBase(res)
				mind.Kindness.SubCurrent(res)
				reportResult(fmt.Sprintf("Their kindness decreased by %v", util.Percent(perc)))
			},
		},
	},
	{
		minAge: -1,
		maxAge: 15,
		query:  "Your human won't eat.",
		choices: choices{
			"Try to coax them into eating nicely.": func(h *human.Human) {
				reportResultHeading(fmt.Sprintf("You did your best to coax them into eating something. %s", resultStr))
				mind := h.Mind()
				if mind.Agreeableness.Current > mind.Stubbornness.Current {
					// It worked! All organs get an extra 1-10% health
					body := h.Body()
					recoverAllOrgans(body, 1, 11)

					// Resilience has a chance of going down
					perc := util.Roll(0, 5)
					toSub := util.WhatIsPercentOf(util.Percent(perc), mind.Resilience.Current)
					mind.Resilience.SubCurrent(toSub)
					mind.Resilience.SubBase(toSub)
					reportResult(fmt.Sprintf("Their resilience decreased by %v", util.Percent(perc)))

					return
				} else {
					// If they are more stubborn than agreeable, organ health has a chance of going down
					body := h.Body()
					damageAllOrgans(body, 1, 11)

					// Kindness has a chance of going up
					perc := util.Roll(0, 15)
					mod := util.WhatIsPercentOf(util.Percent(perc), mind.Kindness.Current)
					mind.Kindness.AddBase(mod)
					mind.Kindness.AddCurrent(mod)
					reportResult(fmt.Sprintf("Their kindness increased by %v", util.Percent(perc)))
				}
			},
			"Ignore them. They'll get over it.": func(h *human.Human) {
				reportResultHeading(fmt.Sprintf("You ignored the tantrum. They'll eat when they're hungry %s", resultStr))
				// Let's see how long it takes them to eat based on their stubbornness
				mind := h.Mind()
				body := h.Body()
				diff := mind.Stubbornness.Current - mind.Agreeableness.Current
				if diff <= 0 {
					reportResult(fmt.Sprintf("As you suspected, they eat right away. They're not that stubborn."))
					recoverAllOrgans(body, 1, 11)

					perc := util.Roll(0, 5)
					toSub := util.WhatIsPercentOf(util.Percent(perc), mind.Resilience.Current)
					mind.Resilience.SubCurrent(toSub)
					mind.Resilience.SubBase(toSub)
					reportResult(fmt.Sprintf("Their resilience decreased by %v", util.Percent(perc)))
				} else {
					var originalHealths []struct {
						name   string
						health int
					}
					for _, o := range body.Organs {
						originalHealths = append(originalHealths, struct {
							name   string
							health int
						}{name: o.Name(), health: o.CurrentHealth})
					}
					// They're pretty stubborn...they hold out for the amount of difference, organs getting damaged each turn
					for i := 0; i < diff; i++ {
						reportResult(fmt.Sprintf("They're still holding out..."))
						damageAllOrgans(body, 0, 6)
					}
					for _, o := range body.Organs {
						for i, originalO := range originalHealths {
							if o.Name() == originalO.name {
								reportResult(fmt.Sprintf("Their %s health decreased by %v", o.Name(), util.GetPercent(o.CurrentHealth, originalO.health)))
								originalHealths = append(originalHealths[:i], originalHealths[i+1:]...)
								break
							}
						}
					}

					// Their resilience has a chance of going up...if they survive
					perc := util.Roll(0, 15)
					toSub := util.WhatIsPercentOf(util.Percent(perc), mind.Resilience.Current)
					mind.Resilience.AddCurrent(toSub)
					mind.Resilience.AddBase(toSub)
				}
			},
		},
	},
	{
		minAge: -1,
		maxAge: -1,
		query:  "You caught your human eating sand. What do you do?",
		choices: choices{
			"Yell at them.": func(h *human.Human) {
				reportResultHeading(fmt.Sprintf("You yelled at them. What kind of an idiot eats sand?! %s", resultStr))
				mind := h.Mind()
				body := h.Body()
				if mind.Stubbornness.Current > mind.Agreeableness.Current {
					reportResult("They dug their heels in and ate even more sand in protest.")
					// their stomach might get a bit damaged from all the sand.
					for _, o := range body.Organs {
						if o.Kind == human.OrganStomacch {
							hitRoll := util.Roll(0, 100)
							immunityPerc := util.GetPercent(body.Immunity.Current, body.Immunity.Max)
							if hitRoll > int(immunityPerc) {
								damagePercRoll := util.Percent(util.Roll(0, 5))
								damage := util.WhatIsPercentOf(damagePercRoll, o.CurrentHealth)
								o.SubHealth(damage)
								reportResult(fmt.Sprintf("They got a stomach-ache; their stomach was damaged by %v", damagePercRoll))
							}
							break
						}
					}
					// If they are not fully mature yet, their max immunity goes up and their immunity attrition goes down for a while
					if body.Maturity.Current < 100 {
						perc := util.Percent(util.Roll(5, 10))
						mod := util.WhatIsPercentOf(perc, body.Immunity.Max)
						body.Immunity.Max += mod
						attrMod := util.WhatIsPercentOf(perc, body.Immunity.BaseAttrition)
						body.Immunity.AddToAttritionModifier(-attrMod, 5)
						reportResult(fmt.Sprintf("Eating all that sand helped reinforce their maximum immunity by %s, and decreased their immunity attrition for a whole.", perc))
					}
					return
				}
				reportResult("They scurry away. Phew - you stopped them before they ate too much.")
				if body.Maturity.Current < 100 {
					perc := util.Percent(util.Roll(0, 6))
					mod := util.WhatIsPercentOf(perc, body.Immunity.Max)
					body.Immunity.Max += mod
					attrMod := util.WhatIsPercentOf(perc, body.Immunity.BaseAttrition)
					body.Immunity.AddToAttritionModifier(-attrMod, 2)
					reportResult(fmt.Sprintf("Eating a little sand helped reinforce their maximum immunity by %s, and decreased their immunity attrition for a whiole.", perc))
				}
				// But their stress goes up from the yelling
				perc := util.Percent(util.Roll(0, 11))
				res := util.WhatIsPercentOf(perc, mind.Stress.Base)
				mind.Stress.Current += res
				reportResult(fmt.Sprintf("All that yelling stressed them out. Their stress increased by %s of the base.", perc))
				return
			},
			"Do nothing. Who are you to judge their dietary choices?": func(h *human.Human) {
				reportResultHeading(fmt.Sprintf("You did nothing. It's a free country. %s", resultStr))
				mind := h.Mind()
				body := h.Body()
				// their stomach might get a bit damaged from all the sand.
				for _, o := range body.Organs {
					if o.Kind == human.OrganStomacch {
						hitRoll := util.Roll(0, 100)
						immunityPerc := util.GetPercent(body.Immunity.Current, body.Immunity.Max)
						if hitRoll > int(immunityPerc) {
							damagePercRoll := util.Percent(util.Roll(0, 15))
							damage := util.WhatIsPercentOf(damagePercRoll, o.CurrentHealth)
							o.SubHealth(damage)
							reportResult(fmt.Sprintf("They got a stomach-ache; their stomach was damaged by %v", damagePercRoll))
						}
						break
					}
				}

				// Maybe sand is a type of stress relief
				perc := util.Percent(util.Roll(0, 11))
				res := util.WhatIsPercentOf(perc, mind.Stress.Base)
				mind.Stress.Current -= res
				reportResult(fmt.Sprintf("The sand crunching between their teeth was oddly soothing...theri stress went down by %d of the base", perc))
			},
		},
	},
	{
		minAge: -1,
		maxAge: -1,
		query:  "You see your human climbing a very tall tree. What do you do?",
		choices: choices{
			"Let them climb, of course!": func(h *human.Human) {
				reportResultHeading(fmt.Sprintf("You watched them climb proudly, offering occasional words of encouragement. %s", resultStr))
				body := h.Body()
				mind := h.Mind()
				fallRoll := util.Roll(0, 1000)

				if mind.Resilience.Current > mind.Stress.Current {
					// chance of falling is pretty low due to their resilience
					if fallRoll > 800 {
						reportResult("They almost reach the top, but then they slip and fall to the ground!!!")
						damageAllOrgans(body, 0, 5)
					} else {
						reportResult("They get to the top! They breathe the crisp canopy air and snack on some bugs.")
						perc := util.Percent(util.Roll(5, 16))
						body.Immunity.AddToAttritionModifier(-util.WhatIsPercentOf(perc, body.Immunity.Current), 5)
						reportResult(fmt.Sprintf("Their immunity attrition decreased by %s for a few years", perc))
					}

					// Their stress goes up from the fall; if they are not mature yet, their base stress goes up as well
					addToMentalProperty(mind.Stress, mind.Maturity.Current, 0, 16, "They are rattled by the fall.")
					return
				} else {
					// chance of falling is slightly more likely because they are stressed
					if fallRoll > 600 {
						reportResult("They are a bit too stressed to be doing this. They fall from a great height!")
						damageAllOrgans(body, 0, 5)
					}
				}

			},
			"Tell them they're not getting dinner unless they get all the way to the top.": func(h *human.Human) {
				reportResultHeading(fmt.Sprintf("You egged them on. They should finish what they started! %s", resultStr))
				body := h.Body()
				mind := h.Mind()

				// Their resilience goes up
				addToMentalProperty(mind.Resilience, mind.Maturity.Current, 0, 16, "")

				// Their stress goes up

			},
			"Yell at them to come down. It's dangerous and you don't have health insurance!": func(h *human.Human) {
				reportResultHeading(fmt.Sprintf("You yelled at them to come down. %s", resultStr))

			},
		},
	},
	{
		minAge: -1,
		maxAge: -1,
		query:  "You caught your human eating sand. What do you do?",
		choices: choices{
			"Yell at them.": func(h *human.Human) {
				reportResultHeading(fmt.Sprintf("You yelled at them. What kind of an idiot eats sand?! %s", resultStr))

			},
		},
	},
	{
		minAge: -1,
		maxAge: -1,
		query:  "You caught your human eating sand. What do you do?",
		choices: choices{
			"Yell at them.": func(h *human.Human) {
				reportResultHeading(fmt.Sprintf("You yelled at them. What kind of an idiot eats sand?! %s", resultStr))

			},
		},
	},
}

func reportResultHeading(result string) {
	color.BgBlue.Println(result)
}

func reportResult(result string) {
	color.Yellow.Println(result)
}

func (d *decision) decide() (func(h *human.Human), error) {
	var selections []string
	for label, _ := range d.choices {
		selections = append(selections, label)
	}
	selection, err := promptSelection(d.query, selections)
	if err != nil {
		return nil, fmt.Errorf("failed to make a decision: %w", err)
	}
	return d.choices[selection], nil

}

func addToMentalProperty(m *human.MentalProperty, maturityPerc util.Percent, minPerc, maxPerc int, context string) {
	perc := util.Percent(util.Roll(minPerc, maxPerc))
	m.AddCurrent(util.WhatIsPercentOf(perc, m.Current))
	if maturityPerc < 100 {
		m.AddBase(util.WhatIsPercentOf(perc, m.Base))
		reportResult(fmt.Sprintf("%s. Their current and base %s went up by %s", context, m.Name, perc))
		return
	}
	reportResult(fmt.Sprintf("%s. Their current %s went up by %s", context, m.Name, perc))
}

func subFromMentalProperty(m *human.MentalProperty, maturityPerc util.Percent, minPerc, maxPerc int, context string) {
	perc := util.Percent(util.Roll(minPerc, maxPerc))
	m.SubCurrent(util.WhatIsPercentOf(perc, m.Current))
	if maturityPerc < 100 {
		m.SubBase(util.WhatIsPercentOf(perc, m.Base))
		reportResult(fmt.Sprintf("%s. Their current and base %s went down by %s", context, m.Name, perc))
		return
	}
	reportResult(fmt.Sprintf("%s. Their current %s went down by %s", context, m.Name, perc))
}

func damageAllOrgans(b *human.Body, minPerc, maxPerc int) {
	for _, o := range b.Organs {
		perc := util.Roll(minPerc, maxPerc)
		toSub := util.WhatIsPercentOf(util.Percent(perc), o.CurrentHealth)
		o.SubHealth(toSub)
		reportResult(fmt.Sprintf("Their %s health decreased by %v", o.Name(), util.Percent(perc)))
	}
}

func recoverAllOrgans(b *human.Body, minPerc, maxPerc int) {
	for _, o := range b.Organs {
		perc := util.Roll(minPerc, maxPerc)
		toSub := util.WhatIsPercentOf(util.Percent(perc), o.CurrentHealth)
		o.AddHealth(toSub)
		reportResult(fmt.Sprintf("Their %s health recovered by %v", o.Name(), util.Percent(perc)))
	}
}
