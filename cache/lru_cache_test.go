package cache

import (
	"fmt"
	"testing"
)

func (this *LRUCache) printLru() {
	print("LRU: ")
	item := this.recentHead
	for item != nil {
		item.printNode()
		item = item.Next
	}
	println()
}

func (this *RecentNode) printNode() {
	fmt.Printf(" RN[%v](%v:%v) ", this, this.Hash.Key, this.Hash.Val)
}

func Test_LruCache(t *testing.T) {
	var target int
	cache := Constructor(2)
	cache.Put(1, 1)
	// cache.printLru()
	cache.Put(2, 2)
	// cache.printLru()

	target = cache.Get(1) // returns 1
	// cache.printLru()
	expect := 1
	if target != expect {
		t.Error("Run Test_LruCache Fail", expect, target)
	}

	cache.Put(3, 3) // evicts key 2
	// cache.printLru()
	target = cache.Get(2) // returns -1 (not found)
	expect = -1
	if target != expect {
		t.Error("Run Test_LruCache Fail", expect, target)
	}

	cache.Put(4, 4) // evicts key 1
	// cache.printLru()
	target = cache.Get(1) // returns -1 (not found)
	expect = -1
	if target != expect {
		t.Error("Run Test_LruCache Fail", expect, target)
	}

	target = cache.Get(3) // returns 3
	expect = 3
	if target != expect {
		t.Error("Run Test_LruCache Fail", expect, target)
	}

	target = cache.Get(4) // returns 4
	expect = 4
	if target != expect {
		t.Error("Run Test_LruCache Fail", expect, target)
	}
	t.Log("Run Test_LruCache Success")
}

/*
func Test_LruCache(t *testing.T) {
	callList := []string{
		"LRUCache", "Put", "Put", "Get", "Put", "Get", "Put", "Get", "Get", "Get",
	}
	sourceCase := [][]int{
		{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4},
	}
	expectCase := []int{
		nil,
		nil,
		nil,
		1,
		nil,
		-1,
		nil,
		-1,
		3,
		4,
	}

	for i, source := range sourceCase {
		expect := expectCase[i]

		tarGet := partitionLabels(source)
		if !compareList(tarGet, expect) {
			t.Error("Run Test_LruCache Fail", expect, tarGet)
		}

		tarGetFast := partitionLabelsFast(source)
		if !compareList(tarGetFast, expect) {
			t.Error("Run Test_PartitionLabelsFast Fail", expect, tarGetFast)
		}
	}

	t.Log("Run Test_LruCache Success")
}
*/
