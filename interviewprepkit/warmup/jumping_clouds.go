package warmup

const SAFE_CLOUD = 0

func JumpingOnClouds(c []int32) int32 {
	var jumps int32 = 0

	for i := 0; i < len(c)-1; {
		// checks if the next cloud after the two-step jump is safe
		if i+2 < len(c) && c[i+2] == SAFE_CLOUD {
			i += 2
		} else {
			i += 1
		}
		// the minimum number of jumps needed to reach the last cloud
		jumps++
	}

	return jumps
}
