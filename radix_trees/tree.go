package radix_trees

type TreeNode struct {
	start    string
	current  string
	content  string
	children []*TreeNode
}

type RadixTree struct {
	root *TreeNode
}

func CreateRadixTree() *RadixTree {
	defer func() {
		println("CREATE EMPTY RADIX TREE")
	}()

	tree := &RadixTree{
		root: &TreeNode{
			children: []*TreeNode{},
		},
	}

	return tree
}

func (tree *RadixTree) insert(word string) {
	stack := []*TreeNode{
		tree.root,
	}

	pos := 0
	isMatch := false
	ll := len(word)

	top := 0
	currentNode := stack[top]

	// println("INPUT:", word)
	for top < len(stack) {
		currentNode = stack[top]
		top++

		// fmt.Printf("%v \n", currentNode)
		for index, node := range currentNode.children {
			if node.start != string(word[pos]) {
				continue
			}

			currentMatch := node.current

			inner := 0
			innerFull := true
			for i := 0; i < len(currentMatch); i++ {
				if pos+i > ll-1 || word[pos+i] != currentMatch[i] {
					innerFull = false
					break
				}
				inner++
			}

			if inner > 0 && innerFull {
				// 如果完全匹配当前节点，需要继续匹配子节点
				// println(node.start, node.current, "@@@@@@@")
				currentNode = node
				pos += inner

				stack = append(stack[:], node)
				break
			}

			// if inner > 0 {
			// 如果需要分裂当前节点的处理
			node.start = string(currentMatch[inner])
			node.current = currentMatch[inner:]
			newNode := &TreeNode{
				start:   string(word[pos]),
				current: word[pos : pos+inner],
				children: []*TreeNode{
					node,
				},
			}

			if pos+inner < ll {
				newNode.children = append(newNode.children, &TreeNode{
					start:    string(word[pos+inner]),
					current:  word[pos+inner:],
					content:  word,
					children: []*TreeNode{},
				})
			} else {
				newNode.content = word
			}
			currentNode.children[index] = newNode
			// }

			isMatch = true
			break
		}
	}

	if isMatch {
		return
	}

	matchPattern := word[pos:]
	if len(matchPattern) > 0 {
		currentNode.children = append(currentNode.children[:], &TreeNode{
			start:    string(word[pos]),
			current:  matchPattern,
			content:  word,
			children: []*TreeNode{},
		})
	}
	// else { // 不支持重复插入节点
	// 	currentNode.content = word
	// }
}

func (tree *RadixTree) find(word string) *TreeNode {
	stack := []*TreeNode{
		tree.root,
	}

	pos := 0
	top := 0
	for top < len(stack) {
		currentNode := stack[top]
		top++

		// fmt.Printf("i=%v \n", currentNode)
		// println(word, pos, currentNode.current)

		index := 0
		for i := 0; i < len(currentNode.current); i++ {
			if i+pos >= len(word) {
				break
			}

			if word[pos+i] != currentNode.current[i] {
				break
			}
			index++
		}

		if index > 0 && index == len(currentNode.current) && pos+index == len(word) && len(currentNode.content) > 0 {
			return currentNode
		}

		pos += index
		for _, node := range currentNode.children {
			if pos < len(word) && string(word[pos]) == node.start {
				stack = append(stack[:], node)
			}
		}
	}
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

	// fmt.Printf("%v < == > %v\n", word, node)
	return node != nil
}
