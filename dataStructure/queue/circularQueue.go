package queue
// LeetCode官方的循环链表实现golang版
type CircularQueue struct {
	Head    int
	Tail    int
	Queue   []int
	Size 	int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) CircularQueue {
	return CircularQueue{-1, -1, make([]int, k), k}
}


/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *CircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		this.Head = 0
	}
	this.Tail = (this.Tail + 1) % this.Size
	this.Queue[this.Tail] = value
	return true
}


/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *CircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	if this.Head == this.Tail {
		this.Head = -1
		this.Tail = -1
		return true
	}
	this.Head = (this.Head + 1) % this.Size
	return true
}


/** Get the front item from the queue. */
func (this *CircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Queue[this.Head]
}


/** Get the last item from the queue. */
func (this *CircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Queue[this.Tail]
}


/** Checks whether the circular queue is empty or not. */
func (this *CircularQueue) IsEmpty() bool {
	return this.Head == -1
}


/** Checks whether the circular queue is full or not. */
func (this *CircularQueue) IsFull() bool {
	return ((this.Tail + 1) % this.Size) == this.Head
}
