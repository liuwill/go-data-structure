package trees

type BinaryNode struct {
	Left  *BinaryNode
	Right *BinaryNode
	Val   int
}

func (node *BinaryNode) insert(val int) {

}

type BinaryTree struct {
	Root *BinaryNode
}

func (tree *BinaryTree) createNode(val int) *BinaryNode {
	return &BinaryNode{
		Left:  nil,
		Right: nil,
		Val:   val,
	}
}

func (tree *BinaryTree) insert(val int) {
	if tree.Root == nil {
		tree.Root = tree.createNode(val)
	} else {
		tree.Root.insert(val)
	}
}
