package leetcode

// 自己套用的LeetCode模板
func openLock752_0(deadends []string, target string) int {
	wheel := []byte{'0','0','0','0'}
	deadSet := make(map[string]bool)
	for _, d := range deadends {
		deadSet[d] = true
	}
	if deadSet[target] || deadSet["0000"] { // 特殊情况
		return -1
	}
	if target == "0000" {
		return 0
	}

	return BFS752_0(wheel, target, deadSet)
}

func BFS752_0(wheel []byte, target string, deadSet map[string]bool) int {
	queue := []string{}
	queue = append(queue, string(wheel))
	used := make(map[string]bool)
	used[string(wheel)] = true
	step := 0
	for len(queue) != 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			cur := queue[0]
			if cur == target {
				return step
			}
			nexts := []string{}
			curWheel := []byte(cur)
			for j := range curWheel {
				o := curWheel[j]
				curWheel[j] = (curWheel[j] - '0' + 1) % 10 + '0'
				nexts = append(nexts, string(curWheel))
				curWheel[j] = o

				curWheel[j] = (curWheel[j] - '0' + 9) % 10 + '0'
				nexts = append(nexts, string(curWheel))
				curWheel[j] = o
			}
			for _, next := range nexts {
				if !deadSet[next] && !used[next] {
					queue = append(queue, next)
					used[next] = true
				}
			}
			queue = queue[1:]
		}
		step++
	}
	return -1
}

// 大佬解法
func openLock752_1(deadends []string, target string) int {
	deadSet := make(map[string]bool) // 使用hashSet快速查找死亡数字
	for _, s := range deadends {     // 可以快速查看每个数字是否是死亡数字
		deadSet[s] = true
	}

	if deadSet[target] || deadSet["0000"] { // 特殊情况讨论
		return -1
	}
	if target == "0000" {
		return 0
	}

	// BFS
	queue := make([]string, 0)      // 申请一个队列,存储当前密码锁的状态
	queue = append(queue, "0000")   // 添加锁的初始状态"0000"
	visited := make(map[string]int) // 使用HashMap记录是否遍历过,还要记录到达某个状态的步数
	visited["0000"] = 0             // 初始"0000",所需要到达的步数为0
	for len(queue) > 0 {            // 只要队列不为空就继续执行循环
		curs := queue[0]         // 取出当前的状态
		queue = queue[1:]        // 从curs状态出发,向上拨一下向下拨一下,达到其他8种状态
		curarray := []rune(curs) // 遍历curs中每一位,对每一位进行遍历

		var nexts []string       // 从当前状态出发,可以到达的其他状态
		for i := 0; i < 4; i++ { // 对每一位进行相应的变动
			o := curarray[i]                           // 保存原始的字符
			curarray[i] = (curarray[i]-'0'+1)%10 + '0' // 0-9字符减'0'变为整型加'0'转为字符
			nexts = append(nexts, string(curarray))    // 新的状态(向后拨一位)，添加到nexts
			curarray[i] = o                            // 归位原始字符

			curarray[i] = (curarray[i]-'0'+9)%10 + '0' // 向前拨一位,避免负数,+9再对10求余
			nexts = append(nexts, string(curarray))    // 新的状态(向前拨一位),添加到nexts
			curarray[i] = o                            // 归位原始字符
		}

		for _, next := range nexts { // 从nexts数组中取出一个个下一个状态
			// 对每个next检查,1. 不能是死亡数组,2. 之前没有入队,没有访问过
			if _, ok := visited[next]; !deadSet[next] && !ok {
				queue = append(queue, next)       // next放入队列
				visited[next] = visited[curs] + 1 // 对应达到这个next需要多少步
				if next == target {               // 如果已经达到目标
					return visited[next]
				}

			}

		}
	}
	return -1
}
/*
作者：custergo
链接：https://leetcode-cn.com/problems/open-the-lock/solution/custerxue-xi-bi-ji-tu-bfs-by-custergo-2/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/