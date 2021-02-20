package human

type Human struct {
	Name string
	Age int
	mind *mind
	body *body
}

func New() *Human {
	return &Human{
		Name:            "",
		Age:             0,
		mind:            &mind{},
		body:            generateBody(),
	}
}

func (h *Human) Tick() {
	h.Age++
}

func (h *Human) IsDead() bool {
	for _, o := range h.body.organs {
		if o.currentHealth <= 0 {
			return true
		}
	}
	return false
}
