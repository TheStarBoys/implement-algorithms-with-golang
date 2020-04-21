package mysort

import "math"

// 归并排序
// 传入0和最大下标

// 将数组每次对半分，递归的排序对半分后的数组，再将对半分后的数组合并为一个，p == r 时，即数组元素个数为1 不可再分时，停止递归
// [3,4,1,2,6,5]
// [3,4,1]         [2,6,5]
// {[3,4], [1]}    {[2,6], [5]}
// {{[3], [4]}, {[1]}}  {{[2], [6]}, {[5]}}
//
func mergeSort(A []int, p, r int) {
	if p < r {
		q := int(math.Floor(float64((p + r) / 2)))
		mergeSort(A, p, q)
		mergeSort(A, q + 1, r)
		merge(A, p, q, r)
	}
}
func merge(A []int, p, q, r int) {
	n1 := q - p + 1
	n2 := r - q
	left, right := make([]int, n1 + 1), make([]int, n2 + 1)
	copy(left, A[p:q + 1])
	copy(right, A[q + 1:r + 1])
	left[n1] = math.MaxInt64
	right[n2] = math.MaxInt64
	i, j := 0, 0
	for k := p; k <= r; k++ {
		if left[i] <= right[j] {
			A[k] = left[i]
			i++
		} else {
			A[k] = right[j]
			j++
		}
	}
}
