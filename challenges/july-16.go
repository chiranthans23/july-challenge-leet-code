package challenges

func myPow(x float64, n int) float64 {
	var result float64
	result = 1
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if x == 1 {
		return 1
	}
	if x == -1 {
		if n%2 == 0 {
			return 1
		}
		return -1
	}
	if n > 0 {
		for n > 1 {
			result *= x * x
			n -= 2
		}
		if n == 1 {
			result *= x
			n--
		}
	} else {
		for n < -1 {
			result /= x * x
			n += 2
		}
		if n == -1 {
			result /= x
			n++
		}
	}

	return result
}
