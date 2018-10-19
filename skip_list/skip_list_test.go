package skip_list

import (
	"fmt"
	"testing"
)

func printSkipList(skipList *SkipList) {
	println(skipList.keyword)
	head := skipList.head

	for head != nil {
		fmt.Printf("%v ", head)
		println(head)
		if len(head.next) < 1 {
			break
		}
		head = head.next[0]
	}
}

func Test_SkipList(t *testing.T) {
	source := [][]int{
		{4, 1}, {1, 2}, {12, 3}, {16, 4},
		{5, 5}, {21, 6}, {3, 7}, {31, 8},
	}
	skipList := InitSkipList("test")

	for _, item := range source {
		skipList.insert(item[0], item[1])
	}

	printSkipList(skipList)
	t.Error("Test_SkipList Fail", skipList)
	t.Log("Test_SkipList Success", skipList)
}
