package chapter16

// point represents a point in the cartesian plane
type point struct {
	x float64
	y float64
}

// newPoint creates a new point instance
func newPoint(x, y float64) point {
	return point{x, y}
}

func (p *point) setLocation(x, y float64) {
	p.x = x
	p.y = y
}

type line struct {
	slope      float64
	yIntercept float64
}

func newLine(start, end point) line {
	deltaY := end.y - start.y
	deltaX := end.x - start.x
	slope := deltaY / deltaX

	return line{
		slope:      deltaY / deltaX,
		yIntercept: end.y - slope*end.x,
	}
}

// getIntersectionPoint returns the intersection point of 2 lines
func getIntersectionPoint(start1, end1, start2, end2 point) *point {
	// rearranging these si that, in order of x values: start is before end an point 1 is before point 2
	// this wil make some of the later logic simpler

	if start1.x > end1.x {
		swapPoint(start1, end1)
	}

	if start2.x > end2.x {
		swapPoint(start2, end2)
	}

	if start1.x > start2.x {
		swapPoint(start1, start2)
		swapPoint(end1, end2)
	}

	// compute lines including slope and y-intercept
	line1 := newLine(start1, end1)
	line2 := newLine(start2, end2)

	// if these lines are parallel, they intercept only if they have the same y intercept and start 2 is on line 1
	if line1.slope == line2.slope {
		if line1.yIntercept == line2.yIntercept && pointIsBetween(start1, start2, end1) {
			result := start2
			return &result
		}

		return nil
	}

	// get intersection coordinate
	x := (line2.yIntercept - line1.yIntercept) / (line1.slope - line2.slope)
	y := x*line1.slope + line1.yIntercept

	intersection := newPoint(x, y)

	// check if within the line segment range
	if pointIsBetween(start1, intersection, end1) && pointIsBetween(start2, intersection, end2) {
		return &intersection
	}

	return nil
}

// checks if middle is between start and end
func floatIsBetween(start, middle, end float64) bool {
	if start > end {
		return end <= middle && middle <= start
	}

	return start <= middle && middle <= end
}

func pointIsBetween(start, middle, end point) bool {
	return floatIsBetween(start.x, middle.x, end.x) && floatIsBetween(start.y, middle.y, end.y)
}

func swapPoint(one, two point) {
	x := one.x
	y := one.y
	one.setLocation(two.x, two.y)
	two.setLocation(x, y)
}
