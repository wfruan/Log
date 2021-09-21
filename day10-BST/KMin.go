package day10_BST

func KthSmallest(root *TreeNode, k int) int {
	i = 0
	return run(root, k)
}

var i int

func run(root *TreeNode, k int) int {
	if root == nil {
		return -1
	}
	left := run(root.Left, k)
	if left != -1 {
		return left
	}
	i++
	if i == k {
		return root.Val
	}
	right := run(root.Right, k)
	if right != -1 {
		return right
	}
	return -1
}

