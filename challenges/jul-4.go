package challenges

// NthUglyNumber -
func NthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}
	var uglyNumbers [1690]int
	uglyNumbers[0] = 1
	var list2, list3, list5 int
	for i := 1; i < n; i++ {
		uglyNumbers[i] = min(uglyNumbers[list2]*2, min(uglyNumbers[list3]*3, uglyNumbers[list5]*5))
		if uglyNumbers[i] == (2 * uglyNumbers[list2]) {
			list2++
		}
		if uglyNumbers[i] == (3 * uglyNumbers[list3]) {
			list3++
		}
		if uglyNumbers[i] == (5 * uglyNumbers[list5]) {
			list5++
		}

	}
	return uglyNumbers[n-1]
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
