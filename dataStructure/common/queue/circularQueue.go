package queue
// LeetCode官方的循环队列实现golang版
type CircularQueue struct {
	head int
	tail int
	data []interface{}
	size int
}

/** Initialize your data structure here. Set the size of the Queue to be k. */
func NewCircularQueue(k int) CircularQueue {
	return CircularQueue{-1, -1, make([]interface{}, k), k}
}


/** Insert an element into the circular Queue. Return true if the operation is successful. */
func (this *CircularQueue) EnQueue(value interface{}) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		this.head = 0
	}
	this.tail = (this.tail + 1) % this.size
	this.data[this.tail] = value
	return true
}


/** Delete an element from the circular Queue. Return true if the operation is successful. */
func (this *CircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	if this.head == this.tail {
		this.head = -1
		this.tail = -1
		return true
	}
	this.head = (this.head + 1) % this.size
	return true
}


/** Get the front item from the Queue. */
func (this *CircularQueue) Front() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.data[this.head]
}


/** Get the last item from the Queue. */
func (this *CircularQueue) Rear() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.data[this.tail]
}


/** Checks whether the circular Queue is empty or not. */
func (this *CircularQueue) IsEmpty() bool {
	return this.head == -1
}


/** Checks whether the circular Queue is full or not. */
func (this *CircularQueue) IsFull() bool {
	return ((this.tail + 1) % this.size) == this.head
}
