package chapter8

// CountEval the number of parenthesis required to evaluate a string.
// TODO: Fix algorithm
func CountEval(input string, result bool) int {
	return countEval(input, result, map[string]int{})
}

func countEval(input string, result bool, cache map[string]int) int {
	if len(input) == 0 {
		return 0
	}

	if len(input) == 1 {
		if stringToBool(input) == result {
			return 1
		}
		return 0
	}

	if _, ok := cache[boolToString(result)+input]; ok {
		return cache[boolToString(result)+input]
	}

	ways := 0

	for i := 1; i < len(input); i += 2 {
		char := input[i]
		left := input[0:i]
		right := input[i+1:]

		leftTrue := countEval(left, true, cache)
		leftFalse := countEval(left, false, cache)
		rightTrue := countEval(right, true, cache)
		rightFalse := countEval(right, false, cache)

		total := (leftTrue + leftFalse) * (rightTrue + rightFalse)

		totalTrue := 0
		if char == '^' {
			totalTrue = leftTrue*rightFalse + leftFalse*rightTrue
		} else if char == '&' {
			totalTrue = leftTrue * rightTrue
		} else if char == '|' {
			totalTrue = leftTrue*rightTrue + leftFalse*rightTrue + leftTrue*rightFalse
		}

		subWays := total - totalTrue
		if result {
			subWays = totalTrue
		}

		ways += subWays
	}

	cache[boolToString(result)+input] = ways
	return ways
}

func boolToString(input bool) string {
	if input {
		return "true"
	}
	return "false"
}

func stringToBool(c string) bool {
	if c == "true" {
		return true
	}
	return false
}
