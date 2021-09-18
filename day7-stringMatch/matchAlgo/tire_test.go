package matchAlgo

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()
	words := []string{"Hello", "rwf", "Trie", "Go"}

	for _, word := range words{
		trie.Insert(word)
	}

	//termFind 待测找字符串

	termFind := "rwf"
	if trie.Find(termFind) {
		fmt.Printf("包含单词\"%s\"\n", termFind)
	} else {
		fmt.Printf("不包含单词\"%s\"\n", termFind)
	}

}
