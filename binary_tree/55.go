package binary_tree

// DFS
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return maxInt(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// BFS
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	nodes := []*TreeNode{root}
	dep := 0
	for len(nodes) != 0 {
		tmp := []*TreeNode{}
		for _, node := range nodes {
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		nodes = tmp
		dep += 1
	}
	return dep
}
