package sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	nums := []int{4, 5, 6, 7, 8, 3, 2, 1}
	sortedNums := MergeSort(nums)
	numsTrue :=[]int{1,2,3,4,5,6,7,8}
	if !reflect.DeepEqual(sortedNums,numsTrue) {
		t.Errorf("This is false")
	}

}
