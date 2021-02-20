package human

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"strings"
	"time"
)

type Human struct {
	Name string
	Age  int
	mind *Mind
	body *Body
}

func New() *Human {
	rand.Seed(time.Now().UTC().UnixNano())
	return &Human{
		Name: "",
		Age:  0,
		mind: generateMind(),
		body: generateBody(),
	}
}

func (h *Human) Tick() {
	h.body.tick()
	h.Age++
	j, err := json.Marshal(h)
	if err != nil {
		log.Error(err)
		return
	}
	log.Debug(string(j))
}
func (h *Human) Body() *Body {
	return h.body
}

func (h *Human) Mind() *Mind {
	return h.mind
}
func (h *Human) IsDead() bool {
	for _, o := range h.body.Organs {
		// If Organ is dead, see if there is another version of this Organ available to fall back on. If not, die.
		if o.currentHealth <= 0 && !h.organFallbackExists(o) {
			return true
		}
	}
	return false
}

func (h *Human) CauseOfDeath() string {
	var causes []string
	if h.IsDead() {
		for _, o := range h.body.Organs {
			// If Organ is dead, see if there is another version of this Organ available to fall back on. If not, die.
			if o.currentHealth <= 0 {
				causes = append(causes, fmt.Sprintf("%s failure", o.kind))
			}
		}
	}
	return strings.Join(causes, ", ")
}

func (h *Human) organFallbackExists(o *Organ) bool {
	// See if there is another version of this Organ available to fall back on...
	for _, otherO := range h.body.Organs {
		if otherO == o {
			continue
		}
		if otherO.kind == o.kind {
			return true
		}
	}
	return false
}

func (h *Human) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Name string
		Age  int
		Mind *Mind
		Body *Body
	}{
		Name: h.Name,
		Age:  h.Age,
		Mind: h.mind,
		Body: h.body,
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}
