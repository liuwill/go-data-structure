package sort

func findInsertPos(list []int, val int) int {
	pos := 0
	for _, v := range list {
		if v < val {
			pos++
		} else {
			break
		}
	}
	return pos
}

func insertItem(list []int, val int) []int {
	pos := findInsertPos(list, val)
	tail := list[pos:]
	front := append(list[0:pos], val)
	list = append(front, tail...)

	return list
}

func insertSort(rawList []int) []int {
	list := make([]int, len(rawList))
	for _, val := range rawList {
		list = insertItem(list, val)
	}
	return list
}
