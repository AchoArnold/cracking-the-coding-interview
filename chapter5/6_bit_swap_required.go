package chapter5

func BitSwapRequired(num1, num2 int) (count int) {
	for c := num1 ^ num2; c != 0; c = c >> 1 {
		count += c & 1
	}

	return count
}

func BitSwapRequiredLastSignificantBit(num1, num2 int) (count int) {
	for c := num1 ^ num2; c != 0; c = c & (c - 1) {
		count++
	}

	return count
}
