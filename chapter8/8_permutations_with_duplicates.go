package chapter8

// GetPermutationsWithDuplicates returns all permutations of a string without duplicates.
func GetPermutationsWithDuplicates(s string) []string {
	frequencyTable := buildFrequencyTable(s)
	return getPermutations(frequencyTable, "", len(s), []string{})
}

func getPermutations(freqTable map[string]int, prefix string, remaining int, result []string) (permutations []string) {
	// Best case, permutation has been completed
	if remaining == 0 {
		return append(result, prefix)
	}

	// Try remaining letters for next char, and generate remaining permutations
	for char := range freqTable {
		count := freqTable[char]

		if count > 0 {
			freqTable[char]--
			result = getPermutations(freqTable, prefix+char, remaining-1, result)
			freqTable[char] = count
		}
	}

	return result
}

func buildFrequencyTable(s string) (frequencyTable map[string]int) {
	frequencyTable = map[string]int{}

	for _, char := range s {
		if _, ok := frequencyTable[string(char)]; !ok {
			frequencyTable[string(char)] = 0
		}

		frequencyTable[string(char)]++
	}

	return frequencyTable
}
