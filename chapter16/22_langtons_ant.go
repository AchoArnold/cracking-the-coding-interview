package chapter16

import (
	"fmt"
	"math"
	"strings"
)

type position struct {
	row int
	col int
}

type orientation string

var (
	orientationLeft  = orientation("left")
	orientationRight = orientation("right")
	orientationUp    = orientation("up")
	orientationDown  = orientation("down")
)

func (o orientation) getTurn(clockwise bool) orientation {
	if o == orientationLeft {
		if clockwise {
			return orientationUp
		}
		return orientationDown
	}

	if o == orientationUp {
		if clockwise {
			return orientationRight
		}
		return orientationLeft
	}

	if o == orientationRight {
		if clockwise {
			return orientationDown
		}
		return orientationUp
	}

	if clockwise {
		return orientationLeft
	}
	return orientationRight
}

func (o orientation) string() string {
	if o == orientationLeft {
		return "\u2190"
	}
	if o == orientationUp {
		return "\u2191"
	}

	if o == orientationRight {
		return "\u2192"
	}
	return "\u2193"
}

type ant struct {
	position    position
	orientation orientation
}

func newAnt() ant {
	return ant{
		position: position{
			row: 0,
			col: 0,
		},
		orientation: orientationRight,
	}
}

func (a *ant) turn(clockwise bool) {
	a.orientation = a.orientation.getTurn(clockwise)
}

func (a *ant) move() {
	switch a.orientation {
	case orientationLeft:
		a.position.col--
		break
	case orientationRight:
		a.position.col++
		break
	case orientationUp:
		a.position.row--
		break
	case orientationDown:
		a.position.row++
		break
	}
}

func (a *ant) adjustPosition(shiftRow int, shiftColumn int) {
	a.position.row += shiftRow
	a.position.col += shiftColumn
}

type board struct {
	whites            map[position]struct{}
	ant               ant
	topLeftCorner     position
	bottomRightCorner position
}

func newBoard() board {
	return board{
		whites:            map[position]struct{}{},
		ant:               newAnt(),
		topLeftCorner:     position{0, 0},
		bottomRightCorner: position{0, 0},
	}
}

func (b *board) move() {
	b.ant.turn(b.isWhite(b.ant.position))
	b.flipPosition(b.ant.position)
	b.ant.move()
	b.ensureFit(b.ant.position)
}

func (b *board) ensureFit(position position) {
	row := position.row
	col := position.col

	b.topLeftCorner.row = int(math.Min(float64(b.topLeftCorner.row), float64(row)))
	b.topLeftCorner.col = int(math.Min(float64(b.topLeftCorner.col), float64(col)))

	b.bottomRightCorner.row = int(math.Max(float64(b.bottomRightCorner.row), float64(row)))
	b.bottomRightCorner.col = int(math.Max(float64(b.bottomRightCorner.col), float64(col)))
}

func (b *board) String() string {
	var sb strings.Builder

	rowMin := b.topLeftCorner.row
	rowMax := b.bottomRightCorner.row
	colMin := b.topLeftCorner.row
	colMax := b.bottomRightCorner.col

	for r := rowMin; r <= rowMax; r++ {
		for c := colMin; c <= colMax; c++ {
			if r == b.ant.position.row && c == b.ant.position.col {
				sb.WriteString(b.ant.orientation.string())
			} else if (b.isWhite(position{r, c})) {
				sb.WriteString("X")
			} else {
				sb.WriteString("_")
			}
		}
		sb.WriteString("\n")
	}

	sb.WriteString(fmt.Sprintf("Ant: %s. \n", b.ant.orientation.string()))
	return sb.String()
}

func (b *board) flipPosition(position position) {
	if _, ok := b.whites[position]; ok {
		delete(b.whites, position)
	} else {
		b.whites[position] = struct{}{}
	}
}

func (b *board) isWhite(position position) bool {
	_, ok := b.whites[position]
	return ok
}

// LangtonsAnt performs a simulation of the langtons ant.
func LangtonsAnt(moves int) string {
	board := newBoard()
	for i := 0; i < moves; i++ {
		board.move()
	}

	return board.String()
}
