package leetcode

func checkPossibility(nums []int) bool {
	l, r := isIncrease(nums)
	if l != -1 { // 有错
		tmp := make([]int, len(nums))
		copy(tmp, nums) // append会修改切片的内容，所以需要copy一个新的切片
		a, _ := isIncrease(append(tmp[:l], tmp[r:]...))
		b, _ := isIncrease(append(nums[:r], nums[r+1:]...))
		return a == -1 || b == -1
	}
	return true
}

func isIncrease(nums []int) (l, r int) {
	length := len(nums)
	for i := 0; i < length -1; i++ {
		if nums[i] > nums[i+1] {
			return i, i+1
		}
	}
	return -1, -1
}


// 这个不可行
func checkPossibility0(nums []int) bool {
	count := 0
	length := len(nums)
	for i := 0; count <= 1 && i < length -1; i++ {
		if nums[i] > nums[i+1] {
			count++
		}
	}
	return count <= 1
}


