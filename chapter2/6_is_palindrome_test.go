package chapter2

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	testData := []struct {
		inputSlice     []int
		expectedResult bool
	}{
		{[]int{1, 2, 2}, false},
		{[]int{1, 2, 1}, true},
		{[]int{1, 2, 2, 1}, true},
	}

	for _, data := range testData {
		list := makeFromSlice(data.inputSlice)
		t.Run("List Length "+strconv.Itoa(list.Count()), func(t *testing.T) {
			assert.Equal(t, data.expectedResult, list.IsPalindrome())
		})
	}
}
