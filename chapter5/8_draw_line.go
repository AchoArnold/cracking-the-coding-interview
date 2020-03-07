package chapter5

func DrawLine(screen []byte, width, x1, x2, y int) {
	startOffset := x1 % 8
	firstFullByte := x1 / 8

	if startOffset != 0 {
		firstFullByte++
	}

	endOffset := x2 % 8
	lastFullByte := x2 / 8
	if endOffset != 7 {
		lastFullByte--
	}

	// Set full bytes
	for b := firstFullByte; b <= lastFullByte; b++ {
		screen[(width/8)*y+b] = byte(0xFF)
	}

	// Create masks for start and end of line
	startMask := byte(0xFF >> startOffset)
	endMask := byte(0xFF>>endOffset + 1)

	// Set start and end of line
	if x1/8 == x2/8 {
		mask := startMask & endMask
		screen[(width/8)*y+(x1/8)] |= mask
	} else {
		if startOffset != 0 {
			byteNumber := (width/8)*y + firstFullByte - 1
			screen[byteNumber] |= startMask
		}

		if endOffset != 7 {
			byteNumber := (width/8)*y + lastFullByte + 1
			screen[byteNumber] |= endMask
		}
	}
}
