package chapter5

func SwapOddEvenBits(num int) (result int) {
	return ((num & 0xaaaaaaaa) >> 1) | ((num & 0x55555555) << 1)
}
