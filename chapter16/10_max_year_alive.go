package chapter16

import "sort"

type person struct {
	birth, death int
}

// MaxAliveYear gets the max year in which people are alive withing a certain range.
func MaxAliveYear(people []person, min, _ int) int {
	births := getSortedYears(people, true)
	deaths := getSortedYears(people, false)

	birthIndex := 0
	deathIndex := 0

	currentlyAlive := 0
	maxAlive := 0
	maxAliveYear := min

	for birthIndex < len(births) {
		if births[birthIndex] <= deaths[deathIndex] {
			currentlyAlive++ //include birth
			if currentlyAlive > maxAlive {
				maxAlive = currentlyAlive
				maxAliveYear = births[birthIndex]
			}

			birthIndex++ //move birth index
		} else if births[birthIndex] > deaths[deathIndex] {
			currentlyAlive-- //include death
			deathIndex++     // move death index
		}
	}

	return maxAliveYear
}

func getSortedYears(people []person, copyBirthYear bool) []int {
	years := make([]int, len(people))

	for i := 0; i < len(people); i++ {
		if copyBirthYear {
			years[i] = people[i].birth
		} else {
			years[i] = people[i].death
		}
	}

	sort.Ints(years)
	return years
}
