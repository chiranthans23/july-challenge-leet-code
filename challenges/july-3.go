package challenges

import "fmt"

// PrisonCellsNDays -
func PrisonCellsNDays(cells []int, N int) []int {
	var previousDayCells [8]int

	for i := 1; i <= N; i++ {
		copy(previousDayCells[:], cells)
		for j := 1; j <= 6; j++ {
			cells[j] = xnor(previousDayCells[j-1], previousDayCells[j+1])
		}
		if i == 1 {
			cells[0] = 0
			cells[7] = 0
		}
		fmt.Printf("Cells after %d days:", i)
		fmt.Println(cells)
	}
	return cells
}

func xnor(x, y int) int {

	if x != y {
		return 0
	}
	return 1

}
