package world

import (
	"aah/pkg/human"
	"aah/pkg/util"
	"fmt"
)

// maybe decisions should be based on current state of the human...
type decision struct {
	query   string
	choices choices
}

type choices map[string]func(h *human.Human)

const resultStr = "As a result..."

var decisions = []decision{
	{
		query: "Your human has crawled into a hole. Do you help them get out?",
		choices: choices{
			"Yes! Help them out of the hole.": func(h *human.Human) {
				reportResult(fmt.Sprintf("You helped your human out of the hole. %s", resultStr))
				// Baseline resilience goes down by 1-5%, but no physical damage
				downPerc := util.Roll(1, 6)
				mind := h.Mind()
				down := util.WhatIsPercentOf(util.Percent(downPerc), mind.Resilience.Base)
				mind.Resilience.Base -= down
				reportResult(fmt.Sprintf("They suffered no immediate physical damage. However, their base resilience went down by %d", down))
			},

			"No, crawling out on their own will teach them mental toughness.": func(h *human.Human) {
				reportResult(fmt.Sprintf("You left your human in the hole. %s", resultStr))

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

				// Max immunity goes up a little, maybe
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
		query: "Your human kicked a puppy. How do you react?",
		choices: choices{
			"With a stern lecture about how to treat other animals.": func(h *human.Human) {
				reportResult(fmt.Sprintf("You gave your human a stern lecture about how to treat other animals. %s", resultStr))
				// Baseline kindenss goes up by 1-10%
				perc := util.Roll(1, 11)
				mind := h.Mind()
				res := util.WhatIsPercentOf(util.Percent(perc), mind.Kindness.Base)
				mind.Kindness.Base += res
				reportResult(fmt.Sprintf("Their base kindness increased by %v", util.Percent(perc)))
			},

			"With capital punishment.": func(h *human.Human) {
				reportResult(fmt.Sprintf("You physically punished your human. There's no animal beating allowed around here! %s", resultStr))
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
		},
	},
}

func reportResult(result string) {
	fmt.Println(colorYellow, result)
	fmt.Printf(colorReset)
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
