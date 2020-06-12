package Tencent

/*
package main
import (
	"fmt"
	"sort"
)
type task struct {
	time int
	level int
}
type machine struct {
	time int
	level int
}

测试用例：
4 3
435 76
164 87
759 68
201 16
587 38
862 57
753 46

func main() {
	n, m := 0, 0
	fmt.Scanf("%d %d", &n, &m)
	macs := make([]machine, n)
	tsks := make([]task, m)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d", &macs[i].time, &macs[i].level)
	}
	for j := 0; j < m; j++ {
		fmt.Scanf("%d %d", &tsks[j].time, &tsks[j].level)
	}
	fmt.Println(arrange(n, m, macs, tsks))
}
// 通过率50% 不及格
func arrange(n, m int, macs []machine, tsks []task) (int, int) {
	sort.Slice(macs, func(i, j int) bool {
		if macs[i].time == macs[j].time {
			return macs[i].level > macs[j].level
		}
		return macs[i].time > macs[j].time
	})
	sort.Slice(tsks, func(i, j int) bool {
		if tsks[i].time == tsks[j].time {
			return tsks[i].level > tsks[j].level
		}
		return tsks[i].time > tsks[j].time
	})
	in, count := 0, 0
	used := make([]int, n)
	for i := range tsks {
		for j := range macs {
			if used[j] == 1 {
				continue
			}
			if macs[j].level >= tsks[i].level && macs[j].time >= tsks[i].time{
				in += income(tsks[i].time, tsks[i].level)
				count++
				used[j] = 1
				break
			}
		}
	}
	return count, in
}

func income(x, y int) int {
	return 200 * x + 3 * y
}

*/