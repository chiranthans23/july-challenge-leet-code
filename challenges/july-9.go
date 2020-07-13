package challenges

import "math"

func widthOfBinaryTree(root *TreeNode) int {
	width := 0
	bfs([]*TreeNode{root}, []int{0}, &width)
	return width
}

func bfs(nodes []*TreeNode, order []int, width *int) {
	if len(nodes) == 0 {
		return
	}
	*width = max(*width, order[len(nodes)-1]-order[0]+1)
	nextLevel := []*TreeNode{}
	nextOrder := []int{}
	for i, v := range nodes {
		if v.Left != nil {
			nextLevel = append(nextLevel, v.Left)
			nextOrder = append(nextOrder, order[i]*2)
		}
		if v.Right != nil {
			nextLevel = append(nextLevel, v.Right)
			nextOrder = append(nextOrder, order[i]*2+1)
		}
	}
	bfs(nextLevel, nextOrder, width)
}

func max(values ...int) int {
	maxValue := math.MinInt32
	for _, v := range values {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}
