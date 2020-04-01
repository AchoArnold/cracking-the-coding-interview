package chapter8

func RecursiveMultiply(n, m int) int {
	bigger := n
	smaller := m
	if m > n {
		smaller = n
		bigger = m
	}

	return recursiveMultiply(smaller, bigger)
}

func recursiveMultiply(n int, m int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return m
	}

	mid := n / 2
	halProd := recursiveMultiply(mid, m)

	if n%2 == 0 {
		return halProd + halProd
	}

	return halProd + halProd + m
}
