package search

import (
	"fmt"
	"testing"
)

func TestBinarySearchLast(t *testing.T) {
	nums := []int{1, 2, 3, 3, 4, 5, 6}
	num := 3
	index := BinarySearchLast(nums, num, 0, len(nums) - 1)
	if index != -1 {
		fmt.Printf("Find num %d last at index %d\n", num, index)
	} else {
		fmt.Printf("Num %d not exists in nums\n", num)
	}
}
