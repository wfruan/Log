package search

import (
	"fmt"
	"sort"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{4, 6, 5, 3, 1, 8, 2, 7}
	sort.Ints(nums)  // 先对待排序数据序列进行排序
	fmt.Printf("Sorted nums: %v\n", nums)
	num := 5
	index := BinarySearch(nums, num, 0, len(nums)-1)
	if index != -1 {
		fmt.Printf("Find num %d at index %d\n", num, index)
	} else {
		fmt.Printf("Num %d not exists in nums\n", num)
	}
}
