package matchAlgo

import (
	"fmt"
	"testing"
)

func strStrV1(haystack, needle string) []int {
	// 子串长度=0
	if len(needle) == 0 {
		return nil
	}
	//主串长度=0，或者主串长度小于子串长度
	if len(haystack) == 0 || len(haystack) < len(needle) {
		return nil
	}
	// 调用 BF 算法查找子串
	return BfSearch(haystack, needle)
}

func TestBfSearch2(t *testing.T) {
	s := "bacbababadababacambabacaddababacasdsd"
	p := "ababaca"
	pos := strStrV1(s,p)
	fmt.Printf("Find \"%s\" at %d in \"%s\"\n", p, pos, s)
}

/*func TestBfSearch(t *testing.T) {
	s := "bacbababadababacambabacaddababacasdsd"
	p := "ababaca"
	pos := BfSearch(s,p)
	pTrue := 10 //正确结果
	switch pos {
	case 0:fmt.Println("模式串长度为0，无法匹配")
	case -1:fmt.Println("输入有误，主串为空 或者模式串比主串长")
	case -2:fmt.Println("查无此串")
	default:
		if pos!=pTrue {
			t.Errorf("测试结果与预测不符，程序有误")
		}
	}
	fmt.Printf("Find \"%s\" at %d in \"%s\"\n", p, pos, s)
}*/
