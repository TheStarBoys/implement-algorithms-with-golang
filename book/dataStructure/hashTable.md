# 散列表

## 简介

散列表用的是数组支持下标随机访问数据的特性，所以散列表其实就是数组的一种扩展，由数组演化而来。可以说，如果没有数组，就没有散列表。

**散列函数**：将给定输入转换为数组下标的**映射方法**，就叫做散列函数（或Hash函数，哈希函数），散列函数计算得到的值就是**散列值**（或Hash值，哈希值），给定的输入就是**键**（Key）。

散列表通过利用数组支持按下标随机访问时间复杂度为 `O(1)` 的特性。通过散列函数将元素的键值映射为下标，将对应数据存储到数组中对应下标的位置。当按照键值查找该元素时，用同样的散列函数，将键值转化为数组下标，从对应的数组下标位置取出数据。

**散列函数设计的基本要求：**

1. 散列函数计算得到的散列值需要是一个非负整数
2. 如果 `key 1 == key 2`， 那么 `Hash(key1) == Hash(key2)`
3. 如果 `key1 != key2`，那么 `Hash(key1) != Hash(key2)`



**散列冲突：**

由于散列函数是将键值转换为数组下标，将一个无穷大的数映射为一个有穷的数组下标，自然会出现散列冲突的现象。常用的解决散列冲突的有两类：

**1. 开放寻址法**

开放寻址法的核心思想是，如果出现了散列冲突，我们就重新探测一个空闲位置，将其插入。那如何重新探测新的位置呢？我先讲一个比较简单的探测方法，线性探测（Linear Probing）。当我们往散列表中插入数据时，如果某个数据经过散列函数散列之后，存储位置已经被占用了，我们就从当前位置开始，依次往后查找，看是否有空闲位置，直到找到为止。

**2. 链表法**

链表法是一种更加常用的散列冲突解决办法，相比开放寻址法，它要简单很多。在散列表中，每个“桶（bucket）”或者“槽（slot）”会对应一条链表，所有散列值相同的元素我们都放到相同槽位对应的链表中。 



## 哈希表实现

### 哈希集

```go
type MyHashSet struct {
	Set [][]int
}


/** Initialize your data structure here. */
func ConstructorMyHashSet() MyHashSet {
	set := make([][]int, 1000)
	return MyHashSet{set}
}


func (this *MyHashSet) Add(key int)  {
	bucket := key % 1000
	if this.Contains(key) {
		return
	}
	this.Set[bucket] = append(this.Set[bucket], key)
}


func (this *MyHashSet) Remove(key int)  {
	bucket := key % 1000 // 0
	for i, v := range this.Set[bucket] {
		if key == v {
			this.Set[bucket] = append(this.Set[bucket][:i], this.Set[bucket][i+1:]...)
			return
		}
	}
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	bucket := key % 1000
	for _, v := range this.Set[bucket] {
		if key == v {
			return true
		}
	}
	return false
}
```



### 哈希表

```go
type MyHashMap struct {
	Maps [][]pair
}
type pair struct {
	key int
	value int
}

var maxLen = 100 * 1000
/** Initialize your data structure here. */
func ConstructorMyHashMap() MyHashMap {
	maps := make([][]pair, maxLen)
	return MyHashMap{maps}
}


/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int)  {
	bucket := getIndex(key)
	if i := this.getPos(bucket, key); i == -1 {
		this.Maps[bucket] = append(this.Maps[bucket], pair{key, value})
	} else {
		this.Maps[bucket][i].value = value
	}

}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	bucket := getIndex(key)
	if i := this.getPos(bucket, key); i != -1 {
		return this.Maps[bucket][i].value
	}
	return -1
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
	bucket := getIndex(key)
	if i := this.getPos(bucket, key); i != -1 {
		this.Maps[bucket] = append(this.Maps[bucket][:i], this.Maps[bucket][i+1:]...)
	}
}

func getIndex(key int) int {
	return key % 100 * 1000
}

func (this *MyHashMap)getPos(bucket, key int) int {
	for i, v := range this.Maps[bucket] {
		if v.key == key {
			return i
		}
	}
	return -1
}
```

## 应用

### 如何判断一个元素在亿级数据中是否存在？

```go
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
```

