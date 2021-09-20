package AVLTree

// Find 查找指定节点
func (tree *AVLTree) Find(data int) *AVLNode {
	// 空树直接返回空
	if tree.root == nil {
		return nil
	}

	return tree.root.Find(data)
}

func (node *AVLNode) Find(data int) *AVLNode {
	if data == node.Data {
		return node
	} else if data < node.Data {
		// 如果查找的值小于节点值，从节点的左子树开始查找
		if node.Left == nil {
			return nil
		}
		return node.Left.Find(data)
	} else {
		// 如果查找的值大于节点值，从节点的右子树开始查找
		if node.Right == nil {
			return nil
		}
		return node.Right.Find(data)
	}
}
