package arrays

import "math"

func TwoDimensionalArray(arr [][]int32) int32 {
	var maxSum int32 = math.MinInt32

	for i := 0; i < len(arr); i++ {
		if i+2 >= len(arr) {
			break
		}

		for j := 0; j < len(arr[i]); j++ {
			if j+2 >= len(arr[i]) {
				break
			}

			firstRow := arr[i][j] + arr[i][j+1] + arr[i][j+2]
			secondRow := arr[i+1][j+1]
			thirdRow := arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]

			totalSum := firstRow + secondRow + thirdRow

			if totalSum > maxSum {
				maxSum = totalSum
			}
		}
	}

	return maxSum
}
