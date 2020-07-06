package challenges

import (
	"testing"

	"github.com/chiranthans23/july-challenge-leet-code/util"
)

func TestPlusOne(t *testing.T) {
	util.AssertEqualsIntArray(t, plusOne([]int{1, 2, 3, 4}), []int{1, 2, 3, 5})
	util.AssertEqualsIntArray(t, plusOne([]int{0, 0, 0, 0}), []int{0, 0, 0, 1})
	util.AssertEqualsIntArray(t, plusOne([]int{9, 9, 9, 9}), []int{1, 0, 0, 0, 0})
}
