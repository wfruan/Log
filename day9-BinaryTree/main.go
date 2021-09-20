package main

import (
	"day9-BinaryTree/AVLTree"
	"day9-BinaryTree/method"
	"fmt"
)

// 后序遍历
func main() {
	// 初始化一个简单的二叉树
	node1 := method.NewNode(0) // 根节点
	node2 := method.NewNode("1")
	node3 := method.NewNode(2.0)
	node1.Left = node2
	node1.Right = node3

	// 前序遍历这个二叉树
	fmt.Print("前序遍历: ")
	method.PreOrderTraverse(node1)
	fmt.Println()

	// 中序遍历这个二叉树
	fmt.Print("中序遍历: ")
	method.MidOrderTraverse(node1)
	fmt.Println()

	// 后序遍历这个二叉树
	fmt.Print("后序遍历: ")
	method.PostOrderTraverse(node1)
	fmt.Println()

	tree := method.NewBinarySearchTree(nil)
	err := tree.Insert(3)
	if err != nil {
		fmt.Printf("%v\n", err)
	}else {
		fmt.Println("插入3成功")
	}
	err = tree.Insert(3)
	if err != nil {
		fmt.Printf("%v\n", err)
	}else {
		fmt.Println("插入3成功")
	}
	err = tree.Insert(2)
	if err != nil {
		fmt.Printf("%v\n", err)
	}else {
		fmt.Println("插入2成功")
	}
	err = tree.Insert(2)
	if err != nil {
		fmt.Printf("%v\n", err)
	}else {
		fmt.Println("插入2成功")
	}
	err = tree.Insert(5)
	if err != nil {
		fmt.Printf("%v\n", err)
	}else {
		fmt.Println("插入5成功")
	}
	err = tree.Insert(1)
	if err != nil {
		fmt.Printf("%v\n", err)
	}else {
		fmt.Println("插入1成功")
	}
	err = tree.Insert(4)
	if err != nil {
		fmt.Printf("%v\n", err)
	}else {
		fmt.Println("插入4成功")
	}

	fmt.Println(tree.Root)
	fmt.Println("中序遍历二叉排序树：")
	method.MidOrderTraverse(tree.Root)
	fmt.Println()

	fmt.Println("查找值为 3 的节点：")
	node := tree.Find(3)
	fmt.Printf("%v\n", node)

	fmt.Println("查找值为 10 的节点：")
	node = tree.Find(10)
	fmt.Printf("%v\n", node)


	fmt.Println(method.Count(tree.Root))

	fmt.Println("删除值为 3 的节点：")
	err = tree.Delete(3)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println("删除成功")
	}

	fmt.Println(method.Count(tree.Root))


	fmt.Println("中序遍历删除节点后的二叉排序树：")
	method.MidOrderTraverse(tree.Root)
	fmt.Println()

	fmt.Println("生成镜像二叉树")
	tree.Invert(tree.Root)
	method.MidOrderTraverse(tree.Root)
	fmt.Println()

	// 初始化 AVL 树
	avlTree := AVLTree.NewAVLTree(3)
	// 插入节点
	avlTree.Insert(2)
	avlTree.Insert(1)
	avlTree.Insert(4)
	avlTree.Insert(5)
	avlTree.Insert(6)
	avlTree.Insert(7)
	avlTree.Insert(10)
	avlTree.Insert(9)
	avlTree.Insert(8)
	// 判断是否是平衡二叉树
	fmt.Print("avlTree 是平衡二叉树: ")
	fmt.Println(avlTree.IsAVLTree())
	// 中序遍历生成的二叉树看是否是二叉排序树
	fmt.Print("中序遍历结果: ")
	avlTree.Traverse()
	fmt.Println()
	// 查找节点
	fmt.Print("查找值为 5 的节点: ")
	fmt.Printf("%v\n", avlTree.Find(5))
	// 删除节点
	avlTree.Delete(5)
	// 删除后是否还是平衡二叉树
	fmt.Print("avlTree 仍然是平衡二叉树: ")
	fmt.Println(avlTree.IsAVLTree())
	fmt.Print("删除节点后的中序遍历结果: ")
	avlTree.Traverse()
	fmt.Println()


	fmt.Println(method.Count(tree.Root))
	//method.Count(tree.Root)
	//method.Count(avlTree.Root)
}
