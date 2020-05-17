package AVL

import (
	"math"
)

// 定义平衡二叉树的结点结构
type AVLNode struct {
	depth int
	parent *AVLNode
	val int
	lChild *AVLNode
	rChild *AVLNode
}

// LL型调整函数——右旋
// 返回：新父节点
func LL_Rotate(node *AVLNode) *AVLNode {

	parent := node.parent
	son := node.lChild
	if son.rChild != nil {
		son.rChild.parent = node
	}
	node.lChild = son.rChild
	updateDepth(node)
	son.rChild = node
	son.parent = parent
	if parent != nil {
		if parent.lChild == node {
			parent.lChild = son
		} else {
			parent.rChild = son
		}
	}
	node.parent = son
	updateDepth(son)

	return son
}

// RR型调整函数——左旋
// 返回：新父节点
func RR_Rotate(node *AVLNode) *AVLNode {
	parent := node.parent
	son := node.rChild
	if son.lChild != nil {
		son.lChild.parent = node
	}
	node.rChild = son.lChild
	updateDepth(node)
	son.lChild = node
	son.parent = parent
	if parent != nil {
		if parent.lChild == node {
			parent.lChild = son
		} else {
			parent.rChild = son
		}
	}
	node.parent = son
	updateDepth(son)

	return son
}

// LR型，先左旋，后右旋
// 返回：新父节点
func LR_Rotate(node *AVLNode) *AVLNode {
	RR_Rotate(node.lChild)
	return LL_Rotate(node)
}

// RL型：先右旋，后左旋
// 返回：新父节点
func RL_Rotate(node *AVLNode) *AVLNode {
	LL_Rotate(node.rChild)
	return RR_Rotate(node)
}

// 更新当前深度
func updateDepth(node *AVLNode) {
	if node == nil {
		return
	}

	left := getBalance(node.lChild)
	right := getBalance(node.rChild)
	node.depth = int(math.Max(float64(left), float64(right)))
}

// 获取当前节点的深度
func getBalance(node *AVLNode) int {
	if node == nil {
		return 0
	}

	return node.depth
}

// 返回当前平衡因子
func isBalance(node *AVLNode) int {
	if node == nil {
		return 0
	}

	return getBalance(node.lChild) - getBalance(node.rChild)
}