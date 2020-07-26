package challenges

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	if A == B {
		set := make(map[byte]struct{}, len(A))
		for i := range A {
			set[A[i]] = struct{}{}
		}
		return len(set) < len(A)
	}
	i, j := -1, -1
	for k := range A {
		if A[k] == B[k] {
			continue
		}
		if i == -1 {
			i = k
		} else if j == -1 {
			j = k
		} else {
			return false
		}
	}
	if i != -1 && j != -1 {
		return A[i] == B[j] && A[j] == B[i]
	}
	return false
}
