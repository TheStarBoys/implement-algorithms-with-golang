package stack

import "fmt"

type LinkedListStack struct {
	topNode *node // 栈顶节点
}

type node struct {
	val interface{}
	next *node
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{nil}
}

func (l *LinkedListStack) IsEmpty() bool {
	return l.topNode == nil
}

func (l *LinkedListStack) Push(v interface{}) {
	l.topNode = &node{val: v, next: l.topNode}
}

func (l *LinkedListStack) Pop() interface{} {
	if l.IsEmpty() {
		return nil
	}
	v := l.topNode.val
	l.topNode = l.topNode.next

	return v
}

func (l *LinkedListStack) Top() interface{} {
	if l.IsEmpty() {
		return nil
	}
	return l.topNode.val
}

func (l *LinkedListStack) Flush() {
	l.topNode = nil
}

func (l *LinkedListStack) String() string {
	if l.IsEmpty() {
		return "[]"
	}
	res := "["
	cur := l.topNode
	for cur != nil {
		res += fmt.Sprintf("%v ", cur.val)
		cur = cur.next
	}

	return res[:len(res)-1] + "]"
}