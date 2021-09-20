package buildTree

import (
	"fmt"
	"testing"
)

func TestPreMidBuildTree(t *testing.T) {
	Pre := []int{3,9,20,15,7}
	Mid := []int{9,3,15,20,7}
	root := PreMidBuildTree(Pre,Mid)
	fmt.Println(root)


}
