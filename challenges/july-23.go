package challenges

func singleNumber(nums []int) []int {
	num1 := make(map[int]int)

	if len(nums) == 2 {
		return []int{nums[0], nums[1]}
	}

	for i := 0; i < len(nums); i++ {
		if num1[nums[i]] == 1 {
			delete(num1, nums[i])
		} else {
			num1[nums[i]] = 1
		}

	}
	i := 0
	for k := range num1 {
		nums[i] = k
		i++
	}
	return nums[:2]
}
