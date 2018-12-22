package bloom_filter

import (
	"math"
)

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
	func(input string, size int) int {
		hash := 5381

		for charIndex := 0; charIndex < len(input); charIndex += 1 {
			char := int(input[charIndex])

			hash = (hash << 5) + hash + char
		}

		return int(math.Abs(float64(hash % size)))
	},
	func(input string, size int) int {
		hash := 0

		for charIndex := 0; charIndex < len(input); charIndex += 1 {
			char := int(input[charIndex])

			hash = (hash << 5) + hash + char
			hash &= hash
			hash = int(math.Abs(float64(hash)))
		}

		return hash % size
	},
	func(input string, size int) int {
		hash := 0

		for charIndex := 0; charIndex < len(input); charIndex += 1 {
			char := int(input[charIndex])

			hash = (hash << 5) - hash
			hash += char
			hash &= hash
		}

		return int(math.Abs(float64(hash % size)))
	},
}

func (filter *BloomFilter) findHash(input string) []int {
	hashResults := make([]int, len(hashList))
	for i, _ := range hashList {
		hashResults[i] = hashList[i](input, filter.Size)
	}

	return hashResults
}
