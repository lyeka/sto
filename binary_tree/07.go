package binary_tree

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	node := &TreeNode{Val: preorder[0]}
	node.Left = buildTree(preorder[1:len(inorder[:i])], inorder[0:i])
	node.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return node
}
