package bloom_filter

import (
	"testing"
)

func Test_BloomFilter(t *testing.T) {
	filter := newBloomFilterInstance(1000)
	source := []string{
		"217d72eccd3f1f936c4767b5ace2a5dbf2f859f79cbe2a76728428022f883c94",
		"Kihnir rejodu pakud suijber rori epupala tuf vithevupe kukeomi neveso ah cojawul fagek.",
		"wooriken",
		"Manuel Ferguson",
		"Oryx",
		"Cendant Corp",
	}
	notExists := []string{
		"Sheep",
		"Freeport McMoran Copper & Gold Inc.",
	}

	for _, val := range source {
		filter.insert(val)
	}

	for _, val := range source {
		if filter.contain(val) == false {
			t.Error("Test_BloomFilter Contain Fail", val)
		}
	}

	for _, val := range notExists {
		if filter.contain(val) == true {
			t.Error("Test_BloomFilter Not Contain Fail", val)
		}
	}

	t.Log("Test_BloomFilter Success")
}
