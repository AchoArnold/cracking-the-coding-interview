package main

func  isPermutation(string1 string, string2 string) bool  {
	if len(string1) != len(string2) {
		return false
	}

	string2AsRune :=  map[int]int32{}

	for index, character := range string2  {
		string2AsRune[index] = character
	}

	for _,character := range string1 {
		characterFound := false
		for index, string2Character := range string2AsRune {
			if character == string2Character {
				delete(string2AsRune, index)
				characterFound = true
				break
			}
		}

		if characterFound == false {
			return false
		}
	}


	return len(string2AsRune) == 0
}