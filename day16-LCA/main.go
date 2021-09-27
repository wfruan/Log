package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor( root *TreeNode ,  o1 int ,  o2 int ) int {
	if root == nil {
		return -1
	}

	node := isParent(root,o1,o2)
	if node != nil {
		return node.Val
	}
	return -1
}

func isParent(root *TreeNode, o1, o2 int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val == o1 || root.Val == o2 {
		return root
	}

	left := isParent(root.Left, o1, o2)
	right := isParent(root.Right, o1, o2)

	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}

func main()  {
	
}