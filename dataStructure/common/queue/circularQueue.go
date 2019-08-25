package queue
// LeetCode官方的循环链表实现golang版
type CircularQueue struct {
	Head int
	Tail int
	Data []interface{}
	Size int
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
		this.Head = 0
	}
	this.Tail = (this.Tail + 1) % this.Size
	this.Data[this.Tail] = value
	return true
}


/** Delete an element from the circular Queue. Return true if the operation is successful. */
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


/** Get the front item from the Queue. */
func (this *CircularQueue) Front() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.Data[this.Head]
}


/** Get the last item from the Queue. */
func (this *CircularQueue) Rear() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.Data[this.Tail]
}


/** Checks whether the circular Queue is empty or not. */
func (this *CircularQueue) IsEmpty() bool {
	return this.Head == -1
}


/** Checks whether the circular Queue is full or not. */
func (this *CircularQueue) IsFull() bool {
	return ((this.Tail + 1) % this.Size) == this.Head
}
