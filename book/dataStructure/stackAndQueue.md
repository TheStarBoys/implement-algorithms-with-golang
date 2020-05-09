# 栈与队列

## 栈

### 1. 用切片实现栈

```go
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
```

### 2. 最小栈

```go
// 最小栈, SliceStack 为 1. 用切片实现栈 中定义的数据结构
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
```



## 队列

### 1. 用切片实现的队列

```go
// 可以动态扩容的queue
type Queue struct {
	Data []interface{}
	size int
}
func NewQueue() *Queue {
	return &Queue{}
}
func (this *Queue) Enqueue(val interface{}) {
	this.Data = append(this.Data, val)
	this.size++
}

func (this *Queue) Dequeue() {
	this.Data = this.Data[1:]
	this.size--
}


/** Get the last item from the Queue. */
func (this *Queue) Rear() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.Data[this.size- 1]
}


/** Checks whether the circular Queue is empty or not. */
func (this *Queue) IsEmpty() bool {
	return this.size == 0
}
```

### 2. 用栈模拟队列

```go
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
```

### 3. 循环队列

```go
// LeetCode官方的循环链表实现golang版
type CircularQueue struct {
	head int
	tail int
	Data []interface{}
	size int
}

/** Initialize your data structure here. Set the size of the Queue to be k. */
func NewCircularQueue(k int) CircularQueue {
	return CircularQueue{-1, -1, make([]interface{}, k), k}
}


/** Insert an element into the circular Queue. Return true if the operation is successful. */
func (this *CircularQueue) EnQueue(value interface{}) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		this.head = 0
	}
	this.tail = (this.tail + 1) % this.size
	this.Data[this.tail] = value
	return true
}


/** Delete an element from the circular Queue. Return true if the operation is successful. */
func (this *CircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	if this.head == this.tail {
		this.head = -1
		this.tail = -1
		return true
	}
	this.head = (this.head + 1) % this.size
	return true
}


/** Get the front item from the Queue. */
func (this *CircularQueue) Front() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.Data[this.head]
}


/** Get the last item from the Queue. */
func (this *CircularQueue) Rear() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.Data[this.tail]
}


/** Checks whether the circular Queue is empty or not. */
func (this *CircularQueue) IsEmpty() bool {
	return this.head == -1
}


/** Checks whether the circular Queue is full or not. */
func (this *CircularQueue) IsFull() bool {
	return ((this.tail + 1) % this.size) == this.head
}
```

