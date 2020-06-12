package stringMatching

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