package main

func URLify(inputString string, length int) string {
	search := " "
	replace := "%20"
	output := ""

	for index,character := range inputString{
		if index == length {
			break
		}
		if string(character) == search {
			output += replace
		} else {
			output += string(character)
		}
	}

	return output
}