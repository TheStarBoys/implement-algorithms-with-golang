package queue

/*
MovingAverage m = new MovingAverage(3);
m.next(1) = 1
m.next(10) = (1 + 10) / 2
m.next(3) = (1 + 10 + 3) / 3
m.next(5) = (10 + 3 + 5) / 3
 */
type MovingAverage struct {
	CircularQueue
	NumSize int
}


/** Initialize your data structure here. */
func NewMovingAverage(size int) MovingAverage {
	return MovingAverage{Constructor(size), 0}
}


func (this *MovingAverage) Next(val int) float64 {
	if this.IsFull() {
		this.DeQueue()
	}
	this.EnQueue(val)
	if this.NumSize < this.size {
		this.NumSize++
	}
	sum := 0
	for i := 0; i < this.size; i++ {
		sum += this.data[i]
	}
	return float64(sum) / float64(this.NumSize)
}
