package binary_tree

func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}

	if absInt(maxDepth(root.Left), maxDepth(root.Right)) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func absInt(x, y int) int {
	t := x - y
	if t > 0 {
		return t
	}
	return 0 - t
}
