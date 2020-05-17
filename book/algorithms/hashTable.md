# 散列表

## 什么是散列表（哈希表）？

> 引用自：[什么是散列表（哈希表）？](https://mp.weixin.qq.com/s?__biz=MzUyNjQxNjYyMg==&mid=2247486499&idx=4&sn=7c9aee095ec972b1af842788c545d309&chksm=fa0e63a2cd79eab47649274533063bea75550d97aa55f60856f02037bc2f6b657270630f214b&scene=21#wechat_redirect)

### 前言

假设你们班级100个同学每个人的学号是由院系-年级-班级和编号组成，例如学号为01100168表示是1系，10级1班的68号。为了快速查找到68号的成绩信息，可以建立一张表，但是不能用学号作为下标，学号的数值实在太大。因此将学号除以1100100取余，即得到编号作为该表的下标，那么，要查找学号为01100168的成绩的时候，只要直接访问表下标为68的数据即可。这就能够在O（1）时间复杂度内完成成绩查找。

实际上这里就用到了散列的思想。本文重在介绍散列的思想以及散列需要考虑的问题。

### 散列表（哈希表）

理想散列表（哈希表）是一个包含关键字的具有固定大小的数组，它能够**以常数时间执行插入，删除和查找操作**。

- 每个关键字被映射到0到数组大小N-1范围，并且放到合适的位置，这个**映射规则就叫散列函数**
- 理想情况下，两个不同的关键字映射到不同的单元，然而由于数组单元有限，关键字范围可能远超数组单元，因此就会出现两个关键字散列到同一个值得时候，这就是**散列冲突**

### 实例演示

通过前面的描述，我们已经了解了一些基本概念，现在来看一个实例。
假设有一个大小为7的表，现在，要将13,18,19，50，20散列到表中。

- 选择散列函数，例如使用hash(x)=x%7作为散列函数
- 计算数据散列值，并放到合适的位置

计算13 % 7得到6，因此将13放到下标为6的位置：

| 0    | 1    | 2    | 3    | 4    | 5    | 6    |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
|      |      |      |      |      |      | 13   |

计算18 % 7得到4，因此将18放到下标为4的位置：

| 0    | 1    | 2    | 3    | 4    | 5    | 6    |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
|      |      |      |      | 18   |      | 13   |

计算19 % 7得到5，因此将19放到下标为5的位置：

| 0    | 1    | 2    | 3    | 4    | 5    | 6    |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
|      |      |      |      | 18   | 19   | 13   |

计算50 % 7得到1，因此将50放到下标为1的位置：

| 0    | 1    | 2    | 3    | 4    | 5    | 6    |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
|      | 50   |      |      | 18   | 19   | 13   |

计算20 % 7得到6，因此将20放到下标为6的位置，但是此时6的位置已经被占用了，因此就产生了**散列冲突**，关于散列冲突的解决，我们后面再介绍。

将数据散列之后，如何从表中查找呢？例如，查找数值为50的数据位置，只需要计算50 % 7，得到下标1，访问下标1的位置即可。但是如果考虑散列冲突，就没有那么简单了。

通过这个实例，了解了以下几个概念：

- 散列函数，散列函数的选择非常重要
- 散列冲突，涉及散列表时，因尽量避免散列冲突，对于冲突也要有好的解决方案
- 快速从散列表中查找数据

### 冲突解决

解决散列冲突通常有以下几种方法：

- 拉链法
- 开放定址法
- 再散列
- …

#### 拉链法

**分离链接法的做法是将同一个值的关键字保存在同一个表中**。例如，对于前面：

| 0    | 1    | 2    | 3    | 4    | 5    | 6    |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
|      | 50   |      |      | 18   | 19   | 13   |

如果再要插入元素20，则在下标为6的位置存储表头，而表的内容是13和20。

这种方法的特点是需要另外分配新的单元来存储散列到同一个位置的数据。

查找的时候，除了根据计算出来的散列值找到对应位置外，还需要在链表上进行搜索。而在**单链表上的查找速度是很慢**的。另外**散列函数如果设计得好，冲突的概率其实也会很小**。

#### 开放定址法

而开放定址法的思想是，**如果冲突发生，就选择另外一个可用的位置**。

而开放定址法中也有常见的几种策略。

- 线性探测法

还是以前面的为例：

| 0    | 1    | 2    | 3    | 4    | 5    | 6    |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
|      | 50   |      |      | 18   | 19   | 13   |

如果此时再要插入20，则20 % 7 = 6，但是6的位置已有元素，因此探测下一个位置（6+1）%7，在这里就是下标为0的位置。因此20的存储位置如下：

| 0    | 1    | 2    | 3    | 4    | 5    | 6    |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| 20   | 50   |      |      | 18   | 19   | 13   |

但这种方式的一个问题是，可能造成**一次聚集**，因为一旦冲突发生，为了处理冲突就会占用下一个位置，而如果冲突较多时，就会出现数据都**聚集在一块区域**。这样就会导致任何关键字都需要多次尝试才可能解决冲突。

- 平方探测法

顾名思义，如果说前面的探测函数是F（i）= i % 7，那么平方探测法就是F（i）= (i^2 )% 7。
但是这也同样会产生**二次聚集**问题。

- 双散列

为了**避免聚集**，在探测时选择跳跃式的探测，即再使用一个散列函数，用来计算探测的位置。假设前面的散列函数为hash1(X)，用于探测的散列函数为hash2(X)，那么一种流行的选择是F(i) = i * hash2(X)，即第一次冲突时探测hash1(X)+hash2(X)的位置，第二次探测
hash1(X)+2hash2(X)的位置。

可以看到，无论是哪种开放定址法，它都要求表足够大。

#### 再散列

我们前面也说到，散列表可以认为是具有固定大小的数组，那么如果插入新的数据时散列表已满，或者散列表所剩容量不多该怎么办？这个时候就需要再散列，常见做法是，建立一个是原来两倍大小的散列表，将原来表中的关键字重新散列到新表中。

### 散列表的应用

散列表应用很广泛。例如做文件校验或数字签名。当然还有快速查询功能的实现。例如，redis中的字典结构就使用了散列表，使用**MurmurHash算法**来计算字符串的hash值，并采用**拉链法**处理冲突，，当散列表的**装载因子**（关键字个数与散列表大小的比）接近某个大小时，进行**再散列**。

### 总结

一个设计良好的散列表能够几乎在O（1）时间复杂度内完成插入，删除和查找，但前提是**散列函数设计得足够优雅，以及有着合适散列冲突解决方案**。常见冲突解决方案有：

- 拉链法
- 开放地址检测法

其中拉链法在实际中是很常见的一种解决方案。另外本文重点说明**什么是散列表**（哈希表），因此没有涉及具体的代码，后面将会通过实例来看散列表的实际应用。

## 什么是哈希洪水攻击（Hash-Flooding Attack）？

> 引用自：[什么是哈希洪水攻击（Hash-Flooding Attack）？](https://mp.weixin.qq.com/s?__biz=MzUyNjQxNjYyMg==&mid=2247485591&idx=1&sn=54cde6fab4f11a0478f0c6447069d47c&chksm=fa0e6716cd79ee00e3502116a4f2999ee9648b99999317d450d9d6ee8ebef8cf4d39f2f79372&scene=21#wechat_redirect)

你开了一家菜鸟驿站，代收周围几百个小区的所有包裹。因为每天的包裹量很大，如何在取件人到来时快速找出他想要的包裹就成了很重要的问题。

正巧，开菜鸟驿站前你是个程序员，于是很自然就想到可以把包裹按照收件人的手机尾号进行堆放。只要以倒数第二位作为货架号、倒数第一位作为层号就可以了。比如手机尾号 24 的取件人的包裹就应该放在二号货架的第四层。

你家的菜鸟驿站开张了，生意很好，货架也够用。双十一期间虽然包裹多、货架上放不下，但问题不大，货架上放不下就堆在对应货架前的地上：二号货架第四层找不到包裹，就在二号货架前地上的包裹堆里找就可以了。特殊时期特殊对待，大家都是街坊，可以理解。

隔壁吴老二从小和你不对付，听说你赚了钱就气得浑身发抖。他摸清楚你是按照手机尾号放快递之后就想了个法子，特意去营业厅挑了一堆以 2X 结尾的手机号，每天从淘宝上买些不值钱又占地方的玩意儿寄到菜鸟驿站。于是二号货架常年是满的，货架前也常年堆着一堆包裹。别人的包裹都很快可以找到，耗费不到一分钟；而凡是手机尾号是 2X 的居民就都倒了血霉了，货架上基本没法找到，得去快件堆里一个个翻，每次不花个五六分钟别想找到包裹。这要是店里人手不够，尾号非 2X 的人还排队排在几个 2X 尾号的人后面取件，那酸爽……反正从此以后，吐槽源源不断，你的生意也一天不如一天了。

小区里有个产品经理朋友建议你可以多加个货架专门处理二号货架爆仓的情况，但你清楚地知道这样的特殊处理是治标不治本的。问题的根本是只要有人知道你是按照手机尾号放置包裹的，就可以用很小的成本「构造」包裹，让特定手机尾号的包裹像洪水般涌来（嗯，此处点题），降低你店里的工作效率，达到攻击的目的。

所以解决方法也就很明显了：**不要让取件人可以轻易猜到你是如何放置包裹的**。

在苦思冥想一周无果之后，你打听了一下，才发现隔壁村的菜鸟驿站居然是用现成的管理系统的。包裹入站时系统直接生成取件码，取件码均匀分散到货架层数，比如 1-2-1234 表示这是本驿站收到的第 1234 个包裹，应该放在一号货架第二层。这样取件人就没法通过构造特定的包裹进行攻击了。

**你恍然大悟，然后把手边的《编程珠玑》给扔了。**

## Bloom Filter 布隆过滤器

> 引用自：[Bloom Filter](https://mp.weixin.qq.com/s?__biz=MzUyNjQxNjYyMg==&mid=2247486559&idx=2&sn=bd9e3575882181e31e33f0a0094376db&chksm=fa0e63decd79eac8adc455059de6309a510951461496cc5ee3444e3675a0d5ffb94f2745b003&scene=21#wechat_redirect)

以下文章来源于crossoverJie ，作者crossoverJie

作者 | crossoverJie

来源 | crossoverJie

### 前言

最近有朋友问我这么一个面试题目：

> 现在有一个非常庞大的数据，假设全是 int 类型。现在我给你一个数，你需要告诉我它是否存在其中(尽量高效)。

需求其实很清晰，只是要判断一个数据是否存在即可。

但这里有一个比较重要的前提：**非常庞大的数据**。

### 常规实现

先不考虑这个条件，我们脑海中出现的第一种方案是什么？

我想大多数想到的都是用 `HashMap` 来存放数据，因为它的写入查询的效率都比较高。

写入和判断元素是否存在都有对应的 `API`，所以实现起来也比较简单。

为此我写了一个单测，利用 `HashSet` 来存数据（底层也是 `HashMap` ）；同时为了后面的对比将堆内存写死：

```
-Xms64m -Xmx64m -XX:+PrintHeapAtGC -XX:+HeapDumpOnOutOfMemoryError 
```

为了方便调试加入了 `GC` 日志的打印，以及内存溢出后 `Dump` 内存。

**Java**

```java
@Test
public void hashMapTest(){
	long star = System. currentTimeMillis();
	
    Set<Integer> hashset = new HashSet<>(100);
	for(inti=0;i<100;i++){
		hashset.add(i);
    }
	Assert.assertTrue(hashset. contains(1));
	Assert.assertTrue(hashset. contains(2));
	Assert.assertTrue(hashset.contains(3));
	long end = System. currentTimeMillis();
	System.out.println("执行时间: "+ (end - star));
}

```

当我只写入 100 条数据时自然是没有问题的。

还是在这个基础上，写入 1000W 数据试试：

![img](https://mmbiz.qpic.cn/mmbiz_jpg/csD7FygBVl1whlD22IErSG1dFe3uTGianWcpAznJfmqibIKs5fkWic8VlfY1st0hjUgSDp5mMut9hpGA8JiagdckTA/640?tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)

执行后马上就内存溢出。

![img](https://mmbiz.qpic.cn/mmbiz_jpg/csD7FygBVl1whlD22IErSG1dFe3uTGianZQyKHQA2n1aE8AvT7kHDyoXYpe5JrIU3ZfrcMuUNdNNPwErqrCxZ5w/640?tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)

可见在内存有限的情况下我们不能使用这种方式。

实际情况也是如此；既然要判断一个数据是否存在于集合中，考虑的算法的效率以及准确性肯定是要把数据全部 `load` 到内存中的。

### Bloom Filter

基于上面分析的条件，要实现这个需求最需要解决的是 `如何将庞大的数据load到内存中。`

而我们是否可以换种思路，因为只是需要判断数据是否存在，也不是需要把数据查询出来，所以完全没有必要将真正的数据存放进去。

伟大的科学家们已经帮我们想到了这样的需求。

`BurtonHowardBloom` 在 1970 年提出了一个叫做 `BloomFilter`（中文翻译：布隆过滤）的算法。

它主要就是用于解决判断一个元素是否在一个集合中，但它的优势是只需要占用很小的内存空间以及有着高效的查询效率。

所以在这个场景下在合适不过了。

#### Bloom Filter 原理

下面来分析下它的实现原理。

> 官方的说法是：它是一个保存了很长的二级制向量，同时结合 Hash 函数实现的。

听起来比较绕，但是通过一个图就比较容易理解了。

![img](https://mmbiz.qpic.cn/mmbiz_jpg/csD7FygBVl1whlD22IErSG1dFe3uTGiancITgRibThicbh1uSsD00Az1wNtzodfbykcXgrH41vtSVXIdrjjMfoOSQ/640?tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)

如图所示：

- 首先需要初始化一个二进制的数组，长度设为 L（图中为 8），同时初始值全为 0 。
- 当写入一个 `A1=1000` 的数据时，需要进行 H 次 `hash` 函数的运算（这里为 2 次）；与 HashMap 有点类似，通过算出的 `HashCode` 与 L 取模后定位到 0、2 处，将该处的值设为 1。
- `A2=2000` 也是同理计算后将 `4、7` 位置设为 1。
- 当有一个 `B1=1000` 需要判断是否存在时，也是做两次 Hash 运算，定位到 0、2 处，此时他们的值都为 1 ，所以认为 `B1=1000` 存在于集合中。
- 当有一个 `B2=3000` 时，也是同理。第一次 Hash 定位到 `index=4` 时，数组中的值为 1，所以再进行第二次 Hash 运算，结果定位到 `index=5` 的值为 0，所以认为 `B2=3000` 不存在于集合中。

整个的写入、查询的流程就是这样，汇总起来就是：

> 对写入的数据做 H 次 hash 运算定位到数组中的位置，同时将数据改为 1 。当有数据查询时也是同样的方式定位到数组中。一旦其中的有一位为 **0** 则认为数据**肯定不存在于集合**，否则数据**可能存在于集合中**。

所以布隆过滤有以下几个特点：

1. 只要返回数据不存在，则肯定不存在。
2. 返回数据存在，但只能是大概率存在。
3. 同时不能清除其中的数据。

第一点应该都能理解，重点解释下 2、3 点。

为什么返回存在的数据却是可能存在呢，这其实也和 `HashMap` 类似。

在有限的数组长度中存放大量的数据，即便是再完美的 Hash 算法也会有冲突，所以有可能两个完全不同的 `A、B` 两个数据最后定位到的位置是一模一样的。

这时拿 B 进行查询时那自然就是误报了。

删除数据也是同理，当我把 B 的数据删除时，其实也相当于是把 A 的数据删掉了，这样也会造成后续的误报。

基于以上的 `Hash` 冲突的前提，所以 `BloomFilter` 有一定的误报率，这个误报率和 `Hash`算法的次数 H，以及数组长度 L 都是有关的。

#### 自己实现一个布隆过滤

算法其实很简单不难理解，于是利用 `Go` 实现了一个简单的雏形。

```go
type BloomFilter struct {
	data []int
}

func NewBloomFilter(size int) *BloomFilter {
	return &BloomFilter{
		data: make([]int, size),
	}
}

func (bf *BloomFilter) Add(data string) {
	first := bf.hashcode1(data)
	second := bf.hashcode2(data)
	third := bf.hashcode3(data)

	bf.data[first % len(data)] = 1
	bf.data[second % len(data)] = 1
	bf.data[third % len(data)] = 1
}

func (bf *BloomFilter) MightContain(data string) bool {
	first := bf.hashcode1(data)
	second := bf.hashcode2(data)
	third := bf.hashcode3(data)

	if bf.data[first % len(data)] == 0 {
		return false
	}
	if bf.data[second % len(data)] == 0 {
		return false
	}
	if bf.data[third % len(data)] == 0 {
		return false
	}

	return true
}

func (bf *BloomFilter) hashcode1(data string) int {
	hash := 0
	for i := 0; i < len(data); i++ {
		hash = 33 * hash + int(data[i])
	}

	return int(math.Abs(float64(hash)))
}

func (bf *BloomFilter) hashcode2(data string) int {
	p := 16777619
	hash := 2166136261
	for i := 0; i < len(data); i++ {
		hash = (hash ^ int(data[i])) * p
	}
	hash += hash << 13
	hash ^= hash >> 7
	hash += hash << 3
	hash ^= hash >> 17
	hash += hash << 5
	return int(math.Abs(float64(hash)))
}

func (bf *BloomFilter) hashcode3(data string) int {
	var hash int
	for i := 0; i < len(data); i++ {
		hash += int(data[i])
		hash += hash << 10
		hash ^= hash >> 6
	}
	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15
	return int(math.Abs(float64(hash)))
}
```

1. 首先初始化了一个 int 数组。
2. 写入数据的时候进行三次 `hash` 运算，同时把对应的位置置为 1。
3. 查询时同样的三次 `hash` 运算，取到对应的值，一旦值为 0 ，则认为数据不存在。

实现逻辑其实就和上文描述的一样。

下面来测试一下，同样的参数：

```

```

执行结果如下：

![img](https://mmbiz.qpic.cn/mmbiz_jpg/csD7FygBVl1whlD22IErSG1dFe3uTGianCG3ibyEgpTgT9UnGQduAkSY81RUMwh8UbAEVfe9hiaKDthqa8VLMO1eQ/640?tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)

只花了 3 秒钟就写入了 1000W 的数据同时做出来准确的判断。

------

![img](https://mmbiz.qpic.cn/mmbiz_jpg/csD7FygBVl1whlD22IErSG1dFe3uTGianUe1k90uEYBXBvaIP21v3lRA8jzUVzIGJdzsUOxrTFTiaqXYQich2ibxFw/640?tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)

当让我把数组长度缩小到了 100W 时就出现了一个误报， `400230340` 这个数明明没在集合里，却返回了存在。

这也体现了 `BloomFilter` 的误报率。

我们提高数组长度以及 `hash` 计算次数可以降低误报率，但相应的 `CPU、内存`的消耗就会提高；这就需要根据业务需要自行权衡。

#### Guava 实现

![img](https://mmbiz.qpic.cn/mmbiz_jpg/csD7FygBVl1whlD22IErSG1dFe3uTGian0BXm7ZVdIuTYV5XmNWnowZbibibu6eibxicUZ723g9EL44L56a5cRAUq5Q/640?tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)

刚才的方式虽然实现了功能，也满足了大量数据。但其实观察 `GC` 日志非常频繁，同时老年代也使用了 90%，接近崩溃的边缘。

总的来说就是内存利用率做的不好。

其实 Google Guava 库中也实现了该算法，下面来看看业界权威的实现。

```
-Xms64m -Xmx64m -XX:+PrintHeapAtGC 
```

------

```

```

也是同样写入了 1000W 的数据，执行没有问题。

![img](https://mmbiz.qpic.cn/mmbiz_jpg/csD7FygBVl1whlD22IErSG1dFe3uTGianyaOIGZF4tyWFv18InKB7I9aoKxJYibMq9wukIh9rAojqm9IEfU6luJA/640?tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)

观察 GC 日志会发现没有一次 `fullGC`，同时老年代的使用率很低。和刚才的一对比这里明显的要好上很多，也可以写入更多的数据。

#### 源码分析

那就来看看 `Guava` 它是如何实现的。

构造方法中有两个比较重要的参数，一个是预计存放多少数据，一个是可以接受的误报率。我这里的测试 demo 分别是 1000W 以及 0.01。

![img](https://mmbiz.qpic.cn/mmbiz_jpg/csD7FygBVl1whlD22IErSG1dFe3uTGianFVATmstxOkTWMpqD9xqicFASA2bDF0kz0or7GGFRfGE1GwFO52RNayQ/640?tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)

`Guava` 会通过你预计的数量以及误报率帮你计算出你应当会使用的数组大小 `numBits` 以及需要计算几次 Hash 函数 `numHashFunctions` 。

这个算法计算规则可以参考维基百科。

#### put 写入函数

真正存放数据的 `put` 函数如下：

![img](https://mmbiz.qpic.cn/mmbiz_jpg/csD7FygBVl1whlD22IErSG1dFe3uTGianJfDkAQGicvAytSicibBeXMadsSlMMhvVnCY411PVLZRbuJibN6mpsC2IsQ/640?tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)

- 根据 `murmur3_128` 方法的到一个 128 位长度的 `byte[]`。
- 分别取高低 8 位的到两个 `hash` 值。
- 再根据初始化时的到的执行 `hash` 的次数进行 `hash` 运算。

```
bitsChanged |= bits.set((combinedHash & Long.MAX_VALUE) % bitSize);
```

其实也是 `hash取模`拿到 `index` 后去赋值 1.

重点是 `bits.set()` 方法。

![img](https://mmbiz.qpic.cn/mmbiz_jpg/csD7FygBVl1whlD22IErSG1dFe3uTGianorEMr8Xm8raUu4DibxtnratyiaGTdgWWe0LvyHRJjibEB1SaFhQicGhT4g/640?tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)

其实 set 方法是 `BitArray` 中的一个函数， `BitArray` 就是真正存放数据的底层数据结构。

利用了一个 `long[]data` 来存放数据。

所以 `set()` 时候也是对这个 `data` 做处理。

![img](https://mmbiz.qpic.cn/mmbiz_jpg/csD7FygBVl1whlD22IErSG1dFe3uTGianesDCuV4GOqxq2vkDDrKia3LHiaWr1DO1sR5PHgj7TfCWQEzUrgUicL8KQ/640?tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)

- 在 `set` 之前先通过 `get()` 判断这个数据是否存在于集合中，如果已经存在则直接返回告知客户端写入失败。
- 接下来就是通过位运算进行 `位或赋值`。
- `get()` 方法的计算逻辑和 set 类似，只要判断为 0 就直接返回存在该值。

#### mightContain 是否存在函数

![img](https://mmbiz.qpic.cn/mmbiz_jpg/csD7FygBVl1whlD22IErSG1dFe3uTGianVQP0UPh4vw2QtQ9ZRE23zp6grWiaj4kBnIMtkg4NP2v3IkHNS0lbr8A/640?tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)

前面几步的逻辑都是类似的，只是调用了刚才的 `get()` 方法判断元素是否存在而已。

#### 总结

布隆过滤的应用还是蛮多的，比如数据库、爬虫、防缓存击穿等。

特别是需要精确知道某个数据不存在时做点什么事情就非常适合布隆过滤。

这段时间的研究发现算法也挺有意思的，后续应该会继续分享一些类似的内容。

如果对你有帮助那就分享一下吧。

本问的示例代码参考这里：

https://github.com/crossoverJie/JCSprout