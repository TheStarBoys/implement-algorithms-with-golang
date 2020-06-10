# 字符串匹配

## 简介

在开始讲解这个算法之前，我先定义两个概念，方便我后面讲解。它们分别是主串和模式串。

这俩概念很好理解，我举个例子你就懂了。比方说，我们在字符串 A 中查找字符串 B，那字符串 A 就是主串，字符串 B 就是模式串。我们把主串的长度记作 n，模式串的长度记作 m。因为我们是在主串中查找模式串，所以 n>m。 

## BF算法

### 简介

BF 算法中的 BF 是 Brute Force 的缩写，中文叫作暴力匹配算法，也叫朴素匹配算法。从名字可以看出，这种算法的字符串匹配方式很“暴力”，当然也就会比较简单、好懂，但相应的性能也不高。 

作为最简单、最暴力的字符串匹配算法，BF 算法的思想可以用一句话来概括，那就是，**我们在主串中，检查起始位置分别是 0、1、2…n-m 且长度为 m 的 n-m+1 个子串，看有没有跟模式串匹配的**。 

### 代码实现

```go
// Brute Force
// 时间复杂度：O(n*m) n 为主串长度，m 为模式串长度
func BFSearch(str, sub string) int {
	n, m := len(str), len(sub)
	Flag:
	for i := 0; i < n - m + 1; i++ {
		// cur 为主串和模式串对齐的下标
		for cur, j := i, 0; j < m; cur, j = cur+1, j+1 {
			if str[cur] != sub[j] {
				continue Flag
			}
		}
		return i
	}

	return -1
}
```

### 复杂度分析

**时间复杂度**：`O(n * m)`

- 极端情况下，每次将比对 `n - m + 1` 次，每次将比对 `m` 个字符
- 实际软件开发中，大部分情况下，模式串和主串长度不会太长。而且每次模式串与主串中的子串匹配的时候，当中途不能匹配的时候，就可以停止了，不需要比对 `m` 个字符，所以统计意义上，时间复杂度比 `O(n * m)` 要低很多

**空间复杂度**：`O(1)`

## RK算法

### 简介

RK 算法的全称叫 Rabin-Karp 算法，是由它的两位发明者 Rabin 和 Karp 的名字来命名的。这个算法理解起来也不是很难。我个人觉得，它其实就是刚刚讲的 BF 算法的升级版。 

每次检查主串与子串是否匹配，需要依次比对每个字符，所以 BF 算法的时间复杂度就比较高，是 `O(n*m)`。我们对朴素的字符串匹配算法稍加改造，引入哈希算法，时间复杂度立刻就会降低。 

RK 算法的思路是这样的：我们通过哈希算法对主串中的 n-m+1 个子串分别求哈希值，然后逐个与模式串的哈希值比较大小。如果某个子串的哈希值与模式串相等，那就说明对应的子串和模式串匹配了（这里先不考虑哈希冲突的问题，后面我们会讲到）。因为哈希值是一个数字，数字之间比较是否相等是非常快速的，所以模式串和子串比较的效率就提高了。 

不过，通过哈希算法计算子串的哈希值的时候，我们需要遍历子串中的每个字符。尽管模式串与子串比较的效率提高了，但是，算法整体的效率并没有提高。有没有方法可以提高哈希算法计算子串哈希值的效率呢？

这就需要哈希算法设计的非常有技巧了。我们假设要匹配的字符串的字符集中只包含 K 个字符，我们可以用一个 K 进制数来表示一个子串，这个 K 进制数转化成十进制数，作为子串的哈希值。表述起来有点抽象，我举了一个例子，看完你应该就能懂了。

比如要处理的字符串只包含 a～z 这 26 个小写字母，那我们就用二十六进制来表示一个字符串。我们把 a～z 这 26 个字符映射到 0～25 这 26 个数字，a 就表示 0，b 就表示 1，以此类推，z 表示 25。

