package challenges

func exist(board [][]byte, word string) bool {
	// base case
	if len(board) == 0 {
		return false
	}

	// itereating over the matrix
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if DFS(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

// DFS -
func DFS(board [][]byte, i int, j int, word string, index int) bool {
	// knowing if I found the pattern
	if index == len(word) {
		return true
	}

	// boundary conditions
	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 || word[index] != board[i][j] {
		return false
	}

	// * is not present in the pattern or word
	temp := board[i][j]
	board[i][j] = '*'

	a := DFS(board, i+1, j, word, index+1) || DFS(board, i, j+1, word, index+1) || DFS(board, i-1, j, word, index+1) || DFS(board, i, j-1, word, index+1)

	board[i][j] = temp
	return a
}
