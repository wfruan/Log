package buildTree

import "fmt"

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	m := make(map[string]int)
	var result []*TreeNode

	var dfs func(root *TreeNode) string
	dfs = func(root *TreeNode) string {
		if root == nil {
			return ""
		}
		serial := fmt.Sprintf("%d,%s,%s", root.Val, dfs(root.Left), dfs(root.Right))
		m[serial]++
		if m[serial] == 2 {
			result = append(result, root)
		}
		return serial
	}

	dfs(root)
	return result
}


