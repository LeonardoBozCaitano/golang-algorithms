package is_palindrome

func IsPalindrome(number int) bool {
	if number < 0 || (number%10 == 0 && number != 0) {
		return false
	}

	revertedNumber := 0
	for number > revertedNumber {
		revertedNumber = revertedNumber*10 + number%10
		number /= 10
	}
	return true
}
