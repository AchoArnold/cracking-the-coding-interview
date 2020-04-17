package chapter16

// SwapNumbersInPlace swaps 2 numbers without using any extra memory
func SwapNumbersInPlace(a, b *int) {
	*b = *b + *a
	*a = *b - *a
	*b = *b - *a
}

// SwapNumbersInPlaceUsingBitOperations swaps 2 numbers using XOR operations
func SwapNumbersInPlaceUsingBitOperations(a, b *int) {
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}
