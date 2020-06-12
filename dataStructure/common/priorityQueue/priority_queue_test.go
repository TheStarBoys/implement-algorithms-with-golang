package priorityQueue

import (
	"testing"
	"fmt"
	"math"
	"sort"
	"math/rand"
)

func TestPriorityQueue(t *testing.T) {
	queue := NewPriorityQueue(nil,4)
	if node := queue.Front(); node != nil {
		t.Errorf("expect: node is nil")
	}
	queue.Push(QNode{3, 2})
	if node := queue.Front(); node == nil || node.GetPriority() != 2 && node.GetValue().(int) != 3 {
		t.Errorf("node is not Node{3, 2}")
	}
	queue.Push(QNode{1, 7})
	queue.Push(QNode{5, 4})
	queue.Push(QNode{2, 4})
	queue.Push(QNode{999, 999}) // queue is full, operation failed
	for queue.Len() > 0 {
		node := queue.Pop()
		fmt.Printf("value: %d, priority: %d\n", node.GetValue(), node.GetPriority())
	}
	// Output:
	// value: 1, priority: 7
	// value: 5, priority: 4
	// value: 2, priority: 4
	// value: 3, priority: 2
}

// 应用1：模拟赫夫曼编码
func TestHuffman(t *testing.T) {
	/*
	  	    	x
	    	   / \
	          x   a
	         / \
	        x   c
	       / \
	      b   d
	 */
	data := []rune("abaaccd")
	table := map[rune]int{}
	for _, r := range data {
		table[r]++
	}
	root := getHuffmanRoot(table)
	fmt.Println()
	fmt.Println("----- Level Order ----")
	root.LevelOrder()
	// Output:
	// -----Level Order----
	// {2147483647, 7}
	// {2147483647, 4} {97, 3}
	// {2147483647, 2} {99, 2}
	// {98, 1} {100, 1}

	fmt.Println("----- Huffman Code ---")
	res := HuffmanCoding(data)
	type hf struct {
		r rune
		code string
	}

	slice := []*hf{}
	for r, code := range res {
		slice = append(slice, &hf{r, code})
	}
	sort.Slice(slice, func(i, j int) bool {
		if slice[i].code > slice[j].code {
			return true
		}
		return false
	})
	for _, v := range slice {
		fmt.Printf("%c : %s\n", v.r, v.code)
	}
	// Output:
	// ----- Huffman Code ---
	// a : 1
	// c : 01
	// d : 001
	// b : 000

	encodeStr, encodeRule := HuffmanEncode(data)
	fmt.Printf("Huffman Encode String: %s\n", encodeStr)
	decodeStr := HuffmanDecode(encodeStr, encodeRule)
	fmt.Printf("Huffman Decode String: %s\n", decodeStr)

	// Output:
	// Huffman Encode String: 110111110000100101
	// Huffman Decode String: abaaccde

	f := func(data []rune) {
		rawStr := string(data)
		encodeStr, encodeRule := HuffmanEncode(data)
		decodeStr := HuffmanDecode(encodeStr, encodeRule)
		if rawStr != decodeStr {
			t.Errorf("expect equal, actual: rawStr: %s, decodeStr: %s\n", rawStr, decodeStr)
		}
	}

	f([]rune("f"))
	f([]rune(""))
	for i := 0; i < 100; i++ {
		length := rand.Intn(100) + 1
		data := make([]rune, length)
		for i := range data {
			b := rand.Intn(25) + 97 // 生成小写字母
			data[i] = rune(b)
		}
		fmt.Println("rawStr:", string(data))
		f(data)
	}
}

type HuffmanNode struct {
	*QNode // 组合得到基本的优先级队列结点特性
	left *HuffmanNode
	right *HuffmanNode
}

// r 字符, f 频率
func NewHuffmanNode(r rune, f int, left, right *HuffmanNode) *HuffmanNode {
	return &HuffmanNode{
		&QNode{r, f},
		left,
		right,
	}
}

// LevelOrder 层序遍历该二叉树
func (n *HuffmanNode) LevelOrder() {
	queue := []*HuffmanNode{}
	queue = append(queue, n)
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			cur := queue[0]
			queue = queue[1:]
			fmt.Printf("{%v, %d} ", cur.value, cur.priority)
			if cur.left != nil {
				queue = append(queue, cur.left)
			}
			if cur.right != nil {
				queue = append(queue, cur.right)
			}
		}
		fmt.Println()
	}
}

