package heap

import (
	"testing"
	"fmt"
	sdHeap "container/heap"
	"math"
)

type Integer int

func (i Integer) Less(j Elem) bool {
	return i < j.(Integer)
}

func TestHeap(t *testing.T) {
	heap := NewHeap(3)
	fmt.Println("------Insert------")
	heap.Insert(Integer(3))
	fmt.Println(heap.data[:heap.count+1])
	heap.Insert(Integer(2))
	fmt.Println(heap.data[:heap.count+1])
	heap.Insert(Integer(5))
	fmt.Println(heap.data[:heap.count+1])
	heap.Insert(Integer(1))
	fmt.Println(heap.data[:heap.count+1])
	fmt.Println("------RemoveMax------")
	heap.RemoveMax()
	fmt.Println(heap.data[:heap.count+1])
	heap.RemoveMax()
	fmt.Println(heap.data[:heap.count+1])
	heap.RemoveMax()
	fmt.Println(heap.data[:heap.count+1])
	heap.RemoveMax()
	fmt.Println(heap.data[:heap.count+1])

	// 将一个普通的数组建堆
	fmt.Println("------BuildHeap------")
	data := []Elem{Integer(2), Integer(3), Integer(5)}
	data = append([]Elem{nil}, data...)
	BuildHeap(data, len(data)-1)
	fmt.Println(data)
	fmt.Println("------Sort------")
	sort(data, len(data)-1)
	fmt.Println(data)
}

type reverseInteger int

func (i reverseInteger) Less(j Elem) bool {
	return i > j.(reverseInteger)
}

// 通过改变Less，将heap变成小顶堆
func TestReverseHeap(t *testing.T) {
	heap := NewHeap(3)
	fmt.Println("------Insert------")
	heap.Insert(reverseInteger(3))
	fmt.Println(heap.data[:heap.count+1])
	heap.Insert(reverseInteger(2))
	fmt.Println(heap.data[:heap.count+1])
	heap.Insert(reverseInteger(5))
	fmt.Println(heap.data[:heap.count+1])
	heap.Insert(reverseInteger(1))
	fmt.Println(heap.data[:heap.count+1])
	fmt.Println("------RemoveMax------")
	heap.RemoveMax()
	fmt.Println(heap.data[:heap.count+1])
	heap.RemoveMax()
	fmt.Println(heap.data[:heap.count+1])
	heap.RemoveMax()
	fmt.Println(heap.data[:heap.count+1])
	heap.RemoveMax()
	fmt.Println(heap.data[:heap.count+1])

	// 将一个普通的数组建堆
	fmt.Println("------BuildHeap------")
	data := []Elem{reverseInteger(2), reverseInteger(3), reverseInteger(5)}
	data = append([]Elem{nil}, data...)
	BuildHeap(data, len(data)-1)
	fmt.Println(data)
	fmt.Println("------Sort------")
	sort(data, len(data)-1)
	fmt.Println(data)
}


// 调标准库的Heap

// 先实现一个Heap需要的接口
type myHeap []int

// 通过Less的不同，实现大顶堆和小顶堆的效果
// 此实现是小顶堆
func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
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

// 仿sort包里的reverse用法，可以将具体的实现进行转换
// 比如myHeap是实现的小顶堆，通过Reverse()可以变成大顶堆
type reverseInterface interface {
	sdHeap.Interface
	Top() interface{}
}
type reverse struct {
	reverseInterface
}

func (r *reverse) Less(i, j int) bool {
	return r.reverseInterface.Less(j, i)
}

func Reverse(p reverseInterface) *reverse {
	return &reverse{
		reverseInterface: p,
	}
}

// 应用1：堆实现优先级队列的应用
func TestPriorityQueue(t *testing.T) {
	// 优先级队列，优先级高的先出队
	h := Reverse(new(myHeap))
	sdHeap.Push(h, 2) // [2]
	sdHeap.Push(h, 3) // [3 2]
	sdHeap.Push(h, 5) // [5 2 3]
	sdHeap.Push(h, 1) // [5 2 3 1]
	sdHeap.Remove(h, 0) // [3 2 1]
	for h.Len() > 0 {
		fmt.Printf("%d ", sdHeap.Pop(h))
	}
	// Output:
	// 3 2 1
}

// 应用2：Top K问题，找到一个数据集合中，前K大的数据
func TestTopK(t *testing.T) {
	fmt.Println(topK([]int{-1, -3, -2, 3}, 2)) // [-1 3]
}

// 如果是一个静态数据，直接调用即可
// 如果是一个动态数据，那么我们需要一直持有一个大小为K的堆
// 每次的数据添加，我们都跟堆顶进行比较，考虑是否将堆顶元素移除，数组元素插入堆
// 这样，堆中的K个数据就是实时的前K大的数据
func topK(data []int, k int) []int {
	h := new(myHeap)
	// 申请一个大小为 k 的堆，并且往里填充该元素的最小值
	for i := 0; i < k; i++ {
		h.Push(math.MinInt64)
	}
	// 遍历数组
	for _, v := range data {
		// 如果数组当前元素比堆顶元素大
		if v > (*h)[0] {
			// 移除堆顶元素
			sdHeap.Pop(h)
			// 将数组元素插入堆中
			sdHeap.Push(h, v)
		}
	}
	return []int(*h)
}

// 应用3：利用堆求中位数
func TestFindMedian(t *testing.T) {
	data := []int{1, 2, 3}
	fmt.Println(findMedian(data))
	// Output:
	// 2
	data = []int{1, 2, 3, 4}
	fmt.Println(findMedian(data))
	// Output:
	// 2.5
}

func findMedian(data []int) float64 {
	finder := Constructor()
	for _, v := range data {
		finder.AddNum(v)
	}

	return finder.FindMedian()
}

type MedianFinder struct {
	maxHeap *myHeap
	minHeap reverseInterface
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		maxHeap: new(myHeap),
		minHeap: Reverse(new(myHeap)),
	}
}

func (this *MedianFinder) AddNum(num int)  {
	sdHeap.Push(this.maxHeap, num) // Add to max heap

	sdHeap.Push(this.minHeap, sdHeap.Pop(this.maxHeap)) // balancing step

	// maintain size property
	if this.maxHeap.Len() < this.minHeap.Len() {
		sdHeap.Push(this.maxHeap, sdHeap.Pop(this.minHeap))
	}
}


func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeap.Len() == this.minHeap.Len() {
		v1 := this.maxHeap.Top().(int)
		v2 := this.minHeap.Top().(int)
		return float64(v1 + v2) / 2
	}
	return float64(this.maxHeap.Top().(int))
}