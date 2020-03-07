package chapter5

func PrintBinaryRepresentationOfFloat(num float32) (result string) {
	if num >= 1 || num <= 0 {
		return "ERROR"
	}

	result += "0."
	for num > 0 {
		if len(result) > 32 {
			return "ERROR"
		}

		r := num * 2
		if r >= 1 {
			result += "1"
			num = r - 1
		} else {
			result += "0"
			num = r
		}
	}

	return result
}
