package warmup

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepeatedString(t *testing.T) {
	t.Run("should return the frequency of 'a' in the substring", func(t *testing.T) {
		tests := [][]int64{
			{RepeatedString("a", 1000000000), 1000000000},
			{RepeatedString("aba", 10), 7},
			{RepeatedString("d", 590826798023), 0},
			{RepeatedString("beeaabc", 711560125001), 203302892858},
			{RepeatedString("kmretasscityylpdhuwjirnqimlkcgxubxmsxpypgzxtenweirknjtasxtvxemtwxuarabssvqdnktqadhyktagjxoanknhgilnm", 736778906400), 51574523448},
		}

		for _, v := range tests {
			result := v[0]
			expected := v[1]

			require.Equal(t, expected, result)
		}
	})
}
