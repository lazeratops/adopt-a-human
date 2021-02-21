package world

import (
	"aah/pkg/human"
	"aah/pkg/util"
	"fmt"
	"github.com/gookit/color"
)

// maybe decisions should be based on current state of the human...
type decision struct {
	minAge int
	maxAge int
	query   string
	choices choices
}

type choices map[string]func(h *human.Human)

const resultStr = "As a result..."

var decisions = []decision{
	{
		minAge: -1,
		maxAge: -1,
		query: "Your human has crawled into a hole. Do you help them get out?",
		choices: choices{
			"Yes! Help them out of the hole.": func(h *human.Human) {
				reportResultHeading(fmt.Sprintf("You helped your human out of the hole. %s", resultStr))
				// Baseline resilience goes down by 1-5%, but no physical damage
				reportResult(fmt.Sprintf("They suffered no immediate physical damage."))
				body := h.Body()

				for _, o := range body.Organs {
					perc := util.Roll(1, 10)
					toAdd := util.WhatIsPercentOf(util.Percent(perc), o.CurrentHealth)
					if toAdd == 0 {
						toAdd = 1
					}
					o.AddHealth(toAdd)
					reportResult(fmt.Sprintf("Their %s health increased by %v", o.Name(), util.Percent(perc)))
				}

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
		query: "Your human kicked a puppy. How do you react?",
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
					for _, o := range body.Organs {
						perc := util.Roll(1, 11)
						toAdd := util.WhatIsPercentOf(util.Percent(perc), o.CurrentHealth)
						if toAdd == 0 {
							toAdd = 1
						}
						o.AddHealth(toAdd)
						reportResult(fmt.Sprintf("Their %s health increased by %v", o.Name(), util.Percent(perc)))
					}

					// Resilience has a chance of going down
					perc := util.Roll(0,5)
					toSub := util.WhatIsPercentOf(util.Percent(perc), mind.Resilience.Current)
					mind.Resilience.SubCurrent(toSub)
					mind.Resilience.SubBase(toSub)
					reportResult(fmt.Sprintf("Their resilience decreased by %v", util.Percent(perc)))

					return
				} else {
					// If they are more stubborn than agreeable, organ health has a chance of going down
					body := h.Body()
					reportResult("They were very stubborn and your coaxing did not work.")
					for _, o := range body.Organs {
						perc := util.Roll(1, 11)
						toSub := util.WhatIsPercentOf(util.Percent(perc), o.CurrentHealth)
						o.SubHealth(toSub)
						reportResult(fmt.Sprintf("Their %s health decreased by %v", o.Name(), util.Percent(perc)))
					}

					// Kindness has a chance of going up
					perc := util.Roll(0,15)
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
					for _, o := range body.Organs {
						perc := util.Roll(1, 11)
						toAdd := util.WhatIsPercentOf(util.Percent(perc), o.CurrentHealth)
						if toAdd == 0 {
							toAdd = 1
						}
						o.AddHealth(toAdd)
						reportResult(fmt.Sprintf("Their %s health increased by %v", o.Name(), util.Percent(perc)))
					}

					perc := util.Roll(0,5)
					toSub := util.WhatIsPercentOf(util.Percent(perc), mind.Resilience.Current)
					mind.Resilience.SubCurrent(toSub)
					mind.Resilience.SubBase(toSub)
					reportResult(fmt.Sprintf("Their resilience decreased by %v", util.Percent(perc)))
				} else {
					// They're pretty stubborn...they hold out for the amount of difference, organs getting damaged each turn
					for i := 0; i < diff; i++ {
						reportResult("They are pretty stubborn and they're still holding out...")
						for _, o := range body.Organs {
							perc := util.Roll(0, 5)
							toSub := util.WhatIsPercentOf(util.Percent(perc), o.CurrentHealth)
							o.SubHealth(toSub)
							reportResult(fmt.Sprintf("Their %s health decreased by %v", o.Name(), util.Percent(perc)))
						}
					}
					// Their resilience has a chance o going up...if they survive
					perc := util.Roll(0,15)
					toSub := util.WhatIsPercentOf(util.Percent(perc), mind.Resilience.Current)
					mind.Resilience.AddCurrent(toSub)
					mind.Resilience.AddBase(toSub)
				}
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
