package method

import (
	"errors"
	"fmt"
)

// Insert 插入数据到二叉排序树
func (tree *BinarySearchTree) Insert(data interface{}) error {
	var value int
	var ok bool
	if value, ok = data.(int); !ok {
		return errors.New("暂时只支持int类型数据")
	}
	if node := tree.Root; node == nil {
		// 二叉排序树为空，则初始化
		tree.Root = NewNode(value)
		return nil
	} else {
		// 否则找到要插入的位置插入新的节点
		for node != nil {
			if value < node.Data.(int) {
				if node.Left == nil {
					node.Left = NewNode(value)
					break
				}
				node = node.Left
			} else if value > node.Data.(int) {
				if node.Right == nil {
					node.Right = NewNode(value)
					break
				}
				node = node.Right
			} else {
				return errors.New("对应数据已存在，请勿重复插入")
			}
		}
		return nil
	}
}

func main()  {
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