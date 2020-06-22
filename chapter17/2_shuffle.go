package chapter17

import "math/rand"

// ShuffleArray shuffles an array
func ShuffleArray(array []int) {
	for i := 0; i < len(array); i++ {
		k := rand.Int() % (i + 1)
		temp := array[k]
		array[k] = array[i]
		array[i] = temp
	}
}
