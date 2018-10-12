package sorts

func findInsertPos(list []int, val int, cursor int) int {
	pos := 0
	for i := 0; i < cursor; i++ {
		if list[i] < val {
			pos++
		} else {
			break
		}
	}
	return pos
}

func findBinaryPos(list []int, val int, cursor int) int {
	pos := 0

	low, high := 0, cursor
	for low <= high {
		pos = (low + high) / 2
		if list[pos] > val {
			high = pos - 1
		} else {
			low = pos + 1
		}
	}

	if list[pos] < val && pos < cursor {
		pos++
	}
	return pos
}

func insertItem(list []int, val int, cursor int) []int {
	pos := findBinaryPos(list, val, cursor)

	for i := cursor - 1; i >= pos; i-- {
		list[i+1] = list[i]
	}
	list[pos] = val

	return list
}

func insertSort(rawList []int) []int {
	list := make([]int, len(rawList))
	cursor := 0
	for _, val := range rawList {
		list = insertItem(list, val, cursor)
		cursor++
	}
	return list
}

// Real Insertion Sort
func insertSortReal(rawList []int) []int {
	list := make([]int, len(rawList))
	copy(list, rawList)
	for i := 1; i < len(list); i++ {
		pos := i
		val := list[i]
		for j := i; j > 0; j-- {
			if val >= list[j-1] {
				break
			}
			list[j] = list[j-1]
			pos--
		}
		list[pos] = val
	}
	return list
}

// Reduce Find
func insertFindItem(list []int, val int, cursor int) []int {
	pos := cursor
	for i := cursor; i > 0; i-- {
		if val >= list[i-1] {
			break
		}
		list[i] = list[i-1]
		pos--
	}

	list[pos] = val
	return list
}

func insertSortFind(rawList []int) []int {
	list := make([]int, len(rawList))
	cursor := 0
	for _, val := range rawList {
		list = insertFindItem(list, val, cursor)
		cursor++
	}
	return list
}

// Slice And Dispatch
func insertItemSlice(list []int, val int) []int {
	pos := findInsertPos(list, val, len(list))

	if pos >= len(list) {
		list = append(list, val)
	} else {
		list = append(list[:pos+1], list[pos:]...)
		list[pos] = val
	}

	return list
}

func insertSortSlice(rawList []int) []int {
	list := []int{}
	for _, val := range rawList {
		list = insertItemSlice(list, val)
	}
	return list
}
