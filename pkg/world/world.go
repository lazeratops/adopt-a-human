package world

import (
	human2 "aah/pkg/human"
	"aah/pkg/util"
	"fmt"
	"github.com/gookit/color"
	"github.com/sirupsen/logrus"
)

type World struct {
	human     *human2.Human
	year      int
	watchOnly bool
}

var W *World

func New(watchOnly bool) *World {
	return &World{
		human:     human2.New(),
		year:      2021,
		watchOnly: watchOnly,
	}
}

func (w *World) SetHumanName(name string) {
	w.human.Name = name
}

func (w *World) Run() {
	for {
		w.tick()
		if !w.watchOnly {
			decision := pickDecision(w.human.Age)
			if decision != nil {
				consequence, err := decision.decide()
				if err != nil {
					logrus.WithError(err).Fatalf("failed to get decision: %v", err)
				}
				consequence(w.human)
			}
		}
		if w.human.IsDead() {
			break
		}
		w.stateReport()
	}
	fmt.Println()
	color.Red.Println("~~~~~~~~~~~~~~~~~~~")
	color.Red.Println("~~~~~GAME OVER~~~~~")
	color.Red.Println("~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("\nYour human died at %d years of age", w.human.Age)
	fmt.Printf("\nCause Of Death: %v", w.human.CauseOfDeath())
	fmt.Println()
}

func (w *World) stateReport() {
	state := w.human.StateReport()
	color.Println(state)
}

func (w *World) tick() {
	w.human.Tick()
}

func pickDecision(humanAge int) *decision {
	if len(decisions) == 0 {
		return nil
	}
	idx := util.Roll(0, len(decisions))
	decision := decisions[idx]
	if decision.called {
		logrus.Debug("This decision has already been called; skipping")
		pickDecision(humanAge)
	}
	if decision.minAge > -1 && humanAge < decision.minAge {
		pickDecision(humanAge)
	}
	if decision.maxAge > -1 && humanAge > decision.maxAge {
		pickDecision(humanAge)
	}
	// Remove decision from list, we don't want one to be repeated (at least for now)
	decisions = append(decisions[:idx], decisions[idx+1:]...)
	return decision
}