// HuffmanEncode 编码
func HuffmanEncode(data []rune) (string, map[rune]string) {
	if len(data) == 0 {
		// 不需要编码
		return "", nil
	}
	var res string
	rule := HuffmanCoding(data)
	for _, r := range data {
		res += rule[r]
	}

	return res, rule
}

// HuffmanDecode 解码。传入编码后的字符串，跟编码规则
func HuffmanDecode(encodeStr string, encodeRule map[rune]string) string {
	if encodeStr == "" {
		return ""
	}
	var (
		res string // 解码结果
		maxCodeLength int // 最长编码长度
		prevIndx int // 前一个可解码的编码的结尾下标
	)
	// 根据编码规则，生成对应的解码规则
	decodeRule := map[string]rune{}
	for r, code := range encodeRule {
		if maxCodeLength < len(code) {
			maxCodeLength = len(code)
		}
		decodeRule[code] = r
	}

	for i := 0; i < len(encodeStr); {
		// 找到最长可解码二进制
		// j 取值范围 [prevIndx : prevIndx + maxCodeLength - 1]
		j := prevIndx + maxCodeLength - 1
		// 如果生成的可取值范围超过了编码的长度
		if j >= len(encodeStr) {
			// 取值范围变为 [prevIndx : len(encodeStr) - 1]
			j = len(encodeStr) - 1
		}
		for ; j >= prevIndx; j-- {
			// code 待解码数据
			code := encodeStr[prevIndx:j+1]
			// 查询是否有对应的解码规则
			_, ok := decodeRule[code]
			if ok {
				break
			}
		}
		j++
		r := encodeStr[prevIndx:j]
		res += string(decodeRule[r])
		prevIndx = j
		i = j
	}

	return res
}

// HuffmanCoding 获取编码规则，出现频率高的字符，编码长度应该尽可能的短，并且无歧义
func HuffmanCoding(data []rune) map[rune]string {
	// 统计频率
	table := map[rune]int{}
	for _, r := range data {
		table[r]++
	}
	if len(table) == 1 {
		return map[rune]string{
			data[0]: "0",
		}
	}
	node := getHuffmanRoot(table)
	res := map[rune]string{}
	getCoding(node, "", res)

	return res
}

// getCoding 返回字符对应的编码
// 二叉树左节点权值为0，右节点权值为1
// 达到叶子节点时所经过的节点的权值拼接的结果，就是该叶子节点的编码
func getCoding(node *HuffmanNode, code string, table map[rune]string) {
	if node == nil {
		return
	}
	if node != nil && node.left == nil && node.right == nil {
		r := node.GetValue().(rune)
		table[r] = code
		return
	}
	getCoding(node.left, code + "0", table)
	getCoding(node.right, code + "1", table)
}

// getHuffmanRoot 返回其根节点
// 将数据排成优先级队列，每取出两个节点，构造一个新的节点，作为它们的父节点，最终形成一颗二叉树
func getHuffmanRoot(data map[rune]int) *HuffmanNode {
	// 优先级队列，用频率表示优先级，数值小的优先
	queue := NewPriorityQueue(Reverse(new(myheap)), len(data))
	for r, f := range data {
		node := NewHuffmanNode(r, f, nil, nil)
		queue.Push(node)
	}
	// 由于每次取两个节点，生成一个新节点，最终必定只剩一个节点，故将其作为终止条件
	for queue.Len() > 1 {
		// 从优先级队列中每次取两个节点
		left := queue.Pop().(*HuffmanNode)
		right := queue.Pop().(*HuffmanNode)
		if left.value.(rune) != rune(math.MaxInt32) {
			left, right = right, left
		}
		// 创建一个新的节点，将新取出的节点作为该节点的左右孩子节点，该节点的频率为孩子节点频率之和
		node := NewHuffmanNode(rune(math.MaxInt32), left.GetPriority() + right.GetPriority(), left, right)
		queue.Push(node)
	}

	node := queue.Pop().(*HuffmanNode)

	return node
}