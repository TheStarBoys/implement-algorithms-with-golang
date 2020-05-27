package hashTable

// 仅供参考
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


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
