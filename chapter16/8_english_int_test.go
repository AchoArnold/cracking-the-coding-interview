package chapter16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertNumberToString(t *testing.T) {
	testData := []struct {
		num      int
		expected string
	}{
		{1, "One"},
		{-0, "Zero"},
		{0, "Zero"},
		{1234253, "One Million Two Hundred And Thirty Four Thousand Two Hundred And Fifty Three"},
		{121, "One Hundred And Twenty One"},
	}

	for _, data := range testData {
		t.Run(data.expected, func(t *testing.T) {
			assert.Equal(t, data.expected, ConvertNumberToString(data.num))
		})
	}
}
