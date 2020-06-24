package offer

type CQueue struct {
	pushStack []int // 插入栈
	popStack []int // 弹出栈
}


func Constructor() CQueue {
	return CQueue{}
}


func (this *CQueue) AppendTail(value int)  {
	this.pushStack = append(this.pushStack, value)
}


func (this *CQueue) DeleteHead() int {
	if len(this.popStack) != 0 {
		v := this.popStack[len(this.popStack)-1]
		this.popStack = this.popStack[:len(this.popStack)-1]
		return v
	}
	for len(this.pushStack) > 0 {
		v := this.pushStack[len(this.pushStack)-1]
		this.pushStack = this.pushStack[:len(this.pushStack)-1]
		this.popStack = append(this.popStack, v)
	}
	if len(this.popStack) == 0 { return -1 }
	return this.DeleteHead()
}
