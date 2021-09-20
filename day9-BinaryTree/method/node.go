package method

import "fmt"

// Node 通过链表存储二叉树节点信息
type Node struct {
	Data  interface{}
	Left  *Node
	Right *Node
}

func NewNode(data interface{}) *Node {
	return &Node{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

func (node *Node) GetData() string {
	return fmt.Sprintf("%v", node.Data)
}


