package heap

// 大顶堆
// 往堆中插入一个元素和删除堆顶元素的时间复杂度都是 O(log n)
type Heap struct {
	data []Elem // 由于堆是完全二叉树，所以可以用数组来存，并且下标为 0 处不存储任何数据
	n int // 堆可存储的最大元素个数
	count int // 堆中已经存储的元素个数
}

// 定义堆中元素的接口
// 以此让堆中的元素可比较
type Elem interface {
	Less(elem Elem) bool
}

func NewHeap(capacity int) *Heap {
	return &Heap{
		data: make([]Elem, capacity + 1),
		n: capacity,
	}
}

// 建堆
// 时间复杂度为 O(n)
// n 为当前元素个数，下标为 0 处除外
func BuildHeap(data []Elem, n int) {
	for i := n/2; i >= 1; i-- {
		heapify(data, n, i)
	}
}

// 只能排序 1 ~ 切片长度
// 时间复杂度 O(n log n) 其中
// 建堆为 O(n),
// 循环部分迭代了 n 个节点，每个节点堆化时间复杂度为 O(log n)，所以迭代部分时间复杂度为 O(n log n)。
// 忽略建堆时的低阶复杂度，得到 O(n log n)
func sort(data []Elem, n int) {
	BuildHeap(data, n)
	for k := n; k > 1; {
		swap(data, 1, k)
		k--
		heapify(data, k, 1)
	}
}

func (h *Heap) Insert(elem Elem) {
	if h.count >= h.n {
		// 堆满了
		return
	}
	// 将数据存到叶子节点
	h.count++
	h.data[h.count] = elem
	// 自下往上堆化
	// 当前节点 i 跟其父节点 i / 2 进行比较，如果父节点小于当前节点，则交换
	for i := h.count; i / 2 > 0 && h.data[i/2].Less(h.data[i]); i /= 2 {
		// 交换节点
		//h.data[i/2], h.data[i] = h.data[i], h.data[i/2]
		swap(h.data, i/2, i)
	}
}

func (h *Heap) RemoveMax() {
	if h.count == 0 {
		return
	}
	// 将堆的最后一个节点放到堆顶
	h.data[1] = h.data[h.count]
	h.count--
	// 堆化
	heapify(h.data, h.count, 1)
}

// 自上向下堆化
// n 为当前元素个数, i 为移除节点
func heapify(data []Elem, n, i int) {
	for {
		maxPos := i
		// 从 i 节点的左右孩子中找到最大的
		if i * 2 <= n && data[i].Less(data[i * 2]) {
			maxPos = i * 2
		}
		if i * 2 + 1 <= n && data[maxPos].Less(data[i * 2 + 1]) {
			maxPos = i * 2 + 1
		}
		// 当前节点的左右孩子均小于它
		if maxPos == i {
			break
		}
		swap(data, maxPos, i)
		i = maxPos
	}
}

func swap(data []Elem, i, j int) {
	data[i], data[j] = data[j], data[i]
}