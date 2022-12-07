package fibonacci_test

import (
	"testing"

	"github.com/LeonardoBozCaitano/golang-algorithms/pkg/fibonacci"
	"gotest.tools/v3/assert"
)

func Test_fibonacci(t *testing.T) {

	tests := []struct {
		name  string
		index int
		want  int
	}{
		{
			name:  "Should calculate index 6",
			index: 6,
			want:  8,
		},
		{
			name:  "Should calculate index 7",
			index: 7,
			want:  13,
		},
		{
			name:  "Should calculate index 8",
			index: 8,
			want:  21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := fibonacci.Calculate(tt.index)
			assert.Equal(t, tt.want, output)
		})
	}
}

func Test_fibonacciWithCache(t *testing.T) {

	tests := []struct {
		name  string
		index int
		want  int
	}{
		{
			name:  "Should calculate index 6",
			index: 6,
			want:  8,
		},
		{
			name:  "Should calculate index 7",
			index: 7,
			want:  13,
		},
		{
			name:  "Should calculate index 8",
			index: 8,
			want:  21,
		},
		{
			name:  "Should calculate index 50",
			index: 50,
			want:  12586269025,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := fibonacci.CalculateWithCache(tt.index, make(map[int]int))
			assert.Equal(t, tt.want, output)
		})
	}
}
