package chapter8

// MakeChange returns the amount of combinations of quaters, nikels, dimes and ones needed to make n
func MakeChange(n int) (result int) {
	denoms := []int{25, 10, 5, 1}

	cache := make([][]int, n+1)
	for i := range cache {
		cache[i] = make([]int, len(denoms))
	}

	return makeChange(n, denoms, 0, cache)
}

func makeChange(amount int, denoms []int, index int, cache [][]int) int {
	// get value
	if cache[amount][index] > 0 {
		return cache[amount][index]
	}

	// one denom remaining
	if index >= len(denoms)-1 {
		return 1
	}

	denomAmount := denoms[index]
	ways := 0

	for i := 0; i*denomAmount <= amount; i++ {
		amountRemaining := amount - i*denomAmount
		ways += makeChange(amountRemaining, denoms, index+1, cache)
	}

	cache[amount][index] = ways

	return ways
}
