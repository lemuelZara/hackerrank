package warmup

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SalesByMatch(t *testing.T) {
	t.Run("should return the number of pairs", func(t *testing.T) {
		tests := [][]int32{
			{SalesByMatch(9, []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}), 3},
			{SalesByMatch(10, []int32{1, 1, 3, 1, 2, 1, 3, 3, 3, 3}), 4},
			{SalesByMatch(5, []int32{50, 3, 99, 1, 50}), 1},
			{SalesByMatch(14, []int32{82, 82, 10, 20, 31, 651, 9, 2, 2, 10, 910, 723, 8, 287}), 3},
			{SalesByMatch(10, []int32{1, 1, 3, 1, 2, 1, 3, 3, 3, 3}), 4},
		}

		for _, v := range tests {
			result := v[0]
			expected := v[1]

			require.Equal(t, expected, result)
		}
	})
}
