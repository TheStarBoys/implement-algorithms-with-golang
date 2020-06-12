package mysort

import (
	"math"
)

func shellSort(nums []int) {
	gap := 1
	// 区间取值没有定论，大概是在1/2 到 1/3之间
	for gap < len(nums) {
		gap = gap * 3 + 1 // 1 -> 4 -> 13
	}

	for gap > 0 {
		// 当 gap == 1时，就是插入排序
		// 通过gap来跨区间，插排
		for i := gap; i < len(nums); i++ { // 6
			tmp := nums[i] // 4
			j := i - gap // 2
			for j >= 0 && nums[j] > tmp {
				nums[j + gap] = nums[j]
				j -= gap
			}
			nums[j + gap] = tmp
		}
		gap = int(math.Floor(float64(gap) / 3))
	}
}
