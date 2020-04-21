package leetcode

func canPlaceFlowers605_0(flowerbed []int, n int) bool {
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

// 注意边界条件
func canPlaceFlowers605_1(flowerbed []int, n int) bool {
	if n <= 0 { return true }
	pre, mid, next := 0, 1, 2
	count := 0
	if len(flowerbed) == 1 { return flowerbed[pre] == 0 }
	if len(flowerbed) == 2 {
		if flowerbed[pre] == 0 && flowerbed[mid] == 0 {
			count++
		}
	}
	for next < len(flowerbed) {
		if pre == 0 && flowerbed[pre] == flowerbed[mid] &&
			flowerbed[mid] == 0 {
			flowerbed[pre] = 1
			count++
		}
		if next == len(flowerbed)-1 && flowerbed[next] == 0 &&
			flowerbed[next] == flowerbed[mid] {
			flowerbed[next] = 1
			count++
		}
		if flowerbed[pre] == flowerbed[mid] && flowerbed[mid] == flowerbed[next] &&
			flowerbed[mid] == 0 {
			flowerbed[mid] = 1
			count++
		}
		pre++
		mid++
		next++
	}
	return count >= n
}