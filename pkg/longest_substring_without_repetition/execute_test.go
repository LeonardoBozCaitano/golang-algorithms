package longest_substring_without_repetition_test

import (
	"testing"

	"github.com/LeonardoBozCaitano/golang-algorithms/pkg/longest_substring_without_repetition"
	"gotest.tools/v3/assert"
)

func Test_longest_substring_without_repetition(t *testing.T) {

	tests := []struct {
		name string
		word string
		want int
	}{
		{
			name: "Should return 1",
			word: "bbbbbb",
			want: 1,
		},
		{
			name: "Should return 2",
			word: "bbbbbaa",
			want: 2,
		},
		{
			name: "Should return 10",
			word: "aaaaabcqwertyuuuuu",
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := longest_substring_without_repetition.LongestSubstringWithoutChatRepetition(tt.word)
			assert.Equal(t, tt.want, output)
		})
	}
}
