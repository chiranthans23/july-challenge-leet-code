package challenges

// IslandPerimeter -
func IslandPerimeter(grid [][]int) int {
	var perimeter int

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
