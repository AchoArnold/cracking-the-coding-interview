package chapter17

// Range represents a range of numbers from start till small
type Range struct {
	start int
	end   int
}

func newRange(start, end int) Range {
	return Range{start, end}
}

func (r Range) length() int {
	return r.end - r.start
}

func (r Range) shorterThan(other Range) bool {
	return r.length() < other.length()
}

// ShortestSuperSequence returns the shortest sequence of all the elements in `small` in the `big` array
func ShortestSuperSequence(big, small []int) Range {
	closures := getClosures(big, small)
	return getShortestClosure(closures)
}

func getClosures(big, small []int) []int {
	closure := make([]int, len(big))

	for i := 0; i < len(small); i++ {
		sweepForClosure(big, closure, small[i])
	}

	return closure
}

// Do backwards sweep and update the closure list with the next occurrence of value, if it's later than the current closure
func sweepForClosure(big, closures []int, value int) {
	next := -1

	for i := len(big) - 1; i >= 0; i-- {
		if big[i] == value {
			next = i
		}

		if (next == -1 || closures[i] < next) && closures[i] != -1 {
			closures[i] = next
		}
	}
}

// Get shortest closure
func getShortestClosure(closures []int) Range {
	shortest := newRange(0, closures[0])
	for i := 1; i < len(closures); i++ {
		if closures[i] == -1 {
			break
		}

		r := newRange(i, closures[i])
		if !shortest.shorterThan(r) {
			shortest = r
		}
	}

	return shortest
}
