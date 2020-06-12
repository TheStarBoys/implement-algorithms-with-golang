package queue
// 可以动态扩容的SliceQueue
type SliceQueue struct {
	Data []interface{}
	size int
}
func NewSliceQueue() *SliceQueue {
	return &SliceQueue{}
}
func (this *SliceQueue) Enqueue(val interface{}) {
	this.Data = append(this.Data, val)
	this.size++
}

func (this *SliceQueue) Dequeue() {
	this.Data = this.Data[1:]
	this.size--
}


/** Get the last item from the queue. */
func (this *SliceQueue) Rear() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.Data[this.size- 1]
}


/** Checks whether the circular queue is empty or not. */
func (this *SliceQueue) IsEmpty() bool {
	return this.size == 0
}