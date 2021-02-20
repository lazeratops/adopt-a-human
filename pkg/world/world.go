package world

import (
	human2 "aah/pkg/human"
	"aah/pkg/util"
	"fmt"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

type World struct {
	human *human2.Human
	year  int
}

var W *World

func New() *World {
	return &World{
		human: human2.New(),
		year:  2021,
	}
}

func (w *World) SetHumanName(name string) {
	w.human.Name = name
}

func (w *World) Run() {
	for {
		w.tick()
		decision := pickDecision()
		consequence, err := decision.decide()
		if err != nil {
			logrus.WithError(err).Fatalf("failed to get decision: %v", err)
		}
		consequence(w.human)
		if w.human.IsDead() {
			break
		}
	}
	fmt.Println()
	fmt.Println("~~~~~~~~~~~~~~~~~~~")
	fmt.Println("~~~~~GAME OVER~~~~~")
	fmt.Println("~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("\nYour human died at %d years of age", w.human.Age)
	fmt.Printf("\nCause Of Death: %v", w.human.CauseOfDeath())
	fmt.Println()
}

func (w *World) tick() {
	w.human.Tick()
}

func pickDecision() decision {
	rand.Seed(time.Now().UTC().UnixNano())
	idx := util.Roll(0, len(decisions))
	return decisions[idx]
}
