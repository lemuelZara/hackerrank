package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var firstHourglass [][]int32 = [][]int32{
	{1, 1, 1, 0, 0, 0},
	{0, 1, 0, 0, 0, 0},
	{1, 1, 1, 0, 0, 0},
	{0, 9, 2, -4, -4, 0},
	{0, 0, 0, -2, 0, 0},
	{0, 0, -1, -2, -4, 0},
}

var secondHourglass [][]int32 = [][]int32{
	{1, -1, 1, -1},
	{-1, 1, -1, 1},
	{1, -1, 1, -1},
	{-1, 1, -1, 1},
}

var thirdHourglass [][]int32 = [][]int32{
	{5, 10, -15, 20, -25, 30},
	{8, -16, 24, -32, 40, -48},
	{-3, 6, -9, 12, -18, 21},
	{7, -14, 28, -35, 42, -49},
	{-2, 4, -6, 8, -10, 12},
	{11, -22, 33, -44, 55, -66},
}

var fourthHourglass [][]int32 = [][]int32{
	{3, -2, 4, -1},
	{2, -3, 1, -4},
	{-1, 4, -2, 3},
	{-4, 1, -3, 2},
}

var fifthHourglass [][]int32 = [][]int32{
	{-5, 6, -7},
	{8, -9, 10},
	{-11, 12, -13},
}

var sixthHourglass [][]int32 = [][]int32{
	{5, -4, 3, -2},
	{-1, 0, -6, 7},
	{9, -8, 11, -10},
}

func Test_TwoDimensionalArray(t *testing.T) {
	t.Run("should return the maximum hourglass sum", func(t *testing.T) {
		tests := [][]int32{
			{TwoDimensionalArray(firstHourglass), 13},
			{TwoDimensionalArray(secondHourglass), 3},
			{TwoDimensionalArray(thirdHourglass), 87},
			{TwoDimensionalArray(fourthHourglass), 7},
			{TwoDimensionalArray(fifthHourglass), -27},
			{TwoDimensionalArray(sixthHourglass), 16},
		}

		for _, v := range tests {
			result := v[0]
			expected := v[1]

			require.Equal(t, expected, result)
		}
	})
}
