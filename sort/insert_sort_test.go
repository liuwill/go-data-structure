package sort

import (
	"sort"
	"testing"
)

func compareList(source []int, target []int) bool {
	if len(source) != len(target) {
		return false
	}

	for index := 0; index < len(source); index++ {
		if source[index] != target[index] {
			return false
		}
	}

	return true
}

func systemSort(raw []int) []int {
	target := make([]int, len(raw))
	copy(target, raw)
	sort.Ints(target)

	return target
}

func Test_InsertSort(t *testing.T) {
	rawList := []int{4, 3, 5, 21, 7, 8, 15, 1}
	expect := systemSort(rawList)

	list := insertSort(rawList)
	listSlice := insertSortSlice(rawList)
	listFind := insertSortFind(rawList)
	if !compareList(expect, list) {
		t.Error("Test_InsertSort insertSort Fail", list)
	}
	if !compareList(expect, listSlice) {
		t.Error("Test_InsertSort insertSortSlice Fail", listSlice)
	}
	if !compareList(expect, listFind) {
		t.Error("Test_InsertSort insertSortFind Fail", listFind)
	}
	t.Log("Test_InsertSort RUN TEST")
}
