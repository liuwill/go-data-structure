package cache

type LRUCache struct {
	recentHead *RecentNode
	recentTail *RecentNode
	hashList   []*HashNode
	capacity   int
	size       int
	count      int
}

type RecentNode struct {
	Next *RecentNode
	Prev *RecentNode
	Hash *HashNode
	Used int
}

type HashNode struct {
	Key    int
	Val    int
	Pos    int
	Next   *HashNode
	Prev   *HashNode
	Recent *RecentNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		recentHead: nil,
		recentTail: nil,
		hashList:   make([]*HashNode, capacity),
		capacity:   capacity,
		size:       0,
	}
}

func buildLruHash(lruCache *LRUCache, key int) int {
	return key % lruCache.capacity
}

func (this *LRUCache) Get(key int) int {
	hashPos := buildLruHash(this, key)
	hashHead := this.hashList[hashPos]

	val := -1
	for hashHead != nil {
		if hashHead.Key == key {
			val = hashHead.Val

			recentNode := hashHead.Recent
			// if recentNode.Prev != nil {
			// 	recentNode.Prev.Next = recentNode.Next
			// }
			if recentNode.Next != nil {
				recentNode.Next.Prev = recentNode.Prev
				if recentNode == this.recentHead {
					this.recentHead = recentNode.Next
				}
			}

			if recentNode != this.recentTail {
				this.recentTail.Next = recentNode
				recentNode.Next = nil
				recentNode.Prev = this.recentTail

				this.recentTail = recentNode
			}
			break
		}

		hashHead = hashHead.Next
	}

	return val
}

func (this *LRUCache) Put(key int, value int) {
	hashPos := buildLruHash(this, key)

	hashNode := &HashNode{
		Key:  key,
		Val:  value,
		Pos:  hashPos,
		Next: nil,
		Prev: nil,
	}

	recentNode := &RecentNode{
		Next: nil,
		Prev: nil,
	}

	recentNode.Hash = hashNode
	hashNode.Recent = recentNode

	// 淘汰
	if this.size == this.capacity {
		expireRecent := this.recentHead
		expireHash := expireRecent.Hash

		if expireRecent.Next != nil {
			expireRecent.Next.Prev = nil
		}
		this.recentHead = expireRecent.Next

		pos := expireHash.Pos
		// if expireHash.Next != nil {
		// 	expireHash.Next.Prev = expireHash.Prev
		// }
		if expireHash.Prev != nil {
			expireHash.Prev.Next = expireHash.Next
		} else {
			this.hashList[pos] = expireHash.Next
		}
	} else {
		this.size++
	}

	// 插入
	hashHead := this.hashList[hashPos]
	hashNode.Next = hashHead
	this.hashList[hashPos] = hashNode
	if hashHead != nil {
		hashHead.Prev = hashNode
	}
	// if hashHead != nil && hashHead.Next != nil {
	// 	hashHead.Next.Prev = hashNode
	// }

	if this.recentHead == nil {
		this.recentHead = recentNode
	}

	if this.recentTail == nil {
		this.recentTail = recentNode
	} else {
		recentTail := this.recentTail

		recentNode.Prev = recentTail
		this.recentTail = recentNode
		if recentTail != nil {
			recentTail.Next = recentNode
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
