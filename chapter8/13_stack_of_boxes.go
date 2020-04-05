package chapter8

import (
	"math"
	"sort"
)

// Box is a box with height, width and depth dimensions
type Box struct {
	height int
	width  int
	depth  int
}

func (a Box) canBeAbove(b Box) bool {
	if (b == Box{}) {
		return true
	}
	if a.width < b.width && a.height < b.height && a.depth < b.depth {
		return true
	}

	return false
}

// MaxStackHeight calculates the maximum height that can be archived by stacking an array of boxes
func MaxStackHeight(boxes []Box) int {
	sort.Slice(boxes, func(i, j int) bool {
		return boxes[i].height > boxes[j].height
	})

	maxHeight := 0
	stackMap := make([]int, len(boxes))

	for i := 0; i < len(boxes); i++ {
		height := createStack(boxes, i, stackMap)
		maxHeight = int(math.Max(float64(maxHeight), float64(height)))
	}

	return maxHeight
}

func createStack(boxes []Box, bottomIndex int, stackMap []int) int {
	if bottomIndex < len(boxes) && stackMap[bottomIndex] > 0 {
		return stackMap[bottomIndex]
	}

	bottom := boxes[bottomIndex]
	maxHeight := 0

	for i := bottomIndex + 1; i < len(boxes); i++ {
		if boxes[i].canBeAbove(bottom) {
			height := createStack(boxes, i, stackMap)
			maxHeight = int(math.Max(float64(maxHeight), float64(height)))
		}
	}

	maxHeight += bottom.height
	stackMap[bottomIndex] = maxHeight
	return maxHeight
}
