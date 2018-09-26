package radix_sort

func radixSort(list []int) []int {
	ll := len(list)
	bit := 1
	hasBit := true
	for hasBit {
		hasBit = false
		radixList := make([]int, 10)
		radixCountList := make([]int, 10)
		middleList := make([]int, ll)

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
			radixList[index] = start + val
			start = start + val
		}

		// fmt.Printf("%v => %v  +  %v\n", radixCountList, radixList, list)
		for j := len(list) - 1; j >= 0; j-- {
			cursor := (list[j] / bit) % 10
			pos := radixList[cursor] - 1

			radixList[cursor]--
			// println(pos, j, list[j], cursor)
			middleList[pos] = list[j]
		}

		bit = bit * 10
		list = middleList
	}
	return list
}
