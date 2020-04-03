package chapter8

// Color represents a color
type Color string

// PaintFill paints radially outwards
func PaintFill(screen [][]Color, r, c int, nColor Color) bool {
	if len(screen[0]) == 0 || screen[r][c] == nColor {
		return false
	}

	return paintFill(screen, r, c, screen[r][c], nColor)
}

func paintFill(screen [][]Color, r, c int, ocolor Color, ncolor Color) bool {
	if r < 0 || r >= len(screen) || c < 0 || c >= len(screen) {
		return false
	}

	if screen[r][c] == ocolor {
		screen[r][c] = ncolor
		paintFill(screen, r-1, c, ocolor, ncolor) // up
		paintFill(screen, r+1, c, ocolor, ncolor) // down
		paintFill(screen, r, c-1, ocolor, ncolor) // left
		paintFill(screen, r, c+1, ocolor, ncolor) // right
	}

	return true
}
