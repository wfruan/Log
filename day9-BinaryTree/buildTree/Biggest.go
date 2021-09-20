package buildTree

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) < 1 {
		return nil
	}
	maxIndex := Max(nums)
	root := &TreeNode{
		Val:   nums[maxIndex],
		Left:  constructMaximumBinaryTree(nums[:maxIndex]),
		Right: constructMaximumBinaryTree(nums[maxIndex+1:]),
	}
	return root
}

func Max(nums []int) int  {
	maxIndex := 1
	for i := 1; i < len(nums); i++ {
		if nums[maxIndex] < nums[i] {
			maxIndex = i
		}
	}
	return maxIndex
}

/*func ceonstructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums)<1{return nil}
	//首选找到最大值
	index:=findMax(nums)
	//其次构造二叉树
	root:=&TreeNode{
		Val: nums[index],
		Left:ceonstructMaximumBinaryTree(nums[:index]),//左半边
		Right:ceonstructMaximumBinaryTree(nums[index+1:]),//右半边
	}
	return root
}
func findMax(nums []int) (index int){
	for i:=0;i<len(nums);i++{
		if nums[i]>nums[index]{
			index=i
		}
	}
	return
}*/
