package stack

// 切片栈
type SliceStack struct {
	Data []interface{}
	size int
}

func NewSliceStack() *SliceStack {
	return &SliceStack{
		Data: make([]interface{}, 0),
		size: 0,
	}
}

func (s *SliceStack) Push(i interface{}) {
	s.Data = append(s.Data, i)
	s.size++
}

func (s *SliceStack) Pop() (res interface{}) {
	if s.IsEmpty() {
		return 0
	}

	defer func(i interface{}) {
		res = i
	}(s.Data[s.size-1])

	s.Data = s.Data[:s.size-1]
	s.size--

	return
}

func (s *SliceStack) Top() interface{} {
	if s.IsEmpty() {
		return 0
	}

	return s.Data[s.size-1]
}

func (s *SliceStack) IsEmpty() bool {
	return s.size == 0
}