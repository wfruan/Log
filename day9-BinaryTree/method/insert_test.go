package method

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree_Insert(t *testing.T) {
	tree := NewBinarySearchTree(nil)
	err := tree.Insert(3)
	if err != nil {
		return
	}
	err = tree.Insert(2)
	if err != nil {
		return
	}
	err = tree.Insert(5)
	if err != nil {
		return
	}
	err = tree.Insert(1)
	if err != nil {
		return
	}
	err = tree.Insert(4)
	if err != nil {
		return
	}

	fmt.Println("中序遍历二叉排序树：")
	MidOrderTraverse(tree.Root)
	fmt.Println()
}
