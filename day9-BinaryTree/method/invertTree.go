package method

func (tree *BinarySearchTree) Invert(root *Node) *Node {
	// base case
	if root == nil {
		return nil
	}

	/**** 前序遍历位置 ****/
	// root 节点需要交换它的左右子节点
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	// 让左右子节点继续翻转它们的子节点
	tree.Invert(root.Left)
	tree.Invert(root.Right)

	return root
}