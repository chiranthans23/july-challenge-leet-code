package challenges

func zigzagLevelOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}
	order := addNodesInOrder(root, [][]int{}, true, 0)

	return order
}
func addNodesInOrder(root *TreeNode, order [][]int, direction bool, level int) [][]int {
	if root == nil {
		return order
	}
	if level >= len(order) {
		order = append(order, []int{})
	}
	if !direction {
		order[level] = append(order[level], root.Val)
	} else {
		order[level] = append([]int{root.Val}, order[level]...)
	}

	order = addNodesInOrder(root.Left, order, !direction, level+1)
	order = addNodesInOrder(root.Right, order, !direction, level+1)

	return order
}
