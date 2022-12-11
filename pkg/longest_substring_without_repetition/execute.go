package longest_substring_without_repetition

import (
	"math"
	"strings"
)

func LongestSubstringWithoutChatRepetition(word string) int {
	wordLenght := len(word)
	biggestWordLength := 0
	slice := word[0:0]
	for i, j := 0, 0; j < wordLenght; j++ {
		if index := strings.IndexByte(string(slice), word[j]); index != -1 {
			i += index + 1
		}
		slice = word[i : j+1]
		biggestWordLength = int(math.Max(float64(biggestWordLength), float64(j-i+1)))
	}
	return biggestWordLength
}
