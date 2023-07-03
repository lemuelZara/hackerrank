package warmup

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJumpingOnClouds(t *testing.T) {
	t.Run("should jump on the clouds correctly", func(t *testing.T) {
		tests := [][]int32{
			{JumpingOnClouds([]int32{0, 0, 0, 0, 1, 0}), 3},
			{JumpingOnClouds([]int32{0, 0, 1, 0, 0, 1, 0}), 4},
			{JumpingOnClouds([]int32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}), 7},
			{JumpingOnClouds([]int32{0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0}), 6},
			{JumpingOnClouds([]int32{0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0}), 9},
			{JumpingOnClouds([]int32{0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0}), 8},
		}

		for _, v := range tests {
			result := v[0]
			expected := v[1]

			require.Equal(t, expected, result)
		}
	})
}
