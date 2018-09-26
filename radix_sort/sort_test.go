package radix_sort

import "testing"

func isLessFirst(list []int) bool {
	for i := 0; i < len(list) - 1; i++ {
		if list[i] > list[i+1] {
			return false
		}
	}
	return true
}

func Test_RadixSort(t *testing.T){
	source := []int{4, 35, 12, 16}
	target := radixSort(source)

	if !isLessFirst(target) {
		t.Error("Test_RadixSort Fail", target)
	}

	t.Log("Test_RadixSort Success")
}
