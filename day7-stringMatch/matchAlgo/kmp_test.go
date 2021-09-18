package matchAlgo

import (
	"fmt"
	"testing"
)

func strStrV2(haystack, needle string) []int {
	var res []int
	// 子串长度=0
	if len(needle) == 0 {
		return nil
	}
	//主串长度=0，或者主串长度小于子串长度
	if len(haystack) == 0 || len(haystack) < len(needle) {
		return nil
	}
	// 子串长度=1时单独判断
	if len(needle) == 1 {
		i := 0
		for ; i < len(haystack); i++ {
			if haystack[i] == needle[0] {
				res = append(res,i)
				return res
			}
		}
		return nil
	}

	// 其他情况走 KMP 算法
	return KmpSearch(haystack, needle)
}

func TestKmpSearch(t *testing.T) {
	s := "bacbababradababacaeembabacaddababacasdsd"
	p := "ababacaee"
	pos := strStrV2(s,p)
	fmt.Printf("Find \"%s\" at %d in \"%s\"\n", p, pos, s)
}


