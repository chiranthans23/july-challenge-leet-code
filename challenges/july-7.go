package challenges

// IslandPerimeter -
func IslandPerimeter(grid [][]int) int {
	var perimeter int
	// visited := make([][]int, len(grid))
	// for i := range visited {
	// 	visited[i] = make([]int, len(grid))
	// }
	// i, j := 0, 0
	// for i = 0; i < len(grid); i++ {
	// 	for j = 0; j < len(grid); j++ {
	// 		if getValue(i, j, len(grid), grid) == 1 {
	// 			break
	// 		}
	// 	}
	// 	if getValue(i, j, len(grid), grid) == 1 {
	// 		break
	// 	}
	// }
	// fmt.Printf("Got x= %d and y= %d\n", i, j)
	// perimeter, grid, visited = addConnection(i, j, len(grid), perimeter, grid, visited)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {

			if grid[i][j] != 0 {
				perimeter += checkGrid(i, j, grid)
			}
		}
	}

	return perimeter
}
func checkGrid(i, j int, grid [][]int) int {
	edges := 0
	if i-1 < 0 || grid[i-1][j] == 0 {
		edges++
	}
	if j-1 < 0 || grid[i][j-1] == 0 {
		edges++
	}
	if i+1 >= len(grid) || grid[i+1][j] == 0 {
		edges++
	}
	if j+1 >= len(grid[0]) || grid[i][j+1] == 0 {
		edges++
	}
	return edges
}
func addConnection(i, j, n, perimeter int, grid [][]int, visited [][]int) (int, [][]int, [][]int) {
	if outOfMatrix(i, j, len(grid)) || getValue(i, j, len(grid), grid) != 1 {
		return perimeter, grid, visited
	}
	if !outOfMatrix(i, j, len(grid)) && visited[i][j] == 1 {
		return perimeter, grid, visited
	}
	if getValue(i, j, len(grid), grid) == 1 {
		perimeter += 4
		visited[i][j] = 1
	}
	if getValue(i, j-1, len(grid), grid) == 1 {
		perimeter--
	}
	if getValue(i, j+1, len(grid), grid) == 1 {
		perimeter--
	}
	if getValue(i+1, j, len(grid), grid) == 1 {
		perimeter--
	}
	if getValue(i-1, j, len(grid), grid) == 1 {
		perimeter--
	}
	perimeter, grid, visited = addConnection(i, j+1, len(grid), perimeter, grid, visited)
	perimeter, grid, visited = addConnection(i, j-1, len(grid), perimeter, grid, visited)
	perimeter, grid, visited = addConnection(i+1, j, len(grid), perimeter, grid, visited)
	perimeter, grid, visited = addConnection(i-1, j, len(grid), perimeter, grid, visited)
	return perimeter, grid, visited
}

// func findConnection(i, j, n, perimeter int, grid [][]int) (int, [][]int) {

// 	if getValue(i, j, len(grid), grid) != 1 {
// 		return perimeter, grid
// 	}
// 	grid[i][j] = -1
// 	if getValue(i, j-1, len(grid), grid) == 1 || getValue(i, j+1, len(grid), grid) == 1 || getValue(i+1, j, len(grid), grid) == 1 || getValue(i-1, j, len(grid), grid) == 1 {
// 		perimeter += 2
// 		perimeter, grid = findConnection(i, j+1, len(grid), perimeter, grid)
// 		perimeter, grid = findConnection(i, j-1, len(grid), perimeter, grid)
// 		perimeter, grid = findConnection(i+1, j, len(grid), perimeter, grid)
// 		perimeter, grid = findConnection(i-1, j, len(grid), perimeter, grid)

// 	} else {
// 		perimeter += 3
// 	}
// 	return perimeter, grid
// }

// func findIslandPoint(grid [][]int) (int, int) {
// 	for i := 0; i < len(grid); i++ {
// 		for j := 0; j < len(grid[i]); j++ {
// 			if getValue(i, j, len(grid), grid) == 1 && getValue(i, j-1, len(grid), grid) == 1 && getValue(i, j+1, len(grid), grid) == 1 && getValue(i+1, j, len(grid), grid) == 1 && getValue(i-1, j, len(grid), grid) == 1 {
// 				return i, j
// 			}
// 		}

// 	}
// 	return -1, -1
// }

func getValue(i, j, n int, grid [][]int) int {
	if i < n && j < n && i >= 0 && j >= 0 {
		return grid[i][j]
	}
	return -1
}
func outOfMatrix(i, j, n int) bool {
	if i < n && j < n && i >= 0 && j >= 0 {
		return false
	}
	return true
}
