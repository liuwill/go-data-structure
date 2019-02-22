package trees

import (
	"testing"
)

func Test_BinarySearchTree(t *testing.T) {
	source := [][]int{
		{2, 5, 6, 4, 1, 9, 3},
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
