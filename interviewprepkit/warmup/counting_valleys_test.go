package warmup

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_CountingValleys(t *testing.T) {
	t.Run("should return the number of valleys traversed", func(t *testing.T) {
		tests := [][]int32{
			{CountingValleys(1, "UDDDUDUU"), 1},       // 1
			{CountingValleys(1, "UUUDDDDU"), 1},       // 1
			{CountingValleys(1, "UDDU"), 1},           // 1
			{CountingValleys(1, "DUDUDU"), 3},         // 3
			{CountingValleys(1, "DDDUUUUD"), 1},       // 1
			{CountingValleys(1, "DUUUUDDDDDDUUU"), 2}, // 2
		}

		for _, v := range tests {
			result := v[0]
			expected := v[1]

			require.Equal(t, expected, result)
		}
	})
}
