package leetcode

import "math"

// 找零的时候，优先用大额钞票找零
func lemonadeChange860_0(bills []int) bool {
	// 5, 10, 20
	moneys := make([]int, 3)
	for _, amount := range bills {
		if amount == 5 { // 不需要找零，直接收钱
			moneys[0] += 5
		} else {
			find := amount - 5 // 需要找零的钱
			indx := len(moneys)-1 // 2
			for find > 0 && indx >= 0 {
				tmp := int(math.Pow(float64(2), float64(indx))) * 5
				for moneys[indx] > 0 && find - tmp >= 0 {
					find -= tmp
					// 5*1, 5*2, 5*4
					moneys[indx] -= tmp
				}
				indx--
			}
			if find > 0 {
				return false
			}
			if amount == 10 {
				moneys[1] += 10
			} else {
				moneys[2] += 20
			}
		}
	}
	return true
}
