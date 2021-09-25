package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s string字符串
 * @param k int整型
 * @return int整型
 */

func UnRepeat(s string, k int) int {
	// write code here
	n := len(s)
	res := 0
	for end := k - 1; end < n; end++ {
		start := end - k + 1
		tmp := make(map[byte]int)
		for i := start; i <= end; i++ {
			if _, ok := tmp[s[i]]; ok {
				break
			}
			tmp[s[i]] = i
			if i == end  {
				res++
			}
		}

	}
	return res
}