package method

//PNode represent Perfect-Binary-Tree Node
type PNode struct {
	Val int
	Left *PNode
	Right *PNode
	next *PNode
}

func Connect(root *PNode) *PNode  {
	if root == nil {
		return nil
	}
	ConnectTwoNode(root.Left,root.Right)
	return root
}

func ConnectTwoNode(node1 *PNode, node2 *PNode)  {
	if node1 == nil || node2 == nil {
		return
	}

	node1.next = node2

	ConnectTwoNode(node1.Left,node1.Right)
	ConnectTwoNode(node2.Left,node2.Right)

	ConnectTwoNode(node1.Right,node2.Left)
}