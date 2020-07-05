package challenges

import (
	"testing"

	"github.com/chiranthans23/july-challenge-leet-code/util"
)

func TestHammingDistance(t *testing.T) {
	util.AssertEquals(t, hammingDistance(1, 4), 2)

}
