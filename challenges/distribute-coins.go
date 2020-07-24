package challenges

func distributeCoins(root *TreeNode) int {
	ans := 0
	valuesToBalance(root, &ans)
	return ans
}

func valuesToBalance(node *TreeNode, ans *int) int {
	if node == nil {
		return 0
	}
	l := valuesToBalance(node.Left, ans)
	r := valuesToBalance(node.Right, ans)
	*ans += abs(l) + abs(r)
	return l + r + node.Val + -1
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
