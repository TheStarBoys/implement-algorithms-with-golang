package dataStructure

import "errors"

// 单链表
type LinkList struct {
	Data interface{}
	Next *LinkList
}

func (ll *LinkList)GetElem(i int) (interface{}, error) {
	p := ll.Next
	j := 0
	for p != nil && j < i {
		p = p.Next
		j++
	}
	if p != nil {	// 此时p指向了i处，若为空，说明i处没值
		return p.Data, nil
	}
	return nil, errors.New("i处无值")
}
func (ll *LinkList)ListInsert(i int, e interface{}) error{
	p := ll
	j := 0
	for p != nil && j < i {	// 找到i结点
		p = p.Next
		j++
	}
	if p != nil {
		l := &LinkList{}
		l.Next = p.Next
		p.Next = l
		l.Data = e
		return nil
	}
	return errors.New("插入失败")
}
func (ll *LinkList)ListDelete(i int) error{
	p := ll
	j := 0
	for p != nil && j < i {	// 找到i结点
		p = p.Next
		j++
	}
	if p != nil {
		p.Next = p.Next.Next
		return nil
	}
	return errors.New("删除数据失败")
}