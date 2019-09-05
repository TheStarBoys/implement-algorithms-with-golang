package leetcode

/*
思路分析：来源于五分钟学算法
[...] [...]		-> 二分中点在左区间，要找的元素在右区间
 m      t

[...] [...]		-> 二分中点在右区间，要找的元素在左区间
 t      m

[...] [...]		-> 二分中点和要找的元素都在左区间且m在t左边
 m t

[...] [...]		-> 二分中点和要找的元素都在左区间且t在m左边
 t m

[...] [...]		-> 二分中点和要找的元素都在右区间且m在t左边
       m t

[...] [...]		-> 二分中点和要找的元素都在右区间且t在m左边
       t m

合并一些情形：
[...] [...]		-> target在左区间，mid在右区间，或者mid在左区间但是比target大，移动尾指针
 t m  m  m

[...] [...]		-> target和mid都在左区间，但target比mid大，移动首指针
 m t

[...] [...]		-> target在右区间，mid在左区间，或者mid在右区间但比target小，移动首指针
 m m  m  t

[...] [...]		-> target和mid都在右区间，但target比mid小，移动尾指针
      t m
 */
func search033_0(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l, h := 0, len(nums) - 1
	for l + 1 < h {
		mid := l + (h-l)/2
		if target > nums[h] { // target 在左区间
			if nums[mid] > target || nums[mid] <= nums[h] {
				// mid在左区间且比target大 || mid在右区间
				h = mid
			}else if nums[mid] == target {
				return mid
			}else {
				l = mid
			}
		}else {
			if nums[mid] > nums[h] || nums[mid] < target {
				l = mid
			}else if nums[mid] == target {
				return mid
			}else {
				h = mid
			}
		}
	}
	if nums[l] == target {
		return l
	}else if nums[h] == target {
		return h
	}
	return -1
}

// 思想与上面的类似，但代码量更少
func search033_1(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	last := nums[len(nums) - 1]
	l, h := 0, len(nums) - 1

	for l + 1 < h {
		mid := l + (h - l) / 2
		if nums[mid] == target {
			return mid
		}else if nums[mid] > last && (nums[mid] < target || target <= last) ||
			nums[mid] < last && target <= last && target > nums[mid] {
			l = mid
		}else {
			h = mid
		}
	}
	if nums[l] == target {
		return l
	}else if nums[h] == target {
		return h
	}
	return -1
}
