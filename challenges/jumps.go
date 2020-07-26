package challenges

func canReach(arr []int, start int) bool {
	visited := make([]bool, len(arr))
	return dfsJump(start, arr, visited)
}

func dfsJump(start int, arr []int, visited []bool) bool {
	// I am interested in only finding 0
	if arr[start] == 0 {
		return true
	}
	var front, back bool
	visited[start] = true

	// Right part
	if start+arr[start] < len(arr) && !visited[start+arr[start]] {
		front = dfsJump(start+arr[start], arr, visited)
	}

	// Left part
	if start-arr[start] >= 0 && !visited[start-arr[start]] {
		back = dfsJump(start-arr[start], arr, visited)
	}

	// Either of them can be true
	return front || back
}
