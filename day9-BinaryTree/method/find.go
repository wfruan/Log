package method

// Find 查找值为data的节点
func (tree *BinarySearchTree) Find(data int) *Node {
	if node := tree.Root; node == nil {
		// 二叉排序树为空返回空
		return nil
	} else {
		// 否则返回值等于data的节点指针
		for node != nil {
			if data < node.Data.(int) {
				node = node.Left
			} else if data > node.Data.(int) {
				node = node.Right
			} else {
				return node
			}
		}
		return nil
	}
}
