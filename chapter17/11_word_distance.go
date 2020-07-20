package chapter17

import "math"

// LocationPair represents 2 locations
type LocationPair struct {
	location1 int
	location2 int
}

func (lp *LocationPair) setLocations(first, second int) {
	lp.location1 = first
	lp.location2 = second
}

func (lp *LocationPair) setLocationsWithPair(pair LocationPair) {
	lp.setLocations(pair.location1, pair.location2)
}

func (lp LocationPair) distance() int {
	return int(math.Abs(float64(lp.location1 - lp.location2)))
}

func (lp LocationPair) isValid() bool {
	return lp.location1 >= 0 && lp.location2 >= 0
}

func (lp *LocationPair) updateWithMin(locPair LocationPair) {
	if !lp.isValid() || locPair.distance() < lp.distance() {
		lp.setLocationsWithPair(locPair)
	}
}

// FindClosest finds the closest distance between 2 words in a string
func FindClosest(word1, word2 string, locations map[string][]int) *LocationPair {
	locations1 := locations[word1]
	locations2 := locations[word2]
	return findMinDistancePair(locations1, locations2)
}

func findMinDistancePair(array1, array2 []int) *LocationPair {
	if len(array1) == 0 || len(array2) == 0 {
		return nil
	}

	index1 := 0
	index2 := 0

	best := &LocationPair{array1[0], array2[0]}
	current := LocationPair{array1[0], array2[0]}

	for index1 < len(array1) && index2 < len(array2) {
		current.setLocations(array1[index1], array2[index2])
		best.updateWithMin(current) // if shorter update values
		if current.location1 < current.location2 {
			index1++
		} else {
			index2++
		}
	}

	return best
}

// pre-computation
func getWordLocations(words []string) map[string][]int {
	locations := map[string][]int{}
	for i, word := range words {
		locations[word] = append(locations[word], i)
	}
	return locations
}
