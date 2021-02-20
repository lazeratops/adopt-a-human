package human

import (
	"math/rand"
	"time"
)

type Human struct {
	Name string
	Age  int
	mind *mind
	body *body
}

func New() *Human {
	rand.Seed(time.Now().UTC().UnixNano())
	return &Human{
		Name: "",
		Age:  0,
		mind: &mind{},
		body: generateBody(),
	}
}

func (h *Human) Tick() {
	h.Age++
	h.body.tick()
}

func (h *Human) IsDead() bool {
	for _, o := range h.body.organs {
		if o.currentHealth <= 0 {
			return true
		}
	}
	return false
}
