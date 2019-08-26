package main

func main() {
	nums1 := []int{0,1,0,3,12}
	moveZeroes283_1(nums1)
}

// 解法2：
func moveZeroes283_1(nums []int)  {
	last := 0
	for cur := 0; cur < len(nums); cur++ {
		if nums[cur] != 0 {
			nums[last], nums[cur] = nums[cur], nums[last]
			last++
		}
	}
}