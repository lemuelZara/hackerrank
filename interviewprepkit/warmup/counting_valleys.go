package warmup

import "strings"

const (
	UPHILL    = "U"
	DOWNHILL  = "D"
	SEA_LEVEL = 0
)

func CountingValleys(steps int32, path string) int32 {
	var altitude, valleys int32 = 0, 0

	for _, step := range strings.Split(path, "") {
		if step == UPHILL {
			altitude += 1

			if altitude == SEA_LEVEL {
				valleys++
				continue
			}
		}

		if step == DOWNHILL {
			altitude -= 1
		}
	}

	return valleys
}
