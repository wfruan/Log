package method

type BinarySearchTree struct {
	Root *Node
}

// NewBinarySearchTree 初始化二叉排序树
func NewBinarySearchTree(node *Node) *BinarySearchTree {
	return &BinarySearchTree{
		Root: node,
	}
}
