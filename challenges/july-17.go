package challenges

import (
	"fmt"
	"math"
)

func topKFrequent(nums []int, k int) []int {
	frequencyTable := make(map[int]int)
	for _, ele := range nums {
		frequencyTable[ele]++
	}
	fmt.Printf("Frequency table is %v\n", frequencyTable)
	var result []int
	for i := 1; i <= k; i++ {
		result = append(result, getMax(frequencyTable))
		fmt.Printf("After %d iteration I got %d\n", i, result[i-1])
	}
	return result
}

func getMax(freqTable map[int]int) int {
	max := math.MinInt32
	key := 0
	for k, v := range freqTable {
		if v > max {
			max = v
			key = k
		}
	}
	delete(freqTable, key)
	return key
}
