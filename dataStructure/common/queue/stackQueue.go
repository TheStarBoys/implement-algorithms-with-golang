package queue

import "github.com/TheStarBoys/implement-algorithms-with-golang/dataStructure/common/stack"

// 用栈模拟队列
// 假设用栈A与栈B模拟队列，A为插入栈，B为弹出栈
type StackQueue struct {
	aStack *stack.SliceStack
	bStack *stack.SliceStack
}

// 1, 2, 3, 4
// a: 1, 2, 3, 4
// b: 4, 3, 2, 1
func (q *StackQueue) EnQueue(data int) {
	q.aStack.Push(data)
}

func (q *StackQueue) DeQueue() int {
	if q.bStack.IsEmpty() {
		for !q.aStack.IsEmpty() {
			elem := q.aStack.Pop().(int)
			q.bStack.Push(elem)
		}
	}
	return q.bStack.Pop().(int)
}
