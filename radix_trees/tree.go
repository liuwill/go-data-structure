package radix_trees

type TreeNode struct {
	isRoot   bool
	start    byte
	current  string
	children []*TreeNode
}

type RadixTree struct {
	root *TreeNode
}

func CreateRadixTree() *RadixTree {
	tree := &RadixTree{
		root: &TreeNode{
			isRoot:   true,
			children: []*TreeNode{},
		},
	}

	return tree
}

func (tree *RadixTree) insert(word string) {
	currentNode := tree.root
	pos := 0
	isMatch := false

	for _, node := range currentNode.children {
		if node.start == word[pos] {
			isMatch = true
			currentNode = node
		}
	}

	if !isMatch {
		currentNode.children = append(currentNode.children[:], &TreeNode{
			start:    word[pos],
			current:  word[pos:],
			children: []*TreeNode{},
		})
	}
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
