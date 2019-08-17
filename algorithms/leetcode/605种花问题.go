package leetcode

func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)
	count := 0
	for i := 0; i < length; i++ {
		if flowerbed[i] == 1 {
			continue
		}
		pre := i
		if pre == 0 {
			pre = 0
		}else {
			pre = flowerbed[i-1]
		}
		next := i
		if next == length - 1 {
			next = 0
		}else {
			next = flowerbed[i+1]
		}
		if pre == 0 && next == 0 {
			count++
			flowerbed[i] = 1
		}
	}
	return count >= n
}
