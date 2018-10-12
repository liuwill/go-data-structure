package sort

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

func insertItem(list []int, val int, cursor int) []int {
	pos := findInsertPos(list, val, cursor)

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
