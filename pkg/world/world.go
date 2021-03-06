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
	initIndependentDecisions()
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
		if w.human.IsDead() {
			break
		}
		if !w.watchOnly {
			decision := pickDecision(w.human.Age, 3)
			if decision != nil {
				consequence, err := decision.decide()
				if err != nil {
					logrus.WithError(err).Fatalf("failed to get decision: %v", err)
				}
				consequence(w.human)
			}
		}
		independentDecision := pickIndependentDecision(w.human.Age, 3)
		if independentDecision != nil {
			independentDecision.decide(w.human)
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

func pickDecision(humanAge int, attemptsLeft int) *decision {
	if len(decisions) == 0 || attemptsLeft == 0 {
		return nil
	}
	idx := util.Roll(0, len(decisions))
	decision := decisions[idx]
	if decision.called {
		logrus.Debug("This decision has already been called; skipping")
		return pickDecision(humanAge, attemptsLeft - 1)
	}
	if decision.minAge > -1 && humanAge < decision.minAge {
		return pickDecision(humanAge, attemptsLeft - 1)
	}
	if decision.maxAge > -1 && humanAge > decision.maxAge {
		return pickDecision(humanAge, attemptsLeft - 1)
	}
	// Remove decision from list, we don't want one to be repeated (at least for now)
	decisions = append(decisions[:idx], decisions[idx+1:]...)
	return decision
}

func pickIndependentDecision(humanAge int, attemptsLeft int) *independentDecision {
	if len(independentDecisions) == 0 || attemptsLeft == 0 {
		return nil
	}
	idx := util.Roll(0, len(independentDecisions))
	decision := independentDecisions[idx]
	if decision.called {
		logrus.Debug("This decision has already been called; skipping")
		return pickIndependentDecision(humanAge, attemptsLeft - 1)
	}
	if decision.minAge > -1 && humanAge < decision.minAge {
		return pickIndependentDecision(humanAge, attemptsLeft - 1)
	}
	if decision.maxAge > -1 && humanAge > decision.maxAge {
		return pickIndependentDecision(humanAge, attemptsLeft - 1)
	}
	// Remove decision from list, we don't want one to be repeated (at least for now)
	independentDecisions = append(independentDecisions[:idx], independentDecisions[idx+1:]...)
	return decision
}

