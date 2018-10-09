package radix_trees

type TreeNode struct {
	isRoot   bool
	start    string
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
				// println(node.start, node.current, "@@@@@@@")
				currentNode = node
				pos = pos + inner

				stack = append(stack[:], node)
				break
			} else if inner > 0 {
				node.start = string(currentMatch[inner])
				node.current = currentMatch[inner:]
				newNode := &TreeNode{
					start:   string(word[pos]),
					current: word[pos : pos+inner],
					children: []*TreeNode{
						node,
					},
				}

				if pos+inner < ll-1 {
					newNode.children = append(newNode.children, &TreeNode{
						start:    string(word[pos+inner]),
						current:  word[pos+inner:],
						children: []*TreeNode{},
					})
				}
				currentNode.children[index] = newNode
			}

			isMatch = true
			break
		}
	}

	if !isMatch {
		currentNode.children = append(currentNode.children[:], &TreeNode{
			start:    string(word[pos]),
			current:  word[pos:],
			children: []*TreeNode{},
		})
	}
}

/*
func (tree *RadixTree) insertTest(word string) {
	currentNode := tree.root
	pos := 0
	isMatch := false
	ll := len(word)
	// lastNode := tree.root
	println(word, "+++")

	for index, node := range currentNode.children {
		fmt.Printf("%v\n", currentNode)
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
			println(node.start, node.current, "@@@@@@@")
			currentNode = node
			pos = pos + inner
			continue
		} else if inner > 0 {
			node.start = string(currentMatch[inner])
			node.current = currentMatch[inner:]
			newNode := &TreeNode{
				start:   string(word[pos]),
				current: word[pos : pos+inner],
				children: []*TreeNode{
					node,
				},
			}

			if pos+inner < ll-1 {
				newNode.children = append(newNode.children, &TreeNode{
					start:    string(word[pos+inner]),
					current:  word[pos+inner:],
					children: []*TreeNode{},
				})
			}
			currentNode.children[index] = newNode

			isMatch = true
			break
		}
	}

	if !isMatch {
		currentNode.children = append(currentNode.children[:], &TreeNode{
			start:    string(word[pos]),
			current:  word[pos:],
			children: []*TreeNode{},
		})
	}
}
*/

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
		for _, node := range currentNode.children {
			// stack = append(stack[:], node)
			// println(node.start)
			if pos < len(word) && string(word[pos]) == node.start {
				stack = append(stack[:], node)
			}
		}

		if currentNode.isRoot {
			pos++
			continue
		}
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

		if index > 0 && index == len(currentNode.current) && pos+index == len(word) {
			return currentNode
		}
		pos = pos + index
	}
	return nil

	/*
		var result *TreeNode
		result = nil

		pos := 0
		currentNode := tree.root
		for _, node := range currentNode.children {
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

			if index > 0 && index == len(currentNode.current) && pos+index == len(word) {
				result = currentNode
				break
			}

			if pos < len(word) && string(word[pos]) == node.start {
				currentNode = node
				println("****", string(word[pos]), pos)
			}
			println(pos, word, index, node.start)
		}

		fmt.Printf("====== %v", result)
		return result
	*/
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
