package challenges

import (
	"fmt"
	"reflect"
)

// PrisonCellsNDays -
func PrisonCellsNDays(cells []int, N int) []int {
	var cycle int
	var pattern [8]int

	for i := N; i > 0; i-- {
		var temp [8]int
		for j := 1; j <= 6; j++ {
			temp[j] = xnor(cells[j-1], cells[j+1])
		}
		if cycle == 0 {
			copy(pattern[:], temp[:])
		} else if reflect.DeepEqual(temp, pattern) {
			i %= cycle
			fmt.Println(i)
		}
		if i == 0 {
			return cells
		}

		copy(cells, temp[:])
		cycle++
	}

	return cells
}

func xnor(x, y int) int {

	if x != y {
		return 0
	}
	return 1

}
