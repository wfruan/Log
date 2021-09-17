package search

func BinarySearch(nums []int, num int, low int, high int) int {
	// 递归终止条件
	if low > high {
		return -1
	}

	// 通过中间元素进行二分查找
	mid := (low + high) / 2
	// 递归查找
	if num > nums[mid] {
		// 如果待查找数据大于中间元素，则在右区间查找
		return BinarySearch(nums, num, mid + 1, high)
	} else if num < nums[mid] {
		// 如果待查找数据小于中间元素，则在左区间查找
		return BinarySearch(nums, num, low, mid - 1)
	} else {
		// 找到了，返回索引值
		return mid
	}
}
