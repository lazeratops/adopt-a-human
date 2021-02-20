package human

import (
	"github.com/stretchr/testify/require"

	"testing"
)

func TestGenerateHuman(t *testing.T) {
	for i := 0; i < 100; i++ {
		h := New()
		body := h.body
		require.True(t, body.immunity.currentPercentage <= body.immunity.currentPercentage)
		for _, o := range body.organs {
			require.True(t, o.currentHealth <= o.maxHealth)
			require.True(t, o.weightG.current <= o.weightG.ideal)
		}
	}
}
