package skipList

import (
	"math"
	"math/rand"
	"fmt"
)

// 跳表实现
// todo 不是很理解，后面重看

const (
	MAX_LEVEL = 16 // 最高层数
)

// skipListNode 跳表结点
type skipListNode struct {
	// 跳表保存的值
	v        interface{}
	// 用于排序的分值
	score    int
	// 层高
	level    int
	// 每层前进指针
	forwards []*skipListNode
}

func newSkipListNode(v interface{}, score, level int) *skipListNode {
	return &skipListNode{
		v: v,
		score: score,
		level: level,
		forwards: make([]*skipListNode, level, level),
	}
}

// SkipList 跳表
type SkipList struct {
	// 跳表头结点
	head *skipListNode
	// 跳表当前层数
	level int
	// 跳表长度
	length int
}

func NewSkipList() *SkipList {
	return &SkipList{
		head: newSkipListNode(0, math.MinInt32, MAX_LEVEL),
		level: 1,
	}
}

// Len 获取跳表长度
func (sl *SkipList) Len() int {
	return sl.length
}

// Level 获取跳表层级
func (sl *SkipList) Level() int {
	return sl.level
}

// Insert 插入结点到跳表中
func (sl *SkipList) Insert(v interface{}, score int) int {
	if v == nil {
		return 1
	}

	// 查找插入位置
	cur := sl.head
	// 记录每层的路径
	update := [MAX_LEVEL]*skipListNode{}
	for i := MAX_LEVEL - 1; i >= 0; i-- {
		for cur.forwards[i] != nil {
			if cur.forwards[i].v == v {
				return 2
			}
			if cur.forwards[i].score > score {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
		if cur.forwards[i] == nil {
			update[i] = cur
		}
	}

	// 通过随机算法获取该结点层数
	level := 1
	for i := 1; i < MAX_LEVEL; i++ {
		if rand.Int31() % 7 == 1 {
			level++
		}
	}

	// 创建一个新的跳表结点
	newNode := newSkipListNode(v, score, level)

	// 原有结点连接
	for i := 0; i <= level - 1; i++ {
		next := update[i].forwards[i]
		update[i].forwards[i] = newNode
		newNode.forwards[i] = next
	}

	// 如果当前结点的层数大于之前跳表的层数
	// 更新当前跳表层数
	if level > sl.level {
		sl.level = level
	}

	// 更新跳表长度
	sl.length++

	return 0
}

// Find 查找
func (sl *SkipList) Find(v interface{}, score int) *skipListNode {
	if v == nil || sl.length == 0 {
		return nil
	}

	cur := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for cur.forwards[i] != nil {
			if cur.forwards[i].score == score && cur.forwards[i].v == v {
				return cur.forwards[i]
			} else if cur.forwards[i].score > score {
				break
			}
			cur = cur.forwards[i]
		}
	}

	return nil
}

// Delete 删除结点
func (sl *SkipList) Delete(v interface{}, score int) int {
	if v == nil {
		return 1
	}

	// 查找前驱结点
	cur := sl.head
	// 记录前驱路径
	update := [MAX_LEVEL]*skipListNode{}
	for i := sl.level - 1; i >= 0; i-- {
		update[i] = sl.head
		for cur.forwards[i] != nil {
			if cur.forwards[i].score == score && cur.forwards[i].v == v {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
	}

	cur = update[0].forwards[0]
	for i := cur.level - 1; i >= 0; i-- {
		if update[i] == sl.head && cur.forwards[i] == nil {
			sl.level = i
		}

		if update[i].forwards[i] == nil {
			update[i].forwards[i] = nil
		} else {
			update[i].forwards[i] = update[i].forwards[i].forwards[i]
		}
	}

	sl.length--

	return 0
}

func (sl *SkipList) String() string {
	return fmt.Sprintf("level:%+v, length:%+v", sl.level, sl.length)
}