package AVLTree

import "fmt"

// Traverse 中序遍历平衡二叉树
func (tree *AVLTree) Traverse() {
	// 从根节点开始遍历
	tree.root.Traverse()
}

func (node *AVLNode) Traverse() {
	// 节点为空则退出当前递归
	if node == nil {
		return
	}
	// 否则先从左子树最左侧节点开始遍历
	node.Left.Traverse()
	// 打印位于中间的根节点
	fmt.Printf("%d(%d) ", node.Data, node.BalanceFactor())
	// 最后按照和左子树一样的逻辑遍历右子树
	node.Right.Traverse()
}

