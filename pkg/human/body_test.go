package human

import (
	"github.com/stretchr/testify/require"

	"testing"
)


func TestGenerateImmunity(t *testing.T) {
	for i := 0; i < 100; i++ {
		immunity := generateImmunity()
		require.True(t, immunity.currentPercentage <= immunity.maxPercentage)
	}
}
