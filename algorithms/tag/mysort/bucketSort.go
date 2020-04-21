package mysort

import "fmt"

// 桶排序，需要明确输入数组的取值范围
// 时间O(M+N) 空间O(M) M是桶的长度, 如果输出需要是切片的话，空间为O(M+N)
// 输入在0~1000之间

// 桶排序需要知道一个数的取值范围，再根据取值范围生成相应数量的桶
// 遍历nums，每取出一个数在对应的桶中放一个旗子
// 最后遍历桶，将得到正确排序
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
