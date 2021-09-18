package matchAlgo

func BfSearch(s, p string) []int {
	var res []int
	temp := 0
	begin := 0
	i, j := 0, 0
	n, m := len(s), len(p)  // 主串、子串长度
	for i = 0; i < n; begin++ {
		// 通过 BF 算法暴力匹配子串和主串
		for j = 0; j < m; j++ {
			if i < n && s[i] == p[j] {
				// 如果子串和主串对应字符相等，逐一往后匹配
				i++
			} else {
				// 否则退出当前循环，从主串下一个字符继续开始匹配
				break
			}
		}
		if j == m {
			// 子串遍历完，表面已经找到，返回对应下标
			temp = i - j
			res = append(res,temp)
		}
		// 从下一个位置继续开始匹配
		i = begin
		i++
	}
	return res
}
