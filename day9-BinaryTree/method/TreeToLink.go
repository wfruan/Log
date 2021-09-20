package method

func Flatten(root *PNode)  {
	if root == nil {
		return
	}

	Flatten(root.Left)
	Flatten(root.Right)

	left := root.Left
	right := root.Right

	root.Left = nil
	root.Right = left

	p := root

	for p.Right != nil {
		p = p.Right
	}
	p.Right = right

}
