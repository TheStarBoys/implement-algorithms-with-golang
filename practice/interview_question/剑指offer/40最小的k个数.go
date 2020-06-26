package offer

import "container/heap"

// 冒泡排序变形
func getLeastNumbers_1(arr []int, k int) []int {
	for i := 0; i < k; i++ {
		for j := 1; j < len(arr); j++ {
			if arr[j-1] < arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
	return arr[len(arr)-k:]
}

// 大顶堆求topK问题
func getLeastNumbers_2(arr []int, k int) []int {
	if k == 0 { return nil }
	h := new(myHeap)
	for i := 0; i < k; i++ {
		heap.Push(h, arr[i])
	}
	for i := k; i < len(arr); i++ {
		if h.Top().(int) > arr[i] {
			heap.Pop(h)
			heap.Push(h, arr[i])
		}
	}

	return []int(*h)
}

type myHeap []int

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len() - 1], (*h)[h.Len() - 1]
	return
}

func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *myHeap) Top() (v interface{}) {
	if h.Len() == 0 {
		return nil
	}
	return (*h)[0]
}