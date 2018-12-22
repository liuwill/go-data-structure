package bloom_filter

type BloomFilter struct {
	hash []bool
	size int
}

func newBloomFilterInstance(size int) *BloomFilter {
	if size <= 0 {
		size = 100
	}
	return &BloomFilter{
		hash: make([]bool, size),
		size: size,
	}
}

func (filter *BloomFilter) insert(input string) {

}

func (filter *BloomFilter) contain(input string) bool {
	return false
}
