package leetcode

// 解法1: 暴力
// 时间复杂度：O(k * n) 空间复杂度 O(1)
func rotate189_0(nums []int, k int)  {
	if len(nums) == 0 {
		return
	}
	for i := 0; i < k % len(nums); i++ {
		tmp := nums[len(nums) - 1]
		for j := len(nums) - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = tmp
	}
}
// 解法2: 新增一个大小为k % len(nums)大小的数组, 下面中的k默认是 k % len(nums)之后的值
// 时间复杂度：O(k + len(nums)) 120 ms , 在所有 Go 提交中击败了 93.56%
// 空间复杂度：O(k)
func rotate189_1(nums []int, k int)  {
	moveNum := k % len(nums)
	move := make([]int, 0)
	for i := len(nums) - moveNum; i < len(nums); i++ {
		move = append(move, nums[i])
	}
	for i := len(nums) - 1 - moveNum; i >= 0; i-- {
		nums[i + moveNum] = nums[i]
	}
	for i := range move {
		nums[i] = move[i]
	}
}

// 解法3：LeetCode官方新增一个数组的解法
// 时间复杂度：O(2n)
// 空间复杂度：O(n)
func rotate189_2(nums []int, k int) {
	a := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		a[(i + k) % len(nums)] = nums[i]
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = a[i]
	}
}
// 解法4：使用环状替换 这个暂时没理解到
func rotate189_3(nums []int, k int) {
	k = k % len(nums)
	count := 0
	for start := 0; count < len(nums); start++ {
		current  := start
		prev := nums[start]
		flag := 1
		for flag == 1 || start != current {
			flag = 0
			next := (current + k) % len(nums)
			temp := nums[next]
			nums[next] = prev
			prev = temp
			current = next
			count++
		}
	}
}

// 解法5：利用反转
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func rotate189_4(nums []int, k int) {
	k = k % len(nums)
	reverse189(nums, 0, len(nums) - 1)
	reverse189(nums, 0, k - 1)
	reverse189(nums, k, len(nums) - 1)
}

func reverse189(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}