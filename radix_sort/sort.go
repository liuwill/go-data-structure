package radix_sort

func radixSort(list []int) []int {
	radixList := make([]int, 10)
	radixCountList :=  make([]int, 10)

	ll := len(list)
	bit := 1
	hasBit := true
	middleList := make([]int, ll)
	for hasBit {
		hasBit = false

		for i := 0; i < ll; i++ {
			middle := list[i] / bit
			cursor := middle % 10
			radixCountList[cursor]++

			if cursor > 0 || middle > 0 {
				hasBit = true
			}
		}

		if !hasBit {
			break
		}

		start := 0
		for index, val := range radixCountList {
			radixList[index] = start + val - 1
			start = start + val - 1
		}

		for j := len(list) - 1; j >= 0; j-- {
			cursor := (list[j] / bit) % 10
			pos := radixList[cursor]

			radixList[cursor]--
			println(pos, j, cursor)
			middleList[pos] = list[j]
		}

		bit = bit * 10
		list = middleList
	}
	return list
}
