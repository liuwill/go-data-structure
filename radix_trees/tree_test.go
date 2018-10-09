package radix_trees

import (
	"fmt"
	"testing"
)

func Test_RadixSearch(t *testing.T) {
	source := []string{"api", "list"}
	wordExist := "api"
	wordNotExist := ""

	tree := buildRadixTree(source)

	fmt.Printf("==> %v", tree.root.children[1].current)
	if searchRadix(tree, wordExist) {
		t.Error("Test_RadixSearch Exist Fail", wordExist)
	}

	if !searchRadix(tree, wordNotExist) {
		t.Error("Test_RadixSearch NotExist Fail", wordNotExist)
	}

	t.Log("Test_RadixSearch Radix Search Tree")
}
