package events

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRoll(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		min int
		max int
	}{
		{
			min: 0,
			max: 100,
		},
		{
			min: -200,
			max: 230,
		},
		{
			min: -100,
			max: -5,
		},
		{
			min: -5,
			max: -4,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run("TestRoll", func(t *testing.T) {
			t.Run(fmt.Sprintf("min: %d, max: %d", tc.min, tc.max), func(t *testing.T) {
				t.Parallel()
				gotRand := Roll(tc.min, tc.max)
				require.True(t, tc.min <= gotRand && tc.max > gotRand)
			})
		})

	}
}
