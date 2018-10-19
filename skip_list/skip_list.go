package skip_list

import (
	"math/rand"
)

const (
	MaxLevel = 10
	cap      = 1.0 / 2.0
)

type SkipNode struct {
	next   []*SkipNode
	parent *SkipNode
	key    int
	val    int
	isRoot bool
	isNil  bool
	// level  int
}

type SkipList struct {
	keyword string
	head    *SkipNode
	level   int
	num     int
}

func InitSkipList(keyword string) *SkipList {
	list := &SkipList{
		keyword: keyword,
		level:   MaxLevel,
		head: &SkipNode{
			isRoot: true,
			next:   make([]*SkipNode, MaxLevel),
		},
	}
	tail := &SkipNode{
		isNil: true,
	}
	// println("tail", tail)

	for i := 0; i < MaxLevel; i++ {
		list.head.next[i] = tail
	}

	return list
}

func (list *SkipList) insert(key int, val int) {
	updated := make([]*SkipNode, MaxLevel)
	head := list.head
	for i := len(head.next) - 1; i >= 0; i-- {
		for head.next[i].key < key && !head.next[i].isNil {
			head = head.next[i]
		}
		updated[i] = head
	}

	// println(key, val, "==================================================", head)
	head = head.next[0]
	if head.key == key {
		head.val = val
	} else {
		level := randomLevel()
		// println("randomLevel", level)
		/* 如果需要动态增加level
		if level > list.level {
			for i := list.level - 1; i >= level; i-- {
				updated[i] = list.head
			}
			list.level = level
		}
		*/
		head = makeNode(level, key, val)

		// println(head)
		for i := level - 1; i >= 0; i-- {
			// println(i, "%%%%%%%%%%%", updated[i], updated[i].next[i], head)
			head.next[i] = updated[i].next[i]
			updated[i].next[i] = head
		}
		head.parent = updated[0]
		// fmt.Printf("  -- %v\n", head)

		list.num++
	}
	// fmt.Printf("=== %v\n", list.head.next)
}

func (list *SkipList) find(key int) int {
	head := list.head

	for i := len(head.next) - 1; i >= 0; i-- {
		for head.next[i].key < key && !head.next[i].isNil {
			head = head.next[i]
		}
	}
	head = head.next[0]
	if head.key == key {
		return head.val
	}
	return -1
}

// func (list *SkipList) delete() {}

func makeNode(level int, key int, val int) *SkipNode {
	return &SkipNode{
		key:  key,
		val:  val,
		next: make([]*SkipNode, level),
	}
}

func randomLevel() int {
	level := 1
	for level < MaxLevel && rand.Float32() < cap {
		level++
	}
	return level
}
