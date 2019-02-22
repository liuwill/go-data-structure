package trees

type BinaryNode struct {
	Left  *BinaryNode
	Right *BinaryNode
	Val   int
	Level int
}

func CreateBinaryNode(val int, level int) *BinaryNode {
	return &BinaryNode{
		Left:  nil,
		Right: nil,
		Val:   val,
		Level: level,
	}
}

func (node *BinaryNode) insert(val int) {
	if val < node.Val {
		if node.Left == nil {
			node.Left = CreateBinaryNode(val, node.Level+1)
		} else {
			node.Left.insert(val)
		}
	} else {
		if node.Right == nil {
			node.Right = CreateBinaryNode(val, node.Level+1)
		} else {
			node.Right.insert(val)
		}
	}
}

type BinaryTree struct {
	Root  *BinaryNode
	Count int
}

func InitBinaryTree() *BinaryTree {
	return &BinaryTree{
		Root:  nil,
		Count: 0,
	}
}

func (tree *BinaryTree) insert(val int) {
	if tree.Root == nil {
		tree.Root = CreateBinaryNode(val, 0)
	} else {
		tree.Root.insert(val)
	}
	tree.Count++
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
