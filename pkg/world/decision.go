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
				down := util.WhatIsPercentOf(util.Percent(downPerc), mind.BaseResilience)
				mind.BaseResilience -= down
				reportResult(fmt.Sprintf("They suffered no immediate physical damage. However, their base resilience went down by %d", down))
			},

			"No, crawling out on their own will teach them mental toughness": func(h *human.Human) {
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
				body.Immunity.Max += util.Percent(immunityInc)
				body.Immunity.Current += util.Percent(immunityInc)
				reportResult(fmt.Sprintf("Their immunity increased by %v", immunityInc))

				// But baseline resilience goes up by 5-10%
				upPerc := util.Roll(5, 11)
				mind := h.Mind()
				up := util.WhatIsPercentOf(util.Percent(upPerc), mind.BaseResilience)
				mind.BaseResilience += up
				reportResult(fmt.Sprintf("Their baseline mental resilience increased by %v", up))
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
