package hashTable

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

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
