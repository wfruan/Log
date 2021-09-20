package method

import "errors"

// Delete 删除值为 data 的节点
func (tree *BinarySearchTree) Delete(data int) error {
	if tree.Root == nil {
		return errors.New("二叉树为空")
	}

	p := tree.Root
	var pp *Node = nil // p 的父节点

	// 找到待删除节点
	for p != nil && p.Data.(int) != data {
		pp = p
		if p.Data.(int) < data {
			// 当前节点值小于待删除值，去右子树查找
			p = p.Right
		} else {
			// 否则去左子树查找
			p = p.Left
		}
	}
	if p == nil {
		return errors.New("待删除节点不存在")
	}

	// 待删除节点有两个子节点，需要将待删除节点值设置为该节点右子树最小节点值，再删除右子树最小节点
	if p.Left != nil && p.Right != nil {
		minP := p.Right // 右子树中的最小节点
		minPP := p      // minP 的父节点
		// 查找右子树中的最小节点
		for minP.Left != nil {
		minPP = minP
		minP = minP.Left
	}
		p.Data = minP.Data // 将 minP 的数据设置到 p 中
		p = minP           // 下面就变成删除 minP 了
		pp = minPP
	}

	// 应用待删除节点只有左/右子节点的逻辑
	var child *Node
	if p.Left != nil {
		child = p.Left
	} else if p.Right != nil {
		child = p.Right
	} else {
		child = nil
	}

	if pp == nil {
		// 删除跟节点
		tree.Root = nil
	} else if pp.Left == p {
		// 仅有左子节点
		pp.Left = child
	} else if pp.Right == p {
		// 仅有右子节点
		pp.Right = child
	}

	return nil
}
