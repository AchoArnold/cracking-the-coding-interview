package chapter8

type Point struct {
	i int
	j int
}

func GetPath(maze [][]int) (path []Point) {
	if maze == nil || len(maze[0]) == 0 {
		return path
	}

	var pointsToVisit []Point
	var visitedPoints []Point

	maxRow := len(maze) - 1
	maxCol := len(maze[0]) - 1

	start := Point{0, 0}
	pointsToVisit = append(pointsToVisit, start)
	visitedPoints = append(visitedPoints, start)
	paths := [][]Point{{start}}

	for len(pointsToVisit) != 0 {
		var nextPointsToVisit []Point

		for _, point := range pointsToVisit {
			var localPaths [][]Point
			if point.i+1 <= maxRow && NotIn(visitedPoints, Point{point.i + 1, point.j}) {
				nextPointsToVisit = append(nextPointsToVisit, Point{point.i + 1, point.j})
				localPaths = append(localPaths, append(pathEndingWithPoint(paths, point), Point{point.i + 1, point.j}))
			}

			if point.j+1 <= maxCol && NotIn(visitedPoints, Point{point.i, point.j + 1}) {
				nextPointsToVisit = append(nextPointsToVisit, Point{point.i, point.j + 1})
				localPaths = append(localPaths, append(pathEndingWithPoint(paths, point), Point{point.i, point.j + 1}))
			}

			if point.i == maxRow && point.j == maxCol {
				return pathEndingWithPoint(paths, point)
			}

			var newPath [][]Point
			for _, path := range paths {
				val := path[len(path)-1]
				if !(val.i == point.i && val.j == point.j) {
					newPath = append(newPath, path)
				}
			}

			paths = newPath
			for _, path := range localPaths {
				paths = append(paths, path)
			}

			visitedPoints = append(visitedPoints, point)
		}

		pointsToVisit = nextPointsToVisit
	}

	return path
}

func pathEndingWithPoint(paths [][]Point, point Point) (result []Point) {
	for _, path := range paths {
		val := path[len(path)-1]
		if val.i == point.i && val.j == point.j {
			return path
		}
	}

	return result
}
func NotIn(haystack []Point, needle Point) bool {
	for _, val := range haystack {
		if val.i == needle.i && val.j == needle.j {
			return false
		}
	}

	return true
}
