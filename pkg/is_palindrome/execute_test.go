package is_palindrome_test

import (
	"testing"

	"github.com/LeonardoBozCaitano/golang-algorithms/pkg/is_palindrome"
	"gotest.tools/v3/assert"
)

func Test_fibonacci(t *testing.T) {

	tests := []struct {
		name   string
		number int
		want   bool
	}{
		{
			name:   "Should return true",
			number: 121,
			want:   true,
		},
		{
			name:   "Should return false",
			number: -121,
			want:   false,
		},
		{
			name:   "Should return true",
			number: 45554,
			want:   true,
		},
		{
			name:   "Should return false",
			number: 10,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := is_palindrome.IsPalindrome(tt.number)
			assert.Equal(t, tt.want, output)
		})
	}
}
