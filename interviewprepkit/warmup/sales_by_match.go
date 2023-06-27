package warmup

func SalesByMatch(n int32, ar []int32) int32 {
	var totalPairs = 0
	var pairs = make(map[int32]int32)

	for i := 0; int32(i) < n; i++ {
		pairs[ar[int32(i)]]++

		if pairs[ar[int32(i)]] == 2 {
			totalPairs++

			pairs[ar[int32(i)]] = 0
		}
	}

	return int32(totalPairs)
}
