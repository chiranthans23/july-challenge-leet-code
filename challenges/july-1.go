package challenges

import "math"

func arrangeCoins(n int) int {

	return int((-1 + math.Sqrt(1+float64(8*n))) / 2)
}
