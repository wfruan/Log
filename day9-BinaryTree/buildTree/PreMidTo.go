package buildTree

func PreMidBuildTree(preorder []int, inorder []int) *TreeNode  {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0],nil,nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = PreMidBuildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = PreMidBuildTree(preorder[len(inorder[:i])+1:],inorder[i+1:])
	return root
}
