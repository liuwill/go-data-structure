package skip_list

import (
	"fmt"
	"math/rand"
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
	expect := []int{8, 9}
	unexpected := []int{100, 1}
	source := [][]int{
		{4, 1}, {1, 2}, {12, 3}, {16, 4},
		{5, 5}, {21, 6}, {3, 7}, {31, 8},
		expect,
	}
	skipList := InitSkipList("test")

	for _, item := range source {
		skipList.insert(item[0], item[1])
	}

	for _, item := range source {
		if item[1] != skipList.find(item[0]) {
			t.Error("Test_SkipList Fail", item)
			printSkipList(skipList)
		}
	}

	expect[1] = rand.Int()
	skipList.insert(expect[0], expect[1])

	if expect[1] != skipList.find(expect[0]) {
		t.Error("Test_SkipList Fail", expect)
	}

	if skipList.find(unexpected[0]) != -1 {
		t.Error("Test_SkipList Fail", unexpected)
	}
	t.Log("Test_SkipList Success", skipList)
}
