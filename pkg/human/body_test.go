package human

import (
	"aah/pkg/util"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"

	"testing"
)

func TestGenerateHuman(t *testing.T) {
	for i := 0; i < 100; i++ {
		h := New()
		body := h.body
		require.True(t, body.Immunity.Current <= body.Immunity.Max)
		for _, o := range body.Organs {
			require.True(t, o.CurrentHealth > 0)
			require.True(t, o.CurrentHealth <= o.maxHealth)
			require.True(t, o.weightG.Current <= o.weightG.Ideal)
			require.True(t, o.baseGrowthRate > 0)
		}
	}
}

func TestOrganGrowthRateGeneration(t *testing.T) {
	testCases := []struct {
		bodyBaseGrowthRate      util.Percent
		idealWeightG            int
		wantOrganBaseGrowthRate int
		wantOrganGrowthModifier int
	}{
		{
			bodyBaseGrowthRate:      2,
			idealWeightG:            100,
			wantOrganBaseGrowthRate: 2,
			wantOrganGrowthModifier: 1,
		},
		{
			bodyBaseGrowthRate:      15,
			idealWeightG:            256,
			wantOrganBaseGrowthRate: 38,
			wantOrganGrowthModifier: 19,
		},
	}
	t.Run("TestOrganGrowthRateGeneration", func(t *testing.T) {
		for _, tc := range testCases {
			tc := tc
			t.Run(fmt.Sprintf("Body Base growth rate: %d, Ideal Weight: %d", tc.bodyBaseGrowthRate, tc.idealWeightG), func(t *testing.T) {
				organ := &Organ{}
				organ.generateAndSetGrowthRate(tc.bodyBaseGrowthRate, tc.idealWeightG)
				require.EqualValues(t, tc.wantOrganBaseGrowthRate, organ.baseGrowthRate)
				require.EqualValues(t, tc.wantOrganGrowthModifier, organ.growthRateModifier)
			})
		}

	})

}

func TestOrganGrowth(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	h := New()
	organ := h.body.Organs[0]
	require.True(t, organ.baseGrowthRate > 0)
	require.True(t, organ.growthRateModifier > 0)
	require.True(t, h.body.Maturity.Current == 0)
	for !h.IsDead() {
		h.Tick()
	}
}
