package chapter17

import (
	"math"
)

// GetKthMagicNumber finds the kth number such that the only prime factors are 3, 5, 7
func GetKthMagicNumber(k int) int {
	if k < 0 {
		return 0
	}

	val := 0

	var queue3 []int
	var queue5 []int
	var queue7 []int
	queue3 = append(queue3, 1)

	// Include 0th through kth iteration
	for i := 0; i < k; i++ {
		v3 := math.MaxInt32
		if len(queue3) > 0 {
			v3 = queue3[len(queue3)-1]
		}

		v5 := math.MaxInt32
		if len(queue5) > 0 {
			v5 = queue5[len(queue5)-1]
		}

		v7 := math.MaxInt32
		if len(queue7) > 0 {
			v7 = queue7[len(queue7)-1]
		}

		val = int(math.Min(float64(v3), math.Min(float64(v5), float64(v7))))

		if val == v3 { // enqueue into queue 3, 5, 7
			queue3 = queue3[1:]            // queue.remove()
			queue3 = append(queue3, 3*val) // append
			queue5 = append(queue5, 5*val) //append

		} else if val == v5 {
			queue5 = queue5[1:]
			queue5 = append(queue5, 5*val)
		} else if val == v7 {
			queue7 = queue7[1:] // enqueue into q7
		}

		queue7 = append(queue7, 7*val) // always enqueue into q7
	}

	return val
}
