package search

func BinarySearchLast(nums []int, num, low, high int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2
	if num < nums[mid] {
		return BinarySearchLast(nums, num, low, mid - 1)
	} else if num > nums[mid] {
		return BinarySearchLast(nums, num, mid + 1, high)
	} else {
		// 除了需要判断最后一个与 num 相等的元素索引外，其他和普通二分查找逻辑一致
		if mid == len(nums) - 1 || nums[mid + 1] != num {
			return mid
		} else {
			return BinarySearchLast(nums, num, mid + 1, high)
		}
	}
}
