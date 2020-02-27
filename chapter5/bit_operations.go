package chapter5

func getBit(num int, i int) int {
	if num&(1<<i) == 0 {
		return 0
	}

	return 1
}

func setBit(num int, i int) int {
	return num | (1 << i)
}

func clearBit(num int, i int) int {
	mask := ^(1 << i)
	return num & mask
}

// i inclusive
func clearBitFromMostSignificantBitThroughI(num int, i int) int {
	mask := (1 << i) - 1
	return num & mask
}

func clearBitsIThrough0(num int, i int) int {
	mask := -1 << (i + 1)
	return num & mask
}

func updateBit(num int, i int, bit bool) int {
	value := 0
	if bit == true {
		value = 1
	}

	mask := ^(1 << i)

	return num&mask | value<<1
}
