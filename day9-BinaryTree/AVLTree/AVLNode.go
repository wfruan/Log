package AVLTree

// AVLNode 平衡二叉树节点类
type AVLNode struct {
	Data        int
	Left, Right *AVLNode
	Height      int // 以该节点作为根节点对应子树的高度，用于计算平衡因子
}

// AVLTree 平衡二叉树结构体
type AVLTree struct {
	root *AVLNode // 根节点
}

// NewAVLTree 平衡二叉树构造函数
func NewAVLTree(data int) *AVLTree {
	return &AVLTree{
		root: &AVLNode{Data: data, Height: 1},
	}
}
