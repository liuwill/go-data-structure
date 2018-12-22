package bloom_filter

type BloomFilter struct {
	Hash []bool
	Size int
}

func newBloomFilterInstance(size int) *BloomFilter {
	// if size <= 0 {
	// 	size = 100
	// }

	return &BloomFilter{
		Hash: make([]bool, size),
		Size: size,
	}
}

func (filter *BloomFilter) insert(input string) {
	inputHash := filter.findHash(input)
	for _, pos := range inputHash {
		filter.Hash[pos] = true
	}
}

func (filter *BloomFilter) contain(input string) bool {
	inputHash := filter.findHash(input)
	for _, pos := range inputHash {
		if filter.Hash[pos] == false {
			return false
		}
	}
	return true
}

var hashList = []func(intput string, size int) int{
	func(intput string, size int) int {
		return 0
	},
	func(intput string, size int) int {
		return 0
	},
	func(intput string, size int) int {
		return 0
	},
}

func (filter *BloomFilter) findHash(input string) []int {
	hashResults := make([]int, len(hashList))
	for i, _ := range hashList {
		hashResults[i] = hashList[i](input, filter.Size)
	}

	return hashResults
}
