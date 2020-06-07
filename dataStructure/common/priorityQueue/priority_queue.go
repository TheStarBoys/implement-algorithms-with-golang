package priorityQueue

import "container/heap"

// Node 队列节点
type Node struct {
	value    interface{}
	priority int // 优先级
}

// 实现堆接口
type myheap []Node

func (h myheap) Len() int { return len(h) }
// 如果是大于等于，那么相同优先级的，先进的将先出
// 如果是大于，那么相同优先级的，后进的先出
func (h myheap) Less(i, j int) bool { return h[i].priority >= h[j].priority }
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
	heap *myheap
	capacity int
}

func NewPriorityQueue(capacity int) PriorityQueue {
	return PriorityQueue{
		heap: new(myheap),
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
func (q *PriorityQueue) Pop() *Node {
	if q.Len() == 0 {
		return nil
	}
	node :=  heap.Pop(q.heap).(Node)
	return &node
}

// Front 获取队首节点
func (q *PriorityQueue) Front() *Node {
	if q.Len() == 0 {
		return nil
	}
	node := q.heap.Top().(Node)
	return &node
}