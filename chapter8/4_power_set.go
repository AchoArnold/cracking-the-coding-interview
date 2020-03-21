package chapter8

func GetPowerSet(set []int) [][]int {
	if len(set) == 0 {
		return [][]int{}
	}

	prev := [][]int{{}, {set[0]}}
	for i := 1; i < len(set); i++ {
		prev = addItemToSet(prev, set[i])
	}

	return prev
}

func addItemToSet(set [][]int, item int) [][]int {
	var newSet [][]int
	for _, val := range set {
		newSet = append(newSet, append(val, item))
	}

	return append(newSet, set...)
}
