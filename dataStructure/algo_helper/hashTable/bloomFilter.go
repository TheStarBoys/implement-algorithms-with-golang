package hashTable

import "math"

type BloomFilter struct {
	data []int
}

func NewBloomFilter(size int) *BloomFilter {
	return &BloomFilter{
		data: make([]int, size),
	}
}

func (bf *BloomFilter) Add(data string) {
	first := bf.hashcode1(data)
	second := bf.hashcode2(data)
	third := bf.hashcode3(data)

	bf.data[first % len(data)] = 1
	bf.data[second % len(data)] = 1
	bf.data[third % len(data)] = 1
}

func (bf *BloomFilter) MightContain(data string) bool {
	first := bf.hashcode1(data)
	second := bf.hashcode2(data)
	third := bf.hashcode3(data)

	if bf.data[first % len(data)] == 0 {
		return false
	}
	if bf.data[second % len(data)] == 0 {
		return false
	}
	if bf.data[third % len(data)] == 0 {
		return false
	}

	return true
}

func (bf *BloomFilter) hashcode1(data string) int {
	hash := 0
	for i := 0; i < len(data); i++ {
		hash = 33 * hash + int(data[i])
	}

	return int(math.Abs(float64(hash)))
}

func (bf *BloomFilter) hashcode2(data string) int {
	p := 16777619
	hash := 2166136261
	for i := 0; i < len(data); i++ {
		hash = (hash ^ int(data[i])) * p
	}
	hash += hash << 13
	hash ^= hash >> 7
	hash += hash << 3
	hash ^= hash >> 17
	hash += hash << 5
	return int(math.Abs(float64(hash)))
}

func (bf *BloomFilter) hashcode3(data string) int {
	var hash int
	for i := 0; i < len(data); i++ {
		hash += int(data[i])
		hash += hash << 10
		hash ^= hash >> 6
	}
	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15
	return int(math.Abs(float64(hash)))
}