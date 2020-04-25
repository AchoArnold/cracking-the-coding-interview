package chapter16

import (
	"container/list"
	"strings"
)

var smalls = []string{
	"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen",
	"Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen",
}

var tens = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}

var bigs = []string{"", "Thousand", "Million", "Billion"}

var hundred = "Hundred"
var negative = "Negative"

// ConvertNumberToString converts an integer into it's string counterpart
func ConvertNumberToString(num int) string {
	if num == 0 {
		return smalls[0]
	}

	if num < 0 {
		return negative + " " + ConvertNumberToString(-1*num)
	}

	parts := list.New()
	chunkCount := 0

	for num > 0 {
		if num%1000 != 0 {
			chunk := convertChunk(num%1000) + " " + bigs[chunkCount]
			parts.PushFront(chunk)
		}

		num /= 1000 // shift chunk
		chunkCount++
	}

	return listToString(parts)
}

func convertChunk(number int) string {
	parts := list.New()

	// convert hundreds
	if number >= 100 {
		parts.PushBack(smalls[number/100])
		parts.PushBack(hundred)
		parts.PushBack("And")
		number %= 100
	}

	// convert tens
	if number >= 10 && number <= 19 {
		parts.PushBack(smalls[number])
	} else if number >= 20 {
		parts.PushBack(tens[number/10])
		number %= 10
	}

	// convert ones place
	if number >= 1 && number <= 9 {
		parts.PushBack(smalls[number])
	}

	return listToString(parts)
}

// converts a list to string
func listToString(list *list.List) string {
	sb := new(strings.Builder)

	for list.Len() > 1 {
		sb.WriteString(list.Front().Value.(string))
		sb.WriteString(" ")
		list.Remove(list.Front())
	}

	sb.WriteString(list.Front().Value.(string))
	list.Remove(list.Front())

	return strings.Trim(sb.String(), " ")
}
