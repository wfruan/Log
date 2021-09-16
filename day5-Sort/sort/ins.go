package sort

func InsertionSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	for i := 0; i < len(nums); i++ {
		// 每次从未排序区间取一个数据 value
		value := nums[i]
		// 在已排序区间找到插入位置
		j := i - 1
		for ; j >= 0; j-- {
			// 如果比 value 大后移
			if nums[j] > value {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		// 插入数据 value
		nums[j+1] = value
	}

	return nums
}
