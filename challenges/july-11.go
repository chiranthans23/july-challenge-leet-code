package challenges

func Subsets(nums []int) [][]int {
	subsets := [][]int{[]int{}}
	for _, num := range nums {
		for _, subset := range subsets {
			newSubset := append([]int{}, subset...)
			newSubset = append(newSubset, num)
			subsets = append(subsets, newSubset)
		}
		// fmt.Printf("Subsets %v", subsets)
	}
	return subsets
}
