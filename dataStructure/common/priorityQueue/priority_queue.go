package priorityQueue

import (
	"container/heap"
)

// Push等操作传入Node接口，使得优先级队列支持扩展不同的结点
type Node interface {
	GetValue() interface{}
	GetPriority() int
}

// QNode 队列节点
type QNode struct {
	value    interface{}
	priority int // 优先级
}

func (q QNode) GetValue() interface{} {
	return q.value
}

func (q QNode) GetPriority() int {
	return q.priority
}

type myHeapInterface interface {
	heap.Interface
	Top() interface{}
}

type reverse struct {
	myHeapInterface
}

func (r *reverse) Less(i, j int) bool {
	return r.myHeapInterface.Less(j, i)
}

func Reverse(p myHeapInterface) *reverse {
	return &reverse{
		myHeapInterface: p,
	}
}

// 实现堆接口
type myheap []Node

func (h myheap) Len() int { return len(h) }
// 如果是大于等于，那么相同优先级的，先进的将先出
// 如果是大于，那么相同优先级的，后进的先出
func (h myheap) Less(i, j int) bool { return h[i].GetPriority() >= h[j].GetPriority() }
func (h myheap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *myheap) Push(v interface{}) {
	*h = append(*h, v.(Node))
}

func (h *myheap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len() - 1], (*h)[h.Len()-1]
	return
}

func (h myheap) Top() (v interface{}) {
	return h[h.Len()-1]
}

// PriorityQueue 优先级队列
type PriorityQueue struct {
	heap myHeapInterface // 堆是一个接口
	capacity int
}

// 可以自定义堆的实现，结点Node在堆中是以数组还是链表等数据结构存储取决于你的实现
// 允许缺省实现
func NewPriorityQueue(heapInterface myHeapInterface, capacity int) PriorityQueue {
	if heapInterface == nil {
		heapInterface = new(myheap)
	}
	return PriorityQueue{
		heap: heapInterface,
		capacity: capacity,
	}
}

func (q PriorityQueue) Len() int {
	return q.heap.Len()
}

// Push 入队
func (q *PriorityQueue) Push(node Node) {
	if q.Len() >= q.capacity {
		// 队列已满
		return
	}
	heap.Push(q.heap, node)
}

// Pop 出队
func (q *PriorityQueue) Pop() Node {
	if q.Len() == 0 {
		return nil
	}
	node :=  heap.Pop(q.heap).(Node)
	return node
}

// Front 获取队首节点
func (q *PriorityQueue) Front() Node {
	if q.Len() == 0 {
		return nil
	}
	node := q.heap.Top().(Node)
	return node
}