package chapter16

// MatrixPoint represents a point in a matrix
type MatrixPoint struct {
	i, j int
}

// PondSizesOptimized finds all ponds with a specified length
func PondSizesOptimized(land [][]int) (sizes []int) {
	for r := 0; r < len(land); r++ {
		for c := 0; c < len(land[r]); c++ {
			if land[r][c] == 0 {
				sizes = append(sizes, computeSize(land, r, c))
			}
		}
	}

	return sizes
}

func computeSize(land [][]int, row, col int) int {
	// if out of bounds or already visited
	if row < 0 || col < 0 || row >= len(land) || col >= len(land[row]) || land[row][col] != 0 {
		return 0
	}

	size := 1
	land[row][col] = -1 // mark as visited

	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			size += computeSize(land, row+dr, col+dc)
		}
	}

	return size
}

// PondSizesBruteForce finds the length of all  ponds in a matrix
func PondSizesBruteForce(plotOfLand [][]int) (sizes []int) {
	cache := map[MatrixPoint]bool{}

	for i := 0; i < len(plotOfLand); i++ {
		for j := 0; j < len(plotOfLand[i]); j++ {
			if plotOfLand[i][j] == 0 && !isInCache(cache, getCacheKey(i, j)) {
				currentCount := 0
				pointsToTraverse := []MatrixPoint{getCacheKey(i, j)}
				for len(pointsToTraverse) > 0 {
					var nextPoints []MatrixPoint
					for _, point := range pointsToTraverse {
						if !isInCache(cache, getCacheKey(point.i, point.j)) {
							currentCount++
							cache[getCacheKey(point.i, point.j)] = true
						}

						// check right
						if (point.j+1 < len(plotOfLand[point.i])) && plotOfLand[point.i][point.j+1] == 0 {
							nextPoints = append(nextPoints, getCacheKey(point.i, point.j+1))
						}

						// check down
						if (point.i+1 < len(plotOfLand)) && plotOfLand[point.i+1][point.j] == 0 {
							nextPoints = append(nextPoints, getCacheKey(point.i+1, point.j))
						}

						// check bottom left diagonal
						if (point.i+1 < len(plotOfLand) && point.j-1 >= 0) && plotOfLand[point.i+1][point.j-1] == 0 {
							nextPoints = append(nextPoints, getCacheKey(point.i+1, point.j-1))
						}

						// check bottom right diagonal
						if (point.i+1 < len(plotOfLand) && point.j+1 < len(plotOfLand[point.i+1])) && plotOfLand[point.i+1][point.j+1] == 0 {
							nextPoints = append(nextPoints, getCacheKey(point.i+1, point.j+1))
						}
					}

					pointsToTraverse = nextPoints
				}
				sizes = append(sizes, currentCount)
			}
		}
	}
	return sizes
}

func getCacheKey(i, j int) MatrixPoint {
	return MatrixPoint{i, j}
}

func isInCache(cache map[MatrixPoint]bool, key MatrixPoint) bool {
	_, ok := cache[key]
	if ok {
		return true
	}

	return false
}
