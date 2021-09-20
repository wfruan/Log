package AVLTree

// Delete 删除指定节点
func (tree *AVLTree) Delete(data int) {
	// 空树直接返回
	if tree.root == nil {
		return
	}

	// 删除指定节点，和插入节点一样，根节点也会随着 AVL 树的旋转动态变化
	tree.root = tree.root.Delete(data)
}

func (node *AVLNode) Delete(data int) *AVLNode {
	// 空节点直接返回 nil
	if node == nil {
		return nil
	}
	if data < node.Data {
		// 如果删除节点值小于当前节点值，则进入当前节点的左子树删除元素
		node.Left = node.Left.Delete(data)
		// 删除后要更新左子树高度
		node.Left.UpdateHeight()
	} else if data > node.Data {
		// 如果删除节点值大于当前节点值，则进入当前节点的右子树删除元素
		node.Right = node.Right.Delete(data)
		// 删除后要更新右子树高度
		node.Right.UpdateHeight()
	} else {
		// 找到待删除节点后
		// 第一种情况，删除的节点没有左右子树，直接删除即可
		if node.Left == nil && node.Right == nil {
			// 返回 nil 表示直接删除该节点
			return nil
		}

		// 第二种情况，待删除节点包含左右子树，选择高度更高的子树下的节点来替换待删除的节点
		// 如果左子树更高，选择左子树中值最大的节点，也就是左子树最右边的叶子节点
		// 如果右子树更高，选择右子树中值最小的节点，也就是右子树最左边的叶子节点
		// 最后，删除这个叶子节点即可
		if node.Left != nil && node.Right != nil {
			// 左子树更高，拿左子树中值最大的节点替换
			if node.Left.Height > node.Right.Height {
				maxNode := node.Left
				for maxNode.Right != nil {
					maxNode = maxNode.Right
				}

				// 将值最大的节点值赋值给待删除节点
				node.Data = maxNode.Data

				// 然后把值最大的节点删除
				node.Left = node.Left.Delete(maxNode.Data)
				// 删除后要更新左子树高度
				node.Left.UpdateHeight()
			} else {
				// 右子树更高，拿右子树中值最小的节点替换
				minNode := node.Right
				for minNode.Left != nil {
					minNode = minNode.Left
				}

				// 将值最小的节点值赋值给待删除节点
				node.Data = minNode.Data

				// 然后把值最小的节点删除
				node.Right = node.Right.Delete(minNode.Data)
				// 删除后要更新右子树高度
				node.Right.UpdateHeight()
			}
		} else {
			// 只有左子树或只有右子树
			// 只有一棵子树，该子树也只是一个节点，则将该节点值赋值给待删除的节点，然后置子树为空
			if node.Left != nil {
				// 第三种情况，删除的节点只有左子树
				// 根据 AVL 树的特征，可以知道左子树其实就只有一个节点，否则高度差就等于 2 了
				node.Data = node.Left.Data
				node.Height = 1
				node.Left = nil
			} else if node.Right != nil {
				// 第四种情况，删除的节点只有右子树
				// 根据 AVL 树的特征，可以知道右子树其实就只有一个节点，否则高度差就等于 2 了
				node.Data = node.Right.Data
				node.Height = 1
				node.Right = nil
			}
		}

		// 找到指定值对应的待删除节点并进行替换删除后，直接返回该节点
		return node
	}

	// 左右子树递归删除节点后需要平衡
	var newNode *AVLNode
	// 相当删除了右子树的节点，左边比右边高了，不平衡
	if node.BalanceFactor() == 2 {
		if node.Left.BalanceFactor() >= 0 {
			newNode = RightRotate(node)
		} else {
			newNode = LeftRightRotation(node)
		}
		//  相当删除了左子树的节点，右边比左边高了，不平衡
	} else if node.BalanceFactor() == -2 {
		if node.Right.BalanceFactor() <= 0 {
			newNode = LeftRotate(node)
		} else {
			newNode = RightLeftRotation(node)
		}
	}

	if newNode == nil {
		node.UpdateHeight()
		return node
	} else {
		newNode.UpdateHeight()
		return newNode
	}
}
