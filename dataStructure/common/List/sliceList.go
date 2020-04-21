package List

import "errors"

/*
线性表
 */
type List struct {
	Data []interface{}
	size int
}

// Operation
// 初始化操作，建立一个空的线性表L
func InitList() *List{
	return &List{}
}
// 判断线性表是否为空表，若线性表为空
// 则返回true,否则返回false
func (l *List)ListEmpty() bool{
	if l.size == 0 {
		return false
	}
	return true
}
// 将线性表清空
func (l *List)ClearList() {
	l.Data = []interface{}{}
	l.size = 0
}
// 将线性表L中的第i下标元素值返回给e
func (l *List)GetElem(i int) (interface{}, error) {
	if l.size == 0 || i < 0 || i > l.size- 1 {
		return 0, errors.New("index out of range")
	}
	return l.Data[i], nil
}
// 在线性表L中查找与给定值e相等的元素，如果查找成功
// 返回该元素在表中序号表示成功
// 否则， 返回-1表示失败
func (l *List)LocateElem(e interface{}) int{
	for i, v := range l.Data {
		if v == e {
			return i
		}
	}
	return -1
}
// 在线性表L中第i个位置插入新元素e
func (l *List)ListInsert(i int, e interface{}) error{
	if  i < 0 || i > l.size {
		return errors.New("index out of range")
	}
	if l.size == i {
		l.Data = append(l.Data, e)
		l.size++
		return nil
	}
	temp := l.Data[i+1:]
	l.Data = append(l.Data[:i], e)
	l.Data = append(l.Data, temp...)
	l.size++
	return nil
}
// 删除线性表L中第i下标元素
// 并用e返回其值
func (l *List)ListDelete(i int) (e interface{}, err error) {
	if i > l.size-1 || i < 0{
		return -1 ,errors.New("index out of range")
	}
	e = l.Data[i]
	if l.size == 0 {
		return e, errors.New("list is empty")
	}
	l.Data = append(l.Data[:i], l.Data[i+1:]...)
	l.size--
	return e , nil
}
// 返回线性表L的元素个数
func (l *List)ListLength() int{
	return l.size
}
