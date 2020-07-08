package challenges

import (
	"reflect"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var triples [][]int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {

			if j+1 < len(nums) && binarySearch(nums[j+1:], -1*(nums[i]+nums[j])) != -1 {
				trip := []int{nums[i], nums[j], -(nums[i] + nums[j])}
				if isDup(triples, trip) {
					continue
				}
				triples = append(triples, trip)
			}

		}
	}
	return triples
}
func isDup(trips [][]int, nums []int) bool {
	for i := 0; i < len(trips); i++ {
		if reflect.DeepEqual(trips[i], nums) {
			return true
		}
	}
	return false
}
func binarySearch(nums []int, key int) int {
	low := 0
	high := len(nums) - 1
	mid := -1
	for low <= high {
		mid := int((low + high) / 2)
		if nums[mid] == key {
			return mid
		}
		if nums[mid] > key {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return mid
}
