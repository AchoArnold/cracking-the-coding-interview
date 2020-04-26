package chapter16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDivingBoardLengths(t *testing.T) {
	testData := []struct {
		name                    string
		number, shorter, longer int
		result                  []int
	}{
		{"different", 3, 1, 2, []int{4, 3, 6, 5}},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.ElementsMatch(t, data.result, GetDivingBoardLengths(data.number, data.shorter, data.longer))
		})
	}
}

func TestGetDivingBoardLoop(t *testing.T) {
	testData := []struct {
		name                    string
		number, shorter, longer int
		result                  []int
	}{
		{"different", 3, 1, 2, []int{4, 3, 6, 5}},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.ElementsMatch(t, data.result, GetDivingBoardLoop(data.number, data.shorter, data.longer))
		})
	}
}
