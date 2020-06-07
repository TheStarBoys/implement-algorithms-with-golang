package priorityQueue

import (
	"testing"
	"fmt"
)

func TestPriorityQueue(t *testing.T) {
	queue := NewPriorityQueue(4)
	if node := queue.Front(); node != nil {
		t.Errorf("expect: node is nil")
	}
	queue.Push(Node{3, 2})
	if node := queue.Front(); node == nil || node.priority != 2 && node.value.(int) != 3 {
		t.Errorf("node is not Node{3, 2}")
	}
	queue.Push(Node{1, 7})
	queue.Push(Node{5, 4})
	queue.Push(Node{2, 4})
	queue.Push(Node{999, 999}) // queue is full, operation failed
	for queue.Len() > 0 {
		node := queue.Pop()
		fmt.Printf("value: %d, priority: %d\n", node.value, node.priority)
	}
	// Output:
	// value: 1, priority: 7
	// value: 5, priority: 4
	// value: 2, priority: 4
	// value: 3, priority: 2
}
