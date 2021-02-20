package human

import (
	"aah/pkg/util"
	"fmt"
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
			require.True(t, o.baseGrowthRate > 0)
			require.True(t, o.growthRateModifier > o.baseGrowthRate)
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
			t.Run(fmt.Sprintf("body base growth rate: %d, ideal weight: %d", tc.bodyBaseGrowthRate, tc.idealWeightG), func(t *testing.T) {
				organ := &organ{}
				organ.generateAndSetGrowthRate(tc.bodyBaseGrowthRate, tc.idealWeightG)
				require.EqualValues(t, tc.wantOrganBaseGrowthRate, organ.baseGrowthRate)
				require.EqualValues(t, tc.wantOrganGrowthModifier, organ.growthRateModifier)
			})
		}

	})

}

func TestOrganGrowth(t *testing.T) {
	h := New()
	organ := h.body.organs[0]
	require.True(t, organ.baseGrowthRate > 0)
	require.True(t, organ.growthRateModifier > 0)
	require.True(t, h.body.maturity.current == 0)
	for i := 0; i < 100; i++ {
		h.Tick()
		fmt.Printf("age: %d, ideal weight: %d, current: %d, baseGrowthRate: %d, growthRateModifier: %d, maturity: %d\n", h.Age, organ.weightG.ideal, organ.weightG.current, organ.baseGrowthRate, organ.growthRateModifier, h.body.maturity.current)
	}
}
