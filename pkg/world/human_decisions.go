package world

import (
	"aah/pkg/human"
	"aah/pkg/util"
)

func (d *decision) decideIndependently() (func(h *human.Human), error) {
	var selections []string
	for label, _ := range d.choices {
		selections = append(selections, label)
	}
	idx := util.Roll(0, len(selections))
	return d.choices[selections[idx]], nil
}