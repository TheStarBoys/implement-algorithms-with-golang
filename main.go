package main

import (
	"fmt"
	"github.com/TheStarBoys/implement-algorithms-with-golang/dataStructure/common/queue"
)
/*
5
1 2 3 3 5
3
1 2 1
2 4 5
3 5 3
 */
func main() {
	m := Constructor(1)
	fmt.Println(m.Next(-1))
	fmt.Println(m.Next(10))
	fmt.Println(m.Next(3))
	fmt.Println(m.Next(5))
}

type MovingAverage struct {
	queue.CircularQueue
	NumSize int
}


/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{queue.Constructor(size), 0}
}


func (this *MovingAverage) Next(val int) float64 {
	if this.IsFull() {
		this.DeQueue()
	}
	this.EnQueue(val)
	if this.NumSize < this.Size {
		this.NumSize++
	}
	sum := 0
	for i := 0; i < this.Size; i++ {
		sum += this.Queue[i].(int)
	}
	return float64(sum) / float64(this.NumSize)
}