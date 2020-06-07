# 堆

## 简介

堆是一种特殊的树，只要满足以下两点，它就是一个堆：

- 堆是一个完全二叉树
- 堆中每一个节点的值都必须大于等于（或小于等于）其子树中每个节点的值



大顶堆：对于每个节点的值都大于等于子树中每个节点值的堆，我们叫作“大顶堆”。

小顶堆：对于每个节点的值都小于等于子树中每个节点值的堆，我们叫作“小顶堆”。

## 如果实现一个堆？

要实现一个堆，首先要知道，**堆都支持哪些操作**以及**如何存储一个堆**。

### 如何存储一个堆

完全二叉树比较适合用数组来存储。用数组来存储完全二叉树是非常节省存储空间的。因为我们不需要存储左右子节点的指针，单纯地通过数组的下标，就可以找到一个节点的左右子节点和父节点。 

我画了一个用数组存储堆的例子，你可以先看下。 

 ![img](https://static001.geekbang.org/resource/image/4d/1e/4d349f57947df6590a2dd1364c3b0b1e.jpg) 

 从图中我们可以看到，数组中下标为 `i` 的节点的左子节点，就是下标为 `i∗2` 的节点，右子节点就是下标为 `i∗2+1` 的节点，父节点就是下标为 `i/2` 的节点。 

### 堆的操作

若没有特别说明，都是以大顶堆进行讲解。

####  1. 往堆中插入一个元素

往堆中插入一个元素后，我们需要继续满足堆的两个特性。

如果我们把新插入的元素放到堆的最后，你可以看我画的这个图，是不是不符合堆的特性了？于是，我们就需要进行调整，让其重新满足堆的特性，这个过程我们起了一个名字，就叫作**堆化**（heapify）。

堆化实际上有两种，从下往上和从上往下。这里我先讲**从下往上**的堆化方法。 

 ![img](https://static001.geekbang.org/resource/image/e5/22/e578654f930002a140ebcf72b11eb722.jpg) 

堆化非常简单，就是顺着节点所在的路径，向上或者向下，对比，然后交换。 

我这里画了一张堆化的过程分解图。我们可以让新插入的节点与父节点对比大小。如果不满足子节点小于等于父节点的大小关系，我们就互换两个节点。一直重复这个过程，直到父子节点之间满足刚说的那种大小关系。 

 ![img](https://static001.geekbang.org/resource/image/e3/0e/e3744661e038e4ae570316bc862b2c0e.jpg) 

我将上面讲的往堆中插入数据的过程，翻译成了代码，你可以结合着一块看。 

```go
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

func swap(data []Elem, i, j int) {
	data[i], data[j] = data[j], data[i]
}
```



#### 2. 删除堆顶元素

从堆的定义的第二条中，任何节点的值都大于等于（或小于等于）子树节点的值，我们可以发现，堆顶元素存储的就是堆中数据的最大值或者最小值。

假设我们构造的是大顶堆，堆顶元素就是最大的元素。当我们删除堆顶元素之后，就需要把第二大的元素放到堆顶，那第二大元素肯定会出现在左右子节点中。然后我们再迭代地删除第二大节点，以此类推，直到叶子节点被删除。

这里我也画了一个分解图。不过这种方法有点问题，就是最后堆化出来的堆并不满足完全二叉树的特性。 

 ![img](https://static001.geekbang.org/resource/image/59/81/5916121b08da6fc0636edf1fc24b5a81.jpg) 

实际上，我们稍微改变一下思路，就可以解决这个问题。你看我画的下面这幅图。我们把最后一个节点放到堆顶，然后利用同样的父子节点对比方法。对于不满足父子节点大小关系的，互换两个节点，并且重复进行这个过程，直到父子节点之间满足大小关系为止。这就是**从上往下的堆化方法**。 

因为我们移除的是数组中的最后一个元素，而在堆化的过程中，都是交换操作，不会出现数组中的“空洞”，所以这种方法堆化之后的结果，肯定满足完全二叉树的特性。 

 ![img](https://static001.geekbang.org/resource/image/11/60/110d6f442e718f86d2a1d16095513260.jpg) 

我把上面的删除过程同样也翻译成了代码，贴在这里，你可以结合着看。 

```go
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
```

我们知道，一个包含 `n` 个节点的完全二叉树，树的高度不会超过 `log2n`。堆化的过程是顺着节点所在路径比较交换的，所以堆化的时间复杂度跟树的高度成正比，也就是 `O(logn)`。插入数据和删除堆顶元素的主要逻辑就是堆化，所以，往堆中插入一个元素和删除堆顶元素的时间复杂度都是 `O(logn)`。 



#### 3. 建堆

我们首先将数组原地建成一个堆。所谓“原地”就是，不借助另一个数组，就在原数组上操作。建堆的过程，有两种思路。 

第一种是借助我们前面讲的，在堆中插入一个元素的思路。尽管数组中包含 n 个数据，但是我们可以假设，起初堆中只包含一个数据，就是下标为 1 的数据。然后，我们调用前面讲的插入操作，将下标从 2 到 n 的数据依次插入到堆中。这样我们就将包含 n 个数据的数组，组织成了堆。 

第二种实现思路，跟第一种截然相反，也是我这里要详细讲的。第一种建堆思路的处理过程是从前往后处理数组数据，并且每个数据插入堆中时，都是从下往上堆化。而第二种实现思路，是从后往前处理数组，并且每个数据都是从上往下堆化。 

我举了一个例子，并且画了一个第二种实现思路的建堆分解步骤图，你可以看下。因为叶子节点往下堆化只能自己跟自己比较，所以我们直接从第一个非叶子节点开始，依次堆化就行了。 

 ![img](https://static001.geekbang.org/resource/image/50/1e/50c1e6bc6fe68378d0a66bdccfff441e.jpg) 

 ![img](https://static001.geekbang.org/resource/image/aa/9d/aabb8d15b1b92d5e040895589c60419d.jpg) 

对于程序员来说，看代码可能更好理解一些，所以，我将第二种实现思路翻译成了代码，你可以看下。 

```go
// 建堆
// 时间复杂度为 O(n)
// n 为当前元素个数，下标为 0 处除外
func BuildHeap(data []Elem, n int) {
	for i := n/2; i >= 1; i-- {
		heapify(data, n, i)
	}
}
```

你可能已经发现了，在这段代码中，我们对下标从 `n/2` 开始到 `1` 的数据进行堆化，下标是 `2n+1` 到 `n` 的节点是叶子节点，我们不需要堆化。实际上，对于完全二叉树来说，下标从 `2n+1` 到 `n` 的节点都是叶子节点。

#### 4. 排序

建堆结束之后，数组中的数据已经是按照大顶堆的特性来组织的。数组中的第一个元素就是堆顶，也就是最大的元素。我们把它跟最后一个元素交换，那最大元素就放到了下标为 `n` 的位置。

这个过程有点类似上面讲的**“删除堆顶元素”**的操作，当堆顶元素移除之后，我们把下标为 `n` 的元素放到堆顶，然后再通过堆化的方法，将剩下的 `n−1` 个元素重新构建成堆。堆化完成之后，我们再取堆顶的元素，放到下标是 `n−1` 的位置，一直重复这个过程，直到最后堆中只剩下标为 `1` 的一个元素，排序工作就完成了。

 ![img](https://static001.geekbang.org/resource/image/23/d1/23958f889ca48dbb8373f521708408d1.jpg) 

堆排序的过程，我也翻译成了代码。结合着代码看，你理解起来应该会更加容易。 

```go
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
```



## 堆的应用

本节内容是基于实现Go语言中 `container/heap` 包中堆的接口，实现的一个堆。本节内容中的`sdHeap` 其实就是 `container/heap`  包，只不过给它起了个别名。以下是实现堆接口的代码：

```go
import (
	sdHeap "container/heap"
)
// 先实现一个Heap需要的接口
type myHeap []int

// 通过Less的不同，实现大顶堆和小顶堆的效果
// 此实现是小顶堆
func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len() - 1], (*h)[h.Len() - 1]
	return
}

func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *myHeap) Top() (v interface{}) {
	if h.Len() == 0 {
		return nil
	}
	return (*h)[0]
}

// 仿sort包里的reverse用法，可以将具体的实现进行转换
// 比如myHeap是实现的小顶堆，通过Reverse()可以变成大顶堆
type reverseInterface interface {
	sdHeap.Interface
	Top() interface{}
}
type reverse struct {
	reverseInterface
}

func (r *reverse) Less(i, j int) bool {
	return r.reverseInterface.Less(j, i)
}

func Reverse(p reverseInterface) *reverse {
	return &reverse{
		reverseInterface: p,
	}
}
```



### 1. 优先级队列

**优先级队列**，顾名思义，它首先应该是一个队列。我们前面讲过，队列最大的特性就是先进先出。不过，在优先级队列中，数据的出队顺序不是先进先出，而是按照优先级来，优先级最高的，最先出队。 

如何实现一个优先级队列呢？方法有很多，但是用堆来实现是最直接、最高效的。这是因为，堆和优先级队列非常相似。一个堆就可以看作一个优先级队列。很多时候，它们只是概念上的区分而已。往优先级队列中插入一个元素，就相当于往堆中插入一个元素；从优先级队列中取出优先级最高的元素，就相当于取出堆顶元素。 

你可别小看这个优先级队列，它的应用场景非常多。我们后面要讲的很多数据结构和算法都要依赖它。比如，赫夫曼编码、图的最短路径、最小生成树算法等等。不仅如此，很多语言中，都提供了优先级队列的实现，比如，Java 的 PriorityQueue，C++ 的 priority_queue 等。 

```go
// 优先级队列，优先级高的先出队
h := Reverse(new(myHeap))
sdHeap.Push(h, 2) // [2]
sdHeap.Push(h, 3) // [3 2]
sdHeap.Push(h, 5) // [5 2 3]
sdHeap.Push(h, 1) // [5 2 3 1]
sdHeap.Remove(h, 0) // [3 2 1]
for h.Len() > 0 {
	fmt.Printf("%d ", sdHeap.Pop(h))
}
// Output:
// 3 2 1
```

完整版的实现：

```go
import "container/heap"

// Node 队列节点
type Node struct {
	value    interface{}
	priority int // 优先级
}

// 实现堆接口
type myheap []Node

func (h myheap) Len() int { return len(h) }
// 如果是大于等于，那么相同优先级的，先进的将先出
// 如果是大于，那么相同优先级的，后进的先出
func (h myheap) Less(i, j int) bool { return h[i].priority >= h[j].priority }
func (h myheap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *myheap) Push(v interface{}) {
	*h = append(*h, v.(Node))
}

func (h *myheap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len() - 1], (*h)[h.Len()-1]
	return
}

func (h myheap) Top() (v interface{}) {
	return h[h.Len()-1]
}

// PriorityQueue 优先级队列
type PriorityQueue struct {
	heap *myheap
	capacity int
}

func NewPriorityQueue(capacity int) PriorityQueue {
	return PriorityQueue{
		heap: new(myheap),
		capacity: capacity,
	}
}

func (q PriorityQueue) Len() int {
	return q.heap.Len()
}

// Push 入队
func (q *PriorityQueue) Push(node Node) {
	if q.Len() >= q.capacity {
		// 队列已满
		return
	}
	heap.Push(q.heap, node)
}

// Pop 出队
func (q *PriorityQueue) Pop() *Node {
	if q.Len() == 0 {
		return nil
	}
	node :=  heap.Pop(q.heap).(Node)
	return &node
}

// Front 获取队首节点
func (q *PriorityQueue) Front() *Node {
	if q.Len() == 0 {
		return nil
	}
	node := q.heap.Top().(Node)
	return &node
}
```



### 2. 求Top K

我把这种求 Top K 的问题抽象成两类。

- 一类是针对静态数据集合，也就是说数据集合事先确定，不会再变。
- 另一类是针对动态数据集合，也就是说数据集合事先并不确定，有数据动态地加入到集合中。 

针对**静态数据**，如何在一个包含 n 个数据的数组中，查找前 K 大数据呢？我们可以维护一个大小为 K 的小顶堆，顺序遍历数组，从数组中取出数据与堆顶元素比较。如果比堆顶元素大，我们就把堆顶元素删除，并且将这个元素插入到堆中；如果比堆顶元素小，则不做处理，继续遍历数组。这样等数组中的数据都遍历完之后，堆中的数据就是前 K 大数据了。 

遍历数组需要 `O(n)` 的时间复杂度，一次堆化操作需要 `O(logK)` 的时间复杂度，所以最坏情况下，`n` 个元素都入堆一次，时间复杂度就是 `O(nlogK)`。 

针对**动态数据**求得 Top K 就是实时 Top K。怎么理解呢？我举一个例子。一个数据集合中有两个操作，一个是添加数据，另一个询问当前的前 K 大数据。 

如果每次询问前 K 大数据，我们都基于当前的数据重新计算的话，那时间复杂度就是 `O(nlogK)`，`n` 表示当前的数据的大小。实际上，我们可以一直都维护一个 K 大小的小顶堆，当有数据被添加到集合中时，我们就拿它与堆顶的元素对比。如果比堆顶元素大，我们就把堆顶元素删除，并且将这个元素插入到堆中；如果比堆顶元素小，则不做处理。这样，无论任何时候需要查询当前的前 K 大数据，我们都可以立刻返回给他。 

这里给了一个代码实现：

```go
// 如果是一个静态数据，直接调用即可
// 如果是一个动态数据，那么我们需要一直持有一个大小为K的堆
// 每次的数据添加，我们都跟堆顶进行比较，考虑是否将堆顶元素移除，数组元素插入堆
// 这样，堆中的K个数据就是实时的前K大的数据
func topK(data []int, k int) []int {
	h := new(myHeap)
	// 申请一个大小为 k 的堆，并且往里填充该元素的最小值
	for i := 0; i < k; i++ {
		h.Push(math.MinInt64)
	}
	// 遍历数组
	for _, v := range data {
		// 如果数组当前元素比堆顶元素大
		if v > (*h)[0] {
			// 移除堆顶元素
			sdHeap.Pop(h)
			// 将数组元素插入堆中
			sdHeap.Push(h, v)
		}
	}
	return []int(*h)
}
```



### 3. 求中位数

```go
// 应用3：利用堆求中位数
func TestFindMedian(t *testing.T) {
	data := []int{1, 2, 3}
	fmt.Println(findMedian(data))
	// Output:
	// 2
	data = []int{1, 2, 3, 4}
	fmt.Println(findMedian(data))
	// Output:
	// 2.5
}

func findMedian(data []int) float64 {
	finder := Constructor()
	for _, v := range data {
		finder.AddNum(v)
	}

	return finder.FindMedian()
}

type MedianFinder struct {
	maxHeap *myHeap
	minHeap reverseInterface
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		maxHeap: new(myHeap),
		minHeap: Reverse(new(myHeap)),
	}
}

func (this *MedianFinder) AddNum(num int)  {
	sdHeap.Push(this.maxHeap, num) // Add to max heap

	sdHeap.Push(this.minHeap, sdHeap.Pop(this.maxHeap)) // balancing step

	// maintain size property
	if this.maxHeap.Len() < this.minHeap.Len() {
		sdHeap.Push(this.maxHeap, sdHeap.Pop(this.minHeap))
	}
}


func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeap.Len() == this.minHeap.Len() {
		v1 := this.maxHeap.Top().(int)
		v2 := this.minHeap.Top().(int)
		return float64(v1 + v2) / 2
	}
	return float64(this.maxHeap.Top().(int))
}
```



## 引用

> [数据结构与算法之美](https://time.geekbang.org/column/intro/100017301)