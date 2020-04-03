package chapter8

// GenerateAllValidParens generates all valid parenthesis having count.
func GenerateAllValidParens(count int) []string {
	str := make([]int32, count*2)
	return addParenthesis([]string{}, count, count, str, 0)
}

func addParenthesis(list []string, leftRem, rightRem int, str []int32, index int) []string {
	// Invalid Sate
	if leftRem < 0 || rightRem < leftRem {
		return list
	}

	if leftRem == 0 && rightRem == 0 { // Out of left and right parenthesis
		list = append(list, string(str))
	} else {
		str[index] = '(' // Add left and recurse
		list = addParenthesis(list, leftRem-1, rightRem, str, index+1)

		str[index] = ')' // add right and recurse
		list = addParenthesis(list, leftRem, rightRem-1, str, index+1)
	}

	return list
}
