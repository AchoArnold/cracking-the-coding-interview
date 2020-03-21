package chapter8

func TripleStep(n int) (result int) {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	if n == 3 {
		return 4
	}

	prev1 := 4
	prev2 := 2
	prev3 := 1

	for i := 4; i <= n; i++ {
		cur := prev1
		prev1 = prev1 + prev2 + prev3
		prev3 = prev2
		prev2 = cur
	}

	return prev1
}

func TripleStepRecursive(n int) (result int) {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	if n == 3 {
		return 4
	}

	return TripleStepRecursive(n-1) + TripleStepRecursive(n-2) + TripleStepRecursive(n-3)
}
