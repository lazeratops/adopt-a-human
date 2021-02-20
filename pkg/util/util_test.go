package util

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWhatIsPercentOf(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		p     Percent
		total int
		want  int
	}{
		{
			p:     50,
			total: 100,
			want:  50,
		},
		{
			p:     15,
			total: 18,
			want:  3,
		},
		{
			p:     100,
			total: 0,
			want:  0,
		},
	}

	t.Run("TestWhatIsPercentOf", func(t *testing.T) {
		for _, tc := range testCases {
			tc := tc
			t.Run(fmt.Sprintf("p: %d, total: %d", tc.p, tc.total), func(t *testing.T) {
				gotRes := WhatIsPercentOf(tc.p, tc.total)
				require.Equal(t, tc.want, gotRes)
			})
		}
	})

}
