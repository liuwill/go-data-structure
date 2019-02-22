package trees

type BinaryNode struct {
	Left  *BinaryNode
	Right *BinaryNode
	Val   int
}

func CreateBinaryNode(val int) *BinaryNode {
	return &BinaryNode{
		Left:  nil,
		Right: nil,
		Val:   val,
	}
}

func (node *BinaryNode) insert(val int) {
	if val < node.Val {
		if node.Left == nil {
			node.Left = CreateBinaryNode(val)
		} else {
			node.Left.insert(val)
		}
	} else {
		if node.Right == nil {
			node.Right = CreateBinaryNode(val)
		} else {
			node.Right.insert(val)
		}
	}
}

type BinaryTree struct {
	Root *BinaryNode
}

func InitBinaryTree() *BinaryTree {
	return &BinaryTree{
		Root: nil,
	}
}

func (tree *BinaryTree) insert(val int) {
	if tree.Root == nil {
		tree.Root = CreateBinaryNode(val)
	} else {
		tree.Root.insert(val)
	}
}

func (tree *BinaryTree) search(val int) bool {
	root := tree.Root
	for root != nil {
		if root.Val == val {
			return true
		}

		if root.Val < val {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return false
}
