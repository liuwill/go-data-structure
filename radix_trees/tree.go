package radix_trees

type TreeNode struct {
	start    rune
	current  string
	children []*TreeNode
}

type RadixTree struct {
	root TreeNode
}

func CreateRadixTree() *RadixTree {
	tree := &RadixTree{
		root: TreeNode{
			children: []*TreeNode{},
		},
	}

	return tree
}

func (tree *RadixTree) insert(word string) {
}

func (tree *RadixTree) find(word string) *TreeNode {
	return nil
}

func buildRadixTree(source []string) *RadixTree {
	newTree := CreateRadixTree()

	for _, val := range source {
		newTree.insert(val)
	}
	return newTree
}

func searchRadix(tree *RadixTree, word string) bool {
	node := tree.find(word)
	return node != nil
}
