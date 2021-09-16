package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	//use the slice to store the int to test
	nums := []int{4,5,6,7,8,3,2,1,}
	nums = BubbleSort(nums)
	numsTest := []int{1,2,3,4,5,6,7,8}
	if !reflect.DeepEqual(numsTest,nums) {
		t.Errorf("sort error")
	}
	fmt.Println(nums)
}
