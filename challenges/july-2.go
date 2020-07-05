package challenges

// TreeNode -
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LevelOrderBottom -
func LevelOrderBottom(root *TreeNode) [][]int {
	var levelOrder [][]int
	if root == nil {
		return levelOrder
	}
	levelOrder = addNodesLevelOrder(root, levelOrder, 0)
	for i, j := 0, len(levelOrder)-1; i < j; i, j = i+1, j-1 {
		levelOrder[i], levelOrder[j] = levelOrder[j], levelOrder[i]
	}
	return levelOrder
}

func addNodesLevelOrder(root *TreeNode, levelOrder [][]int, level int) [][]int {
	if root == nil {
		return levelOrder
	}
	if level > len(levelOrder)-1 {
		levelOrder = append(levelOrder, []int{})
	}
	levelOrder[level] = append(levelOrder[level], root.Val)
	levelOrder = addNodesLevelOrder(root.Left, levelOrder, level+1)
	levelOrder = addNodesLevelOrder(root.Right, levelOrder, level+1)
	return levelOrder
}
