package leetcode

import "sort"





// todo
func findKthLargest0(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	return nums[k-1]
}

func findKthLargest1(nums []int, k int) int {
	k = len(nums) - k
	l, h := 0, len(nums)-1
	for l < h {
		j := partition(nums, l, h)
		if j == k {
			break
		} else if j < k{
			l = j + 1
		} else {
			h = j - 1
		}
	}
	return nums[k]
}

func partition(a []int, l, h int) int {
	i, j := l+1, h
	for {
		for a[i] < a[l] && i < h {
			i++
		}
		for a[j] > a[l] && j > l {
			j--
		}
		if i >= j {
			break
		}
		swap(a, l, j)
	}
	swap(a, l, j)
	return j
}

//func FindKthLagest(a []int, k int) int{
//	return findKthLargest(a, k)
//}