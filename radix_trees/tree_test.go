package radix_trees

import (
	"testing"
)

func printRadixTree(tree *RadixTree) {
	stack := []*TreeNode{
		tree.root,
	}

	top := 0
	for top < len(stack) {
		currentNode := stack[top]
		top++

		println("node:", currentNode, "=>[", currentNode.start, ":", currentNode.current, "]")
		for _, node := range currentNode.children {
			stack = append(stack[:], node)

			println("  Child: ", node, " [Start:", node.start, ", Current:", node.current, "]")
		}

		println("---------------------------------------------------------")
	}
}

func Test_RadixSearch(t *testing.T) {
	source := []string{
		"api", "list", "new", "newest", "little", "long",
		"pie", "pai", "paid",
		"old", "older", "olderer",
	}
	wordExists := []string{
		"api",
		"long",
	}
	wordNotExists := []string{
		"", "ap", "lo", "longer", "newer", "ling",
	}

	tree := buildRadixTree(source)
	// printRadixTree(tree)
	for _, isExist := range wordExists {
		if !searchRadix(tree, isExist) {
			t.Error("Test_RadixSearch Exist Fail", isExist)
		}
	}

	for _, notExist := range wordNotExists {
		if searchRadix(tree, notExist) {
			t.Error("Test_RadixSearch NotExist Fail", notExist)
		}
	}

	t.Log("Test_RadixSearch Radix Search Tree")
}
