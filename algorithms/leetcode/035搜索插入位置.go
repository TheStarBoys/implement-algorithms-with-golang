package leetcode

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		if target == nums[i] {
			return i
		}else if target < nums[i] {
			return i
		}
	}
	return len(nums)
}

func searchInsert1(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l+r)/2
		if target > nums[mid] { // 去右区间找
			l = mid + 1
		} else if target < nums[mid] { // 去左区间找
			r = mid - 1
		} else { // 找到了
			return mid
		}
	}
	return l
}