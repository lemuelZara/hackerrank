package warmup

func RepeatedString(s string, n int64) int64 {
	var frequency int64 = 0
	var strLen int64 = int64(len(s))

	// if the string is only 'a'
	if strLen == 1 && s[0] == 'a' {
		return n
	}

	// taking the string 'aba' as a basis, the frequency of the letter 'a' will be 2
	for i := int64(0); i < strLen; i++ {
		if s[i] == 'a' {
			frequency++
		}
	}

	// represents the number of complete repetitions of the original string
	// ex: 'aba', n = 10, will have 3 complete repetitions -> -> aba aba aba a
	rep := n / strLen

	// represents the additional characters needed beyond the complete repetitions
	// ex: 'aba', n = 10, only the 'a' will remain -> aba aba aba a
	remainingChars := n % strLen

	// multiply the total number of complete repetitions
	frequency *= rep

	for i := 0; i < int(remainingChars); i++ {
		if s[i] == 'a' {
			frequency++
		}
	}

	return frequency
}
