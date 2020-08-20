package chapter17

// SearchAllSubstringsInString returns all substring in a string
func SearchAllSubstringsInString(big string, smalls []string) map[string][]int {
	lookup := map[string][]int{}

	for _, small := range smalls {
		lookup[small] = searchLocations(big, small)
	}

	return lookup
}

func searchLocations(big string, small string) []int {
	locations := []int{}

	for i := 0; i < len(big)-len(small)+1; i++ {
		if isSubstringAtLocation(big, small, i) {
			locations = append(locations, i)
		}
	}

	return locations
}

func isSubstringAtLocation(big, small string, index int) bool {
	for i := 0; i < len(small); i++ {
		if big[i+index] != small[i] {
			return false
		}
	}

	return true
}
