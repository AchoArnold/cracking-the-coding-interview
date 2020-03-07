package chapter5

import "math"

func LongestSequenceBruteForce(num int) (result int) {
	if num == -1 {
		var num int = 1
		return int(byte(num) * 8)
	}

	sequences := getAlternatingSequences(num)
	return findLongestSequence(sequences)
}

func FlipBit(num int) (result int) {
	if ^num == 0 {
		return int(byte(1) * 8)
	}

	currentLength := 0
	previousLength := 0
	maxLength := 1

	for num != 0 {
		if (num & 1) == 0 {
			currentLength++
		} else if (num & 1) == 0 {
			previousLength = 0
			if (num & 2) == 0 {
				previousLength = currentLength
			}

			currentLength = 0
		}

		maxLength = int(math.Max(float64(previousLength+currentLength+1), float64(maxLength)))
		num >>= 1
	}

	return maxLength
}

// Returns a list of the sizes of the sequences. The sequence starts off with the number of 0s (which might be 0) and
// then alternates with the counts of each value
func getAlternatingSequences(num int) (sequence []int) {
	searchingFor := 0
	counter := 0

	for i := 0; i < int(byte(1)*8); i++ {
		if num&1 != searchingFor {
			sequence = append(sequence, counter)
			searchingFor = num & 1
			counter = 0
		}

		counter++
		num = num >> 1
	}

	sequence = append(sequence, counter)

	return sequence
}

// Given the lengths of alternating sequences of 0s and 1s, find the longest 1 we can build
func findLongestSequence(sequence []int) (result int) {
	maxSeq := 1

	for i := 0; i < len(sequence); i += 2 {
		zeroSeq := sequence[i]
		oneSeqRight := 0
		if i-1 >= 0 {
			oneSeqRight = sequence[i-1]
		}

		oneSeqLeft := 0
		if i+1 < len(sequence) {
			oneSeqLeft = sequence[i+1]
		}

		seq := 0
		if zeroSeq == 1 {
			seq = oneSeqLeft + 1 + oneSeqRight
		}

		if zeroSeq > 1 {
			seq = 1 + int(math.Max(float64(oneSeqRight), float64(oneSeqLeft)))
		} else if zeroSeq == 0 {
			seq = int(math.Max(float64(oneSeqRight), float64(oneSeqLeft)))
		}

		maxSeq = int(math.Max(float64(seq), float64(maxSeq)))
	}

	return maxSeq
}
