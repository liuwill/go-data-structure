package skip_list

import "math/rand"

const MaxLevel = 4

type SkipNode struct {
	next   []*SkipNode
	parent *SkipNode
	key    int
	val    int
	isRoot bool
	// level  int
}

type SkipList struct {
	keyword string
	head    *SkipNode
	level   int
	num     int
}

func InitSkipList(keyword string) *SkipList {
	return &SkipList{
		keyword: keyword,
		head: &SkipNode{
			isRoot: true,
			next:   []*SkipNode{},
		},
	}
}

func (list *SkipList) insert(key int, val int) {
	updated := make([]*SkipNode, MaxLevel)
	head := list.head
	for i := len(head.next); i >= 0; i-- {
		for head.next[i].key < key {
			head = head.next[i]
		}
		updated[i] = head
	}

	head = head.next[1]
	if head.key == key {
		head.val = val
	} else {
		level := randomLevel()
		if level < list.level {
			for i := list.level + 1; i >= level; i-- {
				updated[i] = list.head
			}
			list.level = level
		}
		head = makeNode(level, key, val)

		for i := 1; i < level; i++ {
			head.next[i] = updated[i].next[i]
			updated[i].next[i] = head
		}
	}
}

func (list *SkipList) find() {

}

func (list *SkipList) delete() {

}

func makeNode(level int, key int, val int) *SkipNode {
	return &SkipNode{
		key:  key,
		val:  val,
		next: make([]*SkipNode, level),
	}
}

func randomLevel() int {
	level := 1
	for level < MaxLevel && rand.Float32() < 1/2 {
		level++
	}
	return level
}
