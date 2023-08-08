package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var firstArr []int32 = []int32{1, 2, 3, 4, 5}
var secondArr []int32 = []int32{41, 73, 89, 7, 10, 1, 59, 58, 84, 77, 77, 97, 58, 1, 86, 58, 26, 10, 86, 51}
var thirdArr []int32 = []int32{33, 47, 70, 37, 8, 53, 13, 93, 71, 72, 51, 100, 60, 87, 97}
var fourthArr []int32 = []int32{2, 3, 4, 5, 6, 7, 8, 9}
var fiftyArr []int32 = []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func TestLeftRotation(t *testing.T) {
	t.Run("shoud left rotate array", func(t *testing.T) {
		tests := [][][]int32{
			{LeftRotation(firstArr, 4), []int32{5, 1, 2, 3, 4}},
			{LeftRotation(secondArr, 10), []int32{77, 97, 58, 1, 86, 58, 26, 10, 86, 51, 41, 73, 89, 7, 10, 1, 59, 58, 84, 77}},
			{LeftRotation(thirdArr, 13), []int32{87, 97, 33, 47, 70, 37, 8, 53, 13, 93, 71, 72, 51, 100, 60}},
			{LeftRotation(fourthArr, 20), []int32{6, 7, 8, 9, 2, 3, 4, 5}},
			{LeftRotation(fiftyArr, 13), []int32{4, 5, 6, 7, 8, 9, 10, 1, 2, 3}},
		}

		for _, v := range tests {
			result := v[0]
			expected := v[1]

			require.Equal(t, expected, result)
		}
	})
}
