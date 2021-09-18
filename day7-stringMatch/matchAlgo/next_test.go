package matchAlgo

import (
	"fmt"
	"testing"
)

func TestGenerateNext(t *testing.T) {
	s := "ababacd"
	next := GenerateNext(s)
	/*for i := 0; i < len(next); i++ {
		next[i] = next[i]-1
	}
	for i := 0; i < len(next)-1; i++ {
		next[i] = next[i+1]
	}*/
	fmt.Println(next)

}
