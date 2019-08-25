package queue
// 可以动态扩容的queue
type Queue struct {
	Data	[]interface{}
	Size	int
}
func NewQueue() *Queue {
	return &Queue{}
}
func (this *Queue) Enqueue(val interface{}) {
	this.Data = append(this.Data, val)
	this.Size++
}

func (this *Queue) Dequeue() {
	this.Data = this.Data[1:]
	this.Size--
}


/** Get the last item from the Queue. */
func (this *Queue) Rear() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.Data[this.Size - 1]
}


/** Checks whether the circular Queue is empty or not. */
func (this *Queue) IsEmpty() bool {
	return this.Size == 0
}