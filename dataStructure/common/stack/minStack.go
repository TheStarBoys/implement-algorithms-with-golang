package stack

import (
	"math"
)

// 最小栈
type MinStack struct {
	elemStack *SliceStack // 元素栈
	minStack  *SliceStack // 最小栈，栈顶为当前最小值
}

// isMin() 需返回 a <= b 为true
func (m *MinStack) Push(data interface{}, isMin func(a, b interface{}) bool) {
	m.elemStack.Push(data)

	if m.minStack.IsEmpty() {
		m.minStack.Push(data)
	} else {
		if isMin(data, m.minStack.Top()) {
			m.minStack.Push(data)
		}
	}
}

func (m *MinStack) Pop() interface{} {
	topData := m.elemStack.Pop()

	if topData == m.Min() {
		m.minStack.Pop()
	}

	return topData
}

func (m *MinStack) Min() interface{} {
	if m.minStack.IsEmpty() {
		return math.MaxInt32
	} else {
		return m.minStack.Top()
	}
}
