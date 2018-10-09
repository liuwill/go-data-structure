package radix_trees

import (
	"testing"
)

func Test_RadixSearch(t *testing.T) {
	source := []string{"api", "list", "new", "newest", "little"}
	wordExist := "api"
	wordNotExist := ""

	tree := buildRadixTree(source)
	if !searchRadix(tree, wordExist) {
		t.Error("Test_RadixSearch Exist Fail", wordExist)
	}

	if searchRadix(tree, wordNotExist) {
		t.Error("Test_RadixSearch NotExist Fail", wordNotExist)
	}

	t.Log("Test_RadixSearch Radix Search Tree")
}
