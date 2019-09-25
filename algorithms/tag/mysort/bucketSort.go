package mysort

import "fmt"

// 桶排序，需要明确输入数组的取值范围
// 时间O(M+N) 空间O(M) M是桶的长度, 如果输出需要是切片的话，空间为O(M+N)
// 输入在0~1000之间
func bucketSort(nums []int) {
	bucket := make([]int, 1001)
	for _, v := range nums {
		bucket[v]++
	}
	for i := range bucket {
		for j := 0; j < bucket[i]; j++ {
			fmt.Printf("%d ", i)
		}
	}
}
