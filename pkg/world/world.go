package world

import human2 "aah/pkg/human"

type World struct {
	human *human2.Human
	year int
}

var W *World

func New() *World{
	return &World{
		human: human2.New(),
		year:  2021,
	}
}

func (w *World) Run() {
	for !w.human.IsDead() {
		w.tick()
	}
}

func (w *World) tick() {
	w.human.Tick()
}