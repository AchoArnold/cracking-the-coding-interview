package chapter17

import "sort"

// HeightWidth is a struct having height and width
type HeightWidth struct {
	height int
	weight int
}

// returns true if main should be lined up before that.
// note that it's possible that this.isBefore(that) and that.isBefore(this) are both false
func (main HeightWidth) isBefore(that HeightWidth) bool {
	if main.height < that.height && main.weight < that.weight {
		return true
	}
	return false
}

func canAppend(solution []HeightWidth, value HeightWidth) bool {
	last := solution[len(solution)-1]
	return last.isBefore(value)
}

// LongestIncreasingSequence returns the longest way in which items can be stacked on top of each other.
func LongestIncreasingSequence(items []HeightWidth) []HeightWidth {
	sort.Slice(items, func(i, j int) bool {
		first := items[i]
		second := items[j]
		if first.height != second.height {
			return first.height < second.height
		}
		return first.weight < second.weight
	})

	var solutions [][]HeightWidth
	var bestSequence []HeightWidth

	// Find the longest sequence that terminates with each element
	// Track the longest overall subsequence as we go
	for i := 0; i < len(items); i++ {
		longestIndex := bestSequenceAtIndex(items, solutions, i)
		solutions = append(solutions, longestIndex)
		bestSequence = maxSequence(bestSequence, longestIndex)
	}

	return bestSequence
}

func maxSequence(seq1 []HeightWidth, seq2 []HeightWidth) []HeightWidth {
	if seq1 == nil {
		return seq2
	}

	if seq2 == nil {
		return seq1
	}

	if len(seq1) > len(seq2) {
		return seq1
	}

	return seq2
}

// find the longest sequence which terminates with this element
func bestSequenceAtIndex(array []HeightWidth, solutions [][]HeightWidth, index int) []HeightWidth {
	value := array[index]
	var bestSequence []HeightWidth

	// find the longest subsequence that we can append this element to
	for i := 0; i < index; i++ {
		solution := solutions[i]
		if canAppend(solution, value) {
			bestSequence = maxSequence(solution, bestSequence)
		}
	}

	best := append([]HeightWidth{}, bestSequence...)

	return append(best, value)
}
