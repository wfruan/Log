package method

import "fmt"

// PreOrderTraverse 前序遍历
func PreOrderTraverse(treeNode *Node) {
	// 节点为空则退出当前递归
	if treeNode == nil {
		return
	}
	// 否则先打印当前节点值
	fmt.Printf("%s ", treeNode.GetData())
	// 然后对左子树和右子树做前序遍历
	PreOrderTraverse(treeNode.Left)
	PreOrderTraverse(treeNode.Right)
}

// MidOrderTraverse 中序遍历
func MidOrderTraverse(treeNode *Node) {
	// 节点为空则退出当前递归
	if treeNode == nil {
		return
	}
	// 否则先从左子树最左侧节点开始遍历
	MidOrderTraverse(treeNode.Left)
	// 打印位于中间的根节点
	fmt.Printf("%s ", treeNode.GetData())
	// 最后按照和左子树一样的逻辑遍历右子树
	MidOrderTraverse(treeNode.Right)
}

// PostOrderTraverse 后序遍历
func PostOrderTraverse(treeNode *Node) {
	// 节点为空则退出当前递归
	if treeNode == nil {
		return
	}
	// 否则先遍历左子树
	PostOrderTraverse(treeNode.Left)
	// 再遍历右子树
	PostOrderTraverse(treeNode.Right)
	// 最后访问根节点
	fmt.Printf("%s ", treeNode.GetData())
}
