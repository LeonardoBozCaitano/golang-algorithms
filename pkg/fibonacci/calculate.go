package fibonacci

/**
 * The function should return the n-th number of the Fibonacci sequence.
 *
 * The 1st and 2nd number of the sequence is 1.
 * To generate the next number of the sequence, we sum the previous two.
 */

func Calculate(index int) int {
	if index < 3 {
		return 1
	}
	return Calculate(index-1) + Calculate(index-2)
}

func CalculateWithCache(index int, cache map[int]int) int {
	value, found := cache[index]
	if found {
		return value
	}
	if index <= 2 {
		return 1
	}
	cache[index] = CalculateWithCache(index-1, cache) + CalculateWithCache(index-2, cache)
	return cache[index]
}
