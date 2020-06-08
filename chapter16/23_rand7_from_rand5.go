package chapter16

import "math/rand"

func rand5() int {
	return rand.Intn(5)
}

// Rand7 generates a random number from 0-7
func Rand7() int {
	for {
		num := 5*rand5() + rand5()
		if num < 21 {
			return num % 7
		}
	}
}
