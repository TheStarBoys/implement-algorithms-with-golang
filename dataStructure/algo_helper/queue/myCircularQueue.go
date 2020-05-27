package queue
// 自己的循环链表实现
type MyCircularQueue struct {
	Head int
	Tail int
	Data []int
	Size int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func MyConstructor(k int) MyCircularQueue {
	return MyCircularQueue{-1, -1, make([]int, k), k}
}


/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		this.Head = 0
		this.Tail = 0
		this.Data[0] = value
		return true
	}
	if this.Tail < this.Size - 1 {
		this.Tail++
		this.Data[this.Tail] = value
	}else if this.Tail - this.Head < this.Size - 1 {
		this.Tail = 0
		this.Data[0] = value
	}
	return true
}


/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	if this.Head == this.Tail {
		this.Head = -1
		this.Tail = -1
	} else if this.Head + 1 < this.Size {
		this.Head++
	} else {
		this.Head = 0
	}
	return true
}


/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Data[this.Head]
}


/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Data[this.Tail]
}


/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.Head == this.Tail && this.Head == -1
}


/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.Head - this.Tail == 1 || this.Head != -1 && this.Tail - this.Head == this.Size - 1
}
