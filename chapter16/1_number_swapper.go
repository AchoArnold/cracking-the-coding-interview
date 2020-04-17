package chapter16

// SwapNumbersInPlace swaps 2 numbers without using any extra memory
func SwapNumbersInPlace(a, b *int) {
	*b = *b + *a
	*a = *b - *a
	*b = *b - *a
}
