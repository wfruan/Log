package AVLTree

// Insert 插入节点到平衡二叉树
func (tree *AVLTree) Insert(data int) {
	// 从根节点开始插入数据
	// 根节点在动态变化，所以需要不断刷新
	tree.root = tree.root.Insert(data)
}

func (node *AVLNode) Insert(data int) *AVLNode {
	// 如果节点为空，则初始化该节点
	if node == nil {
		return &AVLNode{Data: data, Height: 1}
	}
	// 如果值重复，则什么都不做
	if node.Data == data {
		return node
	}

	// 辅助变量，用于存储（旋转后）子树根节点
	var newTreeNode *AVLNode

	if data > node.Data {
		// 插入的值大于当前节点值，要从右子树插入
		node.Right = node.Right.Insert(data)
		// 计算插入节点后当前节点的平衡因子
		// 按照平衡二叉树的特征，平衡因子绝对值不能大于 1
		bf := node.BalanceFactor()
		// 如果右子树高度变高了，导致左子树-右子树的高度从 -1 变成了 -2
		if bf == -2 {
			if data > node.Right.Data {
				// 表示在右子树中插入右子节点导致失衡，需要单左旋
				newTreeNode = LeftRotate(node)
			} else {
				// 表示在右子树中插上左子节点导致失衡，需要先右后左双旋
				newTreeNode = RightLeftRotation(node)
			}
		}
	} else {
		// 插入的值小于当前节点值，要从左子树插入
		node.Left = node.Left.Insert(data)
		bf := node.BalanceFactor()
		// 左子树的高度变高了，导致左子树-右子树的高度从 1 变成了 2
		if bf == 2 {
			if data < node.Left.Data {
				// 表示在左子树中插入左子节点导致失衡，需要单右旋
				newTreeNode = RightRotate(node)
			} else {
				// 表示在左子树中插入右子节点导致失衡，需要先左后右双旋
				newTreeNode = LeftRightRotation(node)
			}
		}
	}

	if newTreeNode == nil {
		// 根节点没变，直接更新子树高度，并返回当前节点指针
		node.UpdateHeight()
		return node
	} else {
		// 经过旋转处理后，子树根节点变了，需要更新新的子树高度，然后返回新的子树根节点指针
		newTreeNode.UpdateHeight()
		return newTreeNode
	}
}

// UpdateHeight 更新节点树高度
func (node *AVLNode) UpdateHeight() {
	if node == nil {
		return
	}

	// 分别计算左子树和右子树的高度
	leftHeight, rightHeight := 0, 0
	if node.Left != nil {
		leftHeight = node.Left.Height
	}
	if node.Right != nil {
		rightHeight = node.Right.Height
	}

	// 以更高的子树高度作为节点树高度
	maxHeight := leftHeight
	if rightHeight > maxHeight {
		maxHeight = rightHeight
	}

	// 最终高度要加上节点本身所在的那一层
	node.Height = maxHeight + 1
}

// BalanceFactor 计算节点平衡因子（即左右子树的高度差）
func (node *AVLNode) BalanceFactor() int {
	leftHeight, rightHeight := 0, 0
	if node.Left != nil {
		leftHeight = node.Left.Height
	}
	if node.Right != nil {
		rightHeight = node.Right.Height
	}
	return leftHeight - rightHeight
}

// RightRotate 右旋操作
func RightRotate(node *AVLNode) *AVLNode {
	pivot := node.Left    // pivot 表示新插入的节点
	pivotR := pivot.Right // 暂存 pivot 右子树入口节点
	pivot.Right = node    // 右旋后最小不平衡子树根节点 node 变成 pivot 的右子节点
	node.Left = pivotR    // 而 pivot 原本的右子节点需要挂载到 node 节点的左子树上

	// 只有 node 和 pivot 的高度改变了
	node.UpdateHeight()
	pivot.UpdateHeight()

	// 返回右旋后的子树根节点指针，即 pivot
	return pivot
}

// LeftRotate 左旋操作
func LeftRotate(node *AVLNode) *AVLNode {
	pivot := node.Right  // pivot 表示新插入的节点
	pivotL := pivot.Left // 暂存 pivot 左子树入口节点
	pivot.Left = node    // 左旋后最小不平衡子树根节点 node 变成 pivot 的左子节点
	node.Right = pivotL  // 而 pivot 原本的左子节点需要挂载到 node 节点的右子树上

	// 只有 node 和 pivot 的高度改变了
	node.UpdateHeight()
	pivot.UpdateHeight()

	// 返回旋后的子树根节点指针，即 pivot
	return pivot
}

// LeftRightRotation 双旋操作（先左后右）
func LeftRightRotation(node *AVLNode) *AVLNode {
	node.Left = LeftRotate(node.Left)
	return RightRotate(node)
}

// RightLeftRotation 先右旋后左旋
func RightLeftRotation(node *AVLNode) *AVLNode {
	node.Right = RightRotate(node.Right)
	return LeftRotate(node)
}
