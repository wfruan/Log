package AVLTree

import (
	"fmt"
	"math"
)

// IsAVLTree 判断是不是平衡二叉树
func (tree *AVLTree) IsAVLTree() bool {
	if tree == nil || tree.root == nil {
		return true
	}

	// 判断每个节点是否符合平衡二叉树的定义
	if tree.root.IsBalanced() {
		return true
	}

	return false
}

// IsBalanced 判断节点是否符合平衡二叉树的定义
func (node *AVLNode) IsBalanced() bool {
	if node == nil {
		return true
	}

	// 左右子树都为空是叶子节点
	if node.Left == nil && node.Right == nil {
		// 叶子节点高度都是 1
		if node.Height == 1 {
			return true
		} else {
			fmt.Println("叶子节点高度值: ", node.Height)
			return false
		}
	} else if node.Left != nil && node.Right != nil {
		// 左右子树不为空
		// 左子树所有节点值必须比父节点小，右子树所有节点值必须比父节点大（AVL 树首先是二叉排序树）
		if node.Left.Data > node.Data || node.Right.Data < node.Data {
			// 不符合 AVL 树定义
			fmt.Printf("父节点值是 %v, 左子节点值是 %v, 右子节点值是 %v\n", node.Data, node.Left.Data, node.Right.Data)
			return false
		}

		// 计算平衡因子 BF 绝对值
		bf := node.Left.Height - node.Right.Height

		// 平衡因子不能大于 1
		if math.Abs(float64(bf)) > 1 {
			fmt.Println("平衡因子 BF 值: ", bf)
			return false
		}

		// 如果左子树比右子树高，那么父节点的高度等于左子树 +1
		if node.Left.Height > node.Right.Height {
			if node.Height != node.Left.Height+1 {
				fmt.Printf("%#v 高度: %v, 左子树高度: %v, 右子树高度: %v\n", node, node.Height, node.Left.Height, node.Right.Height)
				return false
			}
		} else {
			// 如果右子树比左子树高，那么父节点的高度等于右子树 +1
			if node.Height != node.Right.Height+1 {
				fmt.Printf("%#v 高度: %v, 左子树高度: %v, 右子树高度: %v\n", node, node.Height, node.Left.Height, node.Right.Height)
				return false
			}
		}

		// 递归判断左子树
		if !node.Left.IsBalanced() {
			return false
		}

		// 递归判断右子树
		if !node.Right.IsBalanced() {
			return false
		}

	} else {
		// 只存在一棵子树
		if node.Right != nil {
			// 子树高度只能是 1
			if node.Right.Height == 1 && node.Right.Left == nil && node.Right.Right == nil {
				if node.Right.Data < node.Data {
					// 右子节点值必须比父节点值大
					fmt.Printf("节点值: %v,(%#v,%#v)", node.Data, node.Right, node.Left)
					return false
				}
			} else {
				fmt.Printf("节点值: %v,(%#v,%#v)", node.Data, node.Right, node.Left)
				return false
			}
		} else {
			if node.Left.Height == 1 && node.Left.Left == nil && node.Left.Right == nil {
				if node.Left.Data > node.Data {
					// 左子节点值必须比父节点值小
					fmt.Printf("节点值: %v,(%#v,%#v) child", node.Data, node.Right, node.Left)
					return false
				}
			} else {
				fmt.Printf("节点值: %v,(%#v,%#v) child", node.Data, node.Right, node.Left)
				return false
			}
		}
	}

	return true
}
