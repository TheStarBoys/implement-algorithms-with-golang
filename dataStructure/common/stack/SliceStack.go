package stack

import "fmt"

// 切片栈
type SliceStack struct {
	data []interface{}
	top int
}

func NewSliceStack() *SliceStack {
	return &SliceStack{
		data: make([]interface{}, 0, 32),
		top: -1,
	}
}

func (s *SliceStack) IsEmpty() bool {
	return s.top == -1
}

func (s *SliceStack) Push(i interface{}) {
	s.top++
	if s.top > len(s.data) - 1 {
		s.data = append(s.data, i)
		return
	}

	s.data[s.top] = i
}

func (s *SliceStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	v := s.data[s.top]
	s.top--

	return v
}

func (s *SliceStack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}

	return s.data[s.top]
}

func (s *SliceStack) Flush() {
	s.top = -1
}

func (s *SliceStack) String() string {
	return fmt.Sprintf("data: %v, top: %d", s.data, s.top)
}