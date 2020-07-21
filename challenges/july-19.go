package challenges

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	var carry int
	var result string
	a, b = appendZeroes(a, b)
	A := []rune(a)
	B := []rune(b)
	for i, j := len(A)-1, len(B)-1; i >= 0 && j >= 0; {
		var str string
		str, carry = addBits(A[i], B[j], carry)
		fmt.Printf("After adding a bit I have %s as value and %d as carry", str, carry)
		result = str + result
		i--
		j--
	}
	if carry == 1 {
		return strconv.Itoa(carry) + result
	}
	return result
}
func addBits(a, b rune, carry int) (string, int) {
	if string(a) != string(b) {
		if carry == 1 {
			return "0", 1
		}
		return "1", 0
	}
	if string(a) == "0" {
		return strconv.Itoa(carry), 0
	}
	if carry == 1 {
		return "1", 1
	}
	return "0", 1
}
func appendZeroes(a, b string) (string, string) {
	n, m := len(a), len(b)
	if n > m {
		for i := 0; i < n-m; i++ {
			b = "0" + b
		}
		return a, b
	}
	for i := 0; i < m-n; i++ {
		a = "0" + a
	}
	return a, b
}
