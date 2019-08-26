package leetcode

func plusOne066_0(digits []int) []int {
	for i:=len(digits)-1; i>=0; i-- {
		if digits[i] < 9 {  // 当前位置不用进位，+1，然后直接返回
			digits[i]++
			return digits
		} else {            // 要进位，当前位置置0
			digits[i] = 0
		}
	}
	return append([]int{1}, digits...)
}
