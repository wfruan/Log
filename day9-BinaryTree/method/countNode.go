package method

// Count func (tree *BinarySearchTree) Find(data int) *Node
func  Count(root *Node) int {
	// base case
	if root == nil{
		return 0
	}

	// 自己加上子树的节点数就是整棵树的节点数
	return 1 + Count(root.Left) + Count(root.Right)
}