在十进制的表示法中，一个数字的值是通过下面的方式计算出来的。对应到二十六进制，一个包含 a 到 z 这 26 个字符的字符串，计算哈希的时候，我们只需要把进位从 10 改成 26 就可以。 

 ![img](https://static001.geekbang.org/resource/image/d5/04/d5c1cb11d9fc97d0b28513ba7495ab04.jpg) 

这个哈希算法你应该看懂了吧？现在，为了方便解释，在下面的讲解中，我假设字符串中只包含 a～z 这 26 个小写字符，我们用二十六进制来表示一个字符串，对应的哈希值就是二十六进制数转化成十进制的结果。

这种哈希算法有一个特点，在主串中，相邻两个子串的哈希值的计算公式有一定关系。我这有个个例子，你先找一下规律，再来看我后面的讲解。 

 ![img](https://static001.geekbang.org/resource/image/f9/f5/f99c16f2f899d19935567102c59661f5.jpg) 

从这里例子中，我们很容易就能得出这样的规律：相邻两个子串 s[i-1]和 s[i]（i 表示子串在主串中的起始位置，子串的长度都为 m），对应的哈希值计算公式有交集，也就是说，我们可以使用 s[i-1]的哈希值很快的计算出 s[i]的哈希值。如果用公式表示的话，就是下面这个样子： 

 ![img](https://static001.geekbang.org/resource/image/c4/9c/c47b092408ebfddfa96268037d53aa9c.jpg) 

不过，这里有一个小细节需要注意，那就是 26^(m-1) 这部分的计算，我们可以通过查表的方法来提高效率。我们事先计算好 26^0、26^1、26^2……26^(m-1)，并且存储在一个长度为 m 的数组中，公式中的“次方”就对应数组的下标。当我们需要计算 26 的 x 次方的时候，就可以从数组的下标为 x 的位置取值，直接使用，省去了计算的时间。 

 ![img](https://static001.geekbang.org/resource/image/22/2f/224b899c6e82ec54594e2683acc4552f.jpg) 

### 代码实现

```go
// 时间复杂度：O(n)
func RKSearch(str, sub string) int {
	if len(str) < len(sub) {
		return -1
	}
	var (
		m = len(sub)
		subCode = hashCode3(sub)
		l = 0
		// code 是滑动窗口中子串的哈希值
		code = hashCode3(str[:m])
	)
	// 利用滑动窗口，可以用O(1)的时间从 h[i-1] 递推得到 h[i]，i 为子串起始下标
	for r := m; ; r++ {
		if code == subCode {
			// 如果哈希算法是允许有哈希冲突的，此处还需要额外比较子串跟模式串是否相等
			//if str[l : r] == sub {
			//	return l
			//}
			return l
		}
		if r >= len(str) {
			break
		}
		//code -= int(str[l] - 'a' + 1) * int(math.Pow(26, float64(m-1)))
		code -= int(str[l] - 'a' + 1) * pow(m-1)
		l++
		code *= 26
		code += int(str[r] - 'a' + 1)
	}

	return -1
}

// 基于 进制转换的思想，a ~ z 一共26个字母，把它当成一个26进制，"不存在哈希冲突"(其实a从0开始导致会有一定问题)
// 这种方式其值很容易就超出了整型的范围
// 有一定问题，字符串 "abc" 跟 "bc" 结果其实是一样的，都是28
func hashCode1(str string) int {
	sum := 0
	i := 0
	for j := len(str) - 1; j >= 0; j-- {
		sum += int(str[j] - 'a') * int(math.Pow(26, float64(i)))
		i++
	}

	return sum
}

// 基于第一种方式，将 a 记为 1 而不是0
func hashCode2(str string) int {
	sum := 0
	i := 0
	for j := len(str) - 1; j >= 0; j-- {
		sum += int(str[j] - 'a' + 1) * int(math.Pow(26, float64(i)))
		i++
	}

	return sum
}

// 查表法优化 math.Pow()
func hashCode3(str string) int {
	sum := 0
	i := 0
	for j := len(str) - 1; j >= 0; j-- {
		sum += int(str[j] - 'a' + 1) * pow(i)
		i++
	}

	return sum
}

// pow 可以用查表法优化
var powTable =  []int{
	0: 1,
	1: 26,
	2: 676,
	3: 17576,
	4: 456976,
	5: 11881376,
	6: 308915776,
	7: 8031810176,
	8: 208827064576,
	9: 5429503678976,
	10: 141167095653376,
	11: 3670344486987776,
	12: 95428956661682176,
	13: 2481152873203736576,
}
// pow returns 26 ^ x
func pow(x int) int {
	if x < 0 {
		panic("x less than zero")
	}
	if x > 13 {
		panic("exceed int range")
	}

	return powTable[x]
}
```

### 复杂度分析

**时间复杂度**：`O(n)`

-  如果不存在哈希冲突
  - 整个 RK 算法包含两部分 
    -  计算子串哈希值 
      - 通过设计特殊的哈希算法，只需要扫描一遍主串就能计算出所有子串的哈希值了，所以这部分的时间复杂度是 `O(n) `
    -  模式串哈希值与子串哈希值之间的比较 
      - 模式串哈希值与每个子串哈希值之间的比较的时间复杂度是 `O(1)`，总共需要比较 `n-m+1` 个子串的哈希值，所以，这部分的时间复杂度也是 `O(n)`。
  - 所以，RK 算法整体的时间复杂度就是 `O(n)`。 
- 如果存在哈希冲突
  - 哈希算法的冲突概率要相对控制得低一些，如果存在大量冲突，就会导致 RK 算法的时间复杂度退化，效率下降。极端情况下，如果存在大量的冲突，每次都要再对比子串和模式串本身，那时间复杂度就会退化成 `O(n*m)`。但也不要太悲观，一般情况下，冲突不会很多，RK 算法的效率还是比 BF 算法高的 

**空间复杂度**：`O(1)`

- 尽管进行 `math.Pow()` 运算时转换成了查表法获取值，但powTable大小并不随数据规模增大而增大，所以仍然是常量级的空间。

## BM算法

### 简介

BM（Boyer-Moore）算法。它是一种非常高效的字符串匹配算法，有实验统计，它的性能是著名的 KMP 算法的 3 到 4 倍。 

### 代码实现

```go
const (
	_SIZE = 256
)

func BMSearch(mainStr, patternStr string) int {
	return bm(mainStr, patternStr, len(mainStr), len(patternStr))
}

// n是主串的长度，m是模式串的长度
func bm(mainStr, patternStr string, n, m int) int {
	bc := make([]int, _SIZE) // 记录模式串中每个字符最后出现的下标
	generaetBC(patternStr, m, bc) // 构建坏字符哈希表
	suffix := make([]int, m)
	prefix := make([]bool, m)
	generateGS(patternStr, m, suffix, prefix)
	i := 0 // i 表示主串与模式串对其的第一个字符
	for i <= n - m {
		j := m - 1
		for ; j >= 0; j-- { // 模式串从后往前匹配
			if mainStr[i+j] != patternStr[j] { break } // 坏字符对应模式串中的下标是j
		}
		if j < 0 { return i } // 匹配成功，返回主串与模式串第一个匹配的字符的位置
		// x 为坏字符规则移动距离，该值可能为负数，因为bc数组中对应的是坏字符最后出现的位置
		// 这个下标不一定在 j 之前
		x := j - bc[int(mainStr[i+j])]
		y := 0 // 好后缀字符移动距离
		if j < m - 1 { // 如果有好后缀的话
			y = moveByGS(j, m, suffix, prefix)
		}
		// 从坏字符规则和好后缀规则中选择移动距离最大的值 来进行移动
		i += max(x, y)
	}

	return -1
}

// generaetBC 记录模式串中每个字符最后出现的下标
// 假设字符集不大，数组badCharacter下标对应ASCII码值，存储这个字符在模式串中出现的位置
func generaetBC(patternStr string, m int, badCharacter []int) {
	// 初始化 badCharacter
	for i := 0; i < _SIZE; i++ {
		badCharacter[i] = -1
	}
	for i := 0; i < m; i++ {
		ascii := int(patternStr[i])
		badCharacter[ascii] = i
	}
}

// generateGS 生成好后缀相关的切片
// m 表示模式串长度, suffix, prefix 提前申请好了长度
// suffix 下标k表示后缀子串长度，其值是在模式串中跟好后缀相匹配的子串的起始下标值
// prefix 来记录模式串的后缀子串是否能匹配模式串的前缀子串
func generateGS(patternStr string, m int, suffix []int, prefix []bool) {
	for i := 0; i < m; i++ { // 初始化
		suffix[i] = -1
		prefix[i] = false
	}
	for i := 0; i < m - 1; i++ { // patternStr [0, i]
		j := i
		k := 0 // 公共后缀子串长度
		for j >= 0 && patternStr[j] == patternStr[m-1-k] { // 与 patternStr [0, m-1] 求公共后缀子串
			j--
			k++
			suffix[k] = j + 1 // j+1 表示公共后缀子串在 patternStr [0, i] 中的起始下标
		}
		if j == -1 { prefix[k] = true } // 如果公共后缀子串也是模式串的前缀子串
	}
}

// j 表示坏字符对应的模式串中的字符下标，m 表示模式串长度
func moveByGS(j, m int, suffix []int, prefix []bool) int {
	k := m - 1 - j // 好后缀长度
	if suffix[k] != -1 { // 长度为 k 时有可匹配的子串
		return j - suffix[k] + 1
	}
	for r := j + 2; r <= m - 1; r++ {
		// 长度为 m - r 的有可匹配的前缀子串
		if prefix[m-r] {
			return r
		}
	}

	return m
}

func max(x, y int) int {
	if x >= y {
		return x
	}

	return y
}
```

### 复杂度分析

**时间复杂度**：

- 实际上，BM 算法的时间复杂度分析起来是非常复杂，这篇论文“[A new proof of the linearity of the Boyer-Moore string searching algorithm](http://dl.acm.org/citation.cfm?id=1382431.1382552)”证明了在最坏情况下，BM 算法的比较次数上限是 5n。
- 这篇论文“[Tight bounds on the complexity of the Boyer-Moore string matching algorithm](https://dl.acm.org/doi/10.5555/127787.127830)”证明了在最坏情况下，BM 算法的比较次数上限是 3n。

**空间复杂度**：

- 整个算法用到了额外的 3 个数组，其中 bc 数组的大小跟字符集大小有关，suffix 数组和 prefix 数组的大小跟模式串长度 m 有关。 

## KMP算法

### 简介

KMP 算法是根据三位作者（D.E.Knuth，J.H.Morris 和 V.R.Pratt）的名字来命名的，算法的全称是 Knuth Morris Pratt 算法，简称为 KMP 算法。 

在模式串和主串匹配的过程中，把不能匹配的那个字符仍然叫作**坏字符**，把已经匹配的那段字符串叫作**好前缀**。 

为了表述起来方便，我把好前缀的所有后缀子串中，最长的可匹配前缀子串的那个后缀子串，叫作**最长可匹配后缀子串**；对应的前缀子串，叫作**最长可匹配前缀子串**。 

KMP 算法也可以提前构建一个数组，用来存储模式串中每个前缀（这些前缀都有可能是好前缀）的最长可匹配前缀子串的结尾字符下标。我们把这个数组定义为 **next 数组**，很多书中还给这个数组起了一个名字，叫**失效函数**（failure function）。 

 ![img](https://static001.geekbang.org/resource/image/9e/ad/9e59c0973ffb965abdd3be5eafb492ad.jpg) 

数组的下标是每个前缀结尾字符下标，数组的值是这个前缀的最长可以匹配前缀子串的结尾字符下标。这句话有点拗口，我举了一个例子，你一看应该就懂了。 

 ![img](https://static001.geekbang.org/resource/image/16/a8/1661d37cb190cb83d713749ff9feaea8.jpg) 

### 代码实现

```go
func KMPSearch(mainStr, patternStr string) int {
	return kmp(mainStr, patternStr, len(mainStr), len(patternStr))
}

// n, m分别为主串和模式串长度
func kmp(mainStr, patternStr string, n, m int) int {
	next := getNexts(patternStr, m)
	j := 0
	for i := 0; i < n; i++ {
		// 当有好前缀 并且 此时主串和模式串出现坏字符时
		for j > 0 && mainStr[i] != patternStr[j] {
			// 从next数组中获取好前缀[0, j-1] 可匹配的最长前缀子串结尾下标
			j = next[j - 1] + 1
		}
		if mainStr[i] == patternStr[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}

	return -1
}

// next数组，前缀结束下标 : 最长可匹配前缀子串结尾字符下标
// 例如：模式串：ababacd
// 模式串前缀（好前缀候选） 前缀结尾字符下标 最长可匹配...下标
// a							0 		: 		-1
// ab							1 		: 		-1
// aba							2 		: 		 0
// abab							3 		: 		 1
// ababa						4 		: 		 2
// ababac						5 		: 		-1

// 以其中 abab 好前缀候选为例：
// 后缀	前缀		可匹配结尾下标
// b	a		-1 // 自己不能跟自己匹配
// ab	ab		 1
// bab	aba		-1
// abab	abab	-1 // 自己不能跟自己匹配
// 最长结尾下标 --> 1
func getNexts(patternStr string, m int) []int {
	next := make([]int, m)
	next[0] = -1
	k := -1
	// i 为好前缀候选结尾字符下标, next[m - 1] 的值其实肯定为 -1
	for i := 1; i < m; i++ {
		for k != -1 && patternStr[k + 1] != patternStr[i] {
			k = next[k]
		}
		// 如果相等，说明子串 [0, k+1] 就是 模式串[0, i]的最长可匹配前缀子串
		if patternStr[k + 1] == patternStr[i] {
			k++
		}
		next[i] = k
	}

	return next
}
```

### 复杂度分析

**时间复杂度**：`O(n+m)`

KMP 算法包含两部分 

- 第一部分：构建 next 数组 
  - 我们可以找一些参照变量，i 和 k。i 从 1 开始一直增加到 m，而 k 并不是每次 for 循环都会增加，所以，k 累积增加的值肯定小于 m。
  - 而 while 循环里 `k=next[k]`，实际上是在减小 k 的值，k 累积都没有增加超过 m，所以 while 循环里面 `k=next[k]` 总的执行次数也不可能超过 m。因此，next 数组计算的时间复杂度是 `O(m)`。 
- 第二部分：借助 next 数组匹配 
  - i 从 0 循环增长到 n-1，j 的增长量不可能超过 i，所以肯定小于 n。而 while 循环中的那条语句 `j=next[j-1]+1`，不会让 j 增长的
  - 那有没有可能让 j 不变呢？也没有可能。因为 `next[j-1]` 的值肯定小于 j-1，所以 while 循环中的这条语句实际上也是在让 j 的值减少。
  - 而 j 总共增长的量都不会超过 n，那减少的量也不可能超过 n，所以 while 循环中的这条语句总的执行次数也不会超过 n，所以这部分的时间复杂度是 `O(n)`。 

**空间复杂度**：`O(m)`

KMP 算法只需要一个额外的 next 数组，数组的大小跟模式串相同。所以空间复杂度是 `O(m)`，m 表示模式串的长度。 

## Sunday算法

### 代码实现

```go
// sunday算法，字符串模式匹配
func Sunday(str, sub string) int{
	cur := 0 // 指向当前str和sub对齐的位置
	stri, subi := 0, 0 // 双指针，分别指向str, sub
	Flag :
	for stri < len(str) && subi < len(sub) && cur+len(sub) <= len(str){
		if str[stri] != sub[subi] {
			// 匹配失败，关注src中参加匹配的最末位字符的下一个字符
			i := 0
			for i = len(sub)-1; i >= 0; i-- {
				if cur+len(sub) < len(str) && sub[i] == str[cur+len(sub)] {
					// 末位的下一个字符在sub中
					cur = cur + (len(sub) - i)
					stri = cur
					subi = 0
					continue Flag
				}
			}
			if i < 0 {
				// 末位的下一个字符不在sub中
				cur = cur + len(sub) + 1
				stri = cur
				subi = 0
				continue
			}
		}
		stri++
		subi++
	}
	if subi >= len(sub) {
		return cur
	}
	return -1
}
```



## 引用

> [数据结构与算法之美](https://time.geekbang.org/column/intro/100017301)