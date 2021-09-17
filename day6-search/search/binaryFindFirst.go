package search

func BinarySearchFirst(nums []int, num, low, high int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2
	if num < nums[mid] {
		return BinarySearchFirst(nums, num, low, mid - 1)
	} else if num > nums[mid] {
		return BinarySearchFirst(nums, num, mid + 1, high)
	} else {
		// 除了需要判断第一个与 num 相等的元素索引外，其他和普通二分查找逻辑一致
		if mid == 0 || nums[mid-1] != num {
			return mid
		} else {
			return BinarySearchFirst(nums, num, low, mid - 1)
		}
	}
}
