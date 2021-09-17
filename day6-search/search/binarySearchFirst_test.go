package search

import (
	"fmt"
	"testing"
)

func TestBinarySearchFirst(t *testing.T) {
	nums := []int{1, 2, 3, 3, 4, 5, 6}
	num := 3
	index := BinarySearchFirst(nums, num, 0, len(nums) - 1)
	if index != -1 {
		fmt.Printf("Find num %d first at index %d\n", num, index)
	} else {
		fmt.Printf("Num %d not exists in nums\n", num)
	}
}
