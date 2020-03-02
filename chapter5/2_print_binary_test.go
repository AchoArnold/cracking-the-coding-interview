package chapter5

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestPrintBinaryRepresentationOfFloat(t *testing.T) {
	inputOutput := map[float32]string {
		0.625 : "0.101",
		0.999999999: "ERROR",
		-1 : "ERROR",
		10 : "ERROR",
	}

	for input, output := range inputOutput {
		t.Run(fmt.Sprintf("Input = %f, Expected %s", input, output), func(t *testing.T) {
			assert.Equal(t, output, PrintBinaryRepresentationOfFloat(input))
		})
	}
}
