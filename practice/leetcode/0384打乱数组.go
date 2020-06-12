package leetcode

import (
	"math/rand"
)

type Solution384 struct {
	Raw []int
}


func Constructor384_0(nums []int) Solution384 {
	return Solution384{nums}
}


/** Resets the mathematics to its original configuration and return it. */
func (this *Solution384) Reset384_0() []int {
	return this.Raw
}


/** Returns a random shuffling of the mathematics. */
func (this *Solution384) Shuffle384_0() []int {
	nums := make([]int, len(this.Raw))
	aux := make([]int, len(this.Raw))
	copy(aux, this.Raw)

	for i := range nums {
		removeIndex := rand.Intn(len(aux))
		nums[i] = aux[removeIndex]
		aux = append(aux[:removeIndex], aux[removeIndex + 1:]...)
	}
	return nums
}


// 以下是错误实现

/*

type Solution struct {
    Raw []int
    Set []int
}


func Constructor(nums []int) Solution {
    return Solution{nums, []int{}}
}


func (this *Solution) Reset() []int {
	return this.Raw
}


func (this *Solution) Shuffle() []int {
	aux := make([]int, len(this.Raw))
	copy(aux, this.Raw)
	for len(this.Set) != len(this.Raw) {
		removeIndex := rand.Intn(len(aux))
		this.Set = append(this.Set, this.Raw[removeIndex])
		aux = append(aux[:removeIndex], aux[removeIndex + 1:]...)
	}
	return this.Set
}

 */