package human

import (
	"aah/pkg/util"
	"encoding/json"
	"fmt"
	"github.com/gookit/color"
	log "github.com/sirupsen/logrus"
	"strings"
)

type Human struct {
	Name string
	Age  int
	mind *Mind
	body *Body
}

func New() *Human {
	return &Human{
		Name: "",
		Age:  0,
		mind: generateMind(),
		body: generateBody(),
	}
}

func (h *Human) Tick() {
	h.body.tick()
	h.mind.tick()
	h.Age++
	j, err := json.Marshal(h)
	if err != nil {
		log.Error(err)
		return
	}
	log.Trace(string(j))
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
		if o.CurrentHealth <= 0 && !h.organFallbackExists(o) {
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
			if o.CurrentHealth <= 0 {
				causes = append(causes, fmt.Sprintf("%s failure", o.kind))
			}
		}
	}
	return strings.Join(causes, ", ")
}

func (h* Human) StateReport() string {
	state := fmt.Sprintf("%s is %d years old.", h.Name, h.Age)

	body := h.body
	mind := h.mind

	state += fmt.Sprintf(" Their body is %s.", body.maturity.Descriptor())
	state += fmt.Sprintf(" Their mind is %s.", mind.Maturity.Descriptor())
	state += fmt.Sprintf(" they weigh %d kg.", body.weightKg.Current)

	immunity := body.Immunity
	immunityPerc := util.GetPercent(immunity.Current, immunity.Max)
	state += fmt.Sprintf(" Their immunity is %v of their total capacity.", immunityPerc)
	state += fmt.Sprintf(" %s",mind.StateReport())
	for _, o := range body.Organs {
		healthPerc := util.GetPercent(o.CurrentHealth, o.maxHealth)
		if healthPerc < 50 {
			state += color.Yellow.Sprintf("\nTheir %s is %s.", o.Name(), o.Descriptor())
		}
	}
	return state
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
