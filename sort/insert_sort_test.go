package sort

import (
	"fmt"
	"testing"
)

func Test_InsertSort(t *testing.T) {
	rawList := []int{4, 3, 5, 21, 7, 8, 15, 1}

	list := insertSort(rawList)
	fmt.Printf("%v\n", list)
}
