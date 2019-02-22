package trees

import (
	"fmt"
	"testing"
)

func printBinarySearchTree(tree *BinaryTree) {
	stack := make([]*BinaryNode, tree.Count)
	left := 0
	total := 1
	stack[0] = tree.Root

	result := make([]BinaryNode, tree.Count)

	for left < total {
		current := stack[left]
		result[left] = *current

		pos := left
		if current.Left != nil {
			total++
			pos++
			stack[pos] = current.Left
		}
		if current.Right != nil {
			pos++
			stack[pos] = current.Right
		}

		left++
	}

	fmt.Printf("%v => %v\n", tree.Count, result)
}

func Test_BinarySearchTree(t *testing.T) {
	source := [][]int{
		{4, 5, 6, 2, 1, 9, 3},
	}

	for _, list := range source {
		tree := InitBinaryTree()
		for _, item := range list {
			tree.insert(item)
		}

		for _, item := range list {
			if !tree.search(item) {
				t.Error("Test_BinarySearchTree Fail", item, list)
			}
		}
	}

	t.Log("Test_BinarySearchTree Success")
}
