package leetcode

import (
	"fmt"
	"math"
)
// todo 存疑，坐标点[0,0]为障碍的时候，按理来说，不应该可以走动
func robotSim876_0(commands []int, obstacles [][]int) int {
	// obstacles [][]int{}存放的是障碍物坐标[x, y]
	barrier := make(map[string]bool) // 存放障碍物
	var di = 1 // 默认向北朝向
	for _, coordinate := range obstacles {
		barrier[listToStr876_0(coordinate)] = true
	}
	if barrier[listToStr876_0([]int{0, 0})] {
		return 0
	}
	coordinate := []int{0, 0}
	for _, com := range commands {
		if com < 0 { // 转向命令
			di = turn876_0(com, di)
		} else {
			// 尝试往前走，直到遇到障碍物
			for com > 0 && !barrier[listToStr876_0([]int{coordinate[0], coordinate[1]})] {
				com--
				coordinate[0] += dx[di]
				coordinate[1] += dy[di]
			}
			// 如果遇到了障碍物, 退回到没有障碍物的位置
			if com > 0 {
				coordinate[0] -= dx[di]
				coordinate[1] -= dy[di]
			}
		}
	}
	ans := int(math.Pow(float64(coordinate[0]), float64(2)))
	ans += int(math.Pow(float64(coordinate[1]), float64(2)))
	return ans
}

// 对应方向 左，上，右，下

//var dx = []int{-1, 0, 1, 0}
//var dy = []int{0, 1, 0, -1}

// (命令，方向)
func turn876_0(command, di int) int {
	if command == -1 {
		di = (di + 1) % 4
	} else {
		di = (di + 3) % 4
	}
	return di
}

func listToStr876_0(list []int) string {
	return fmt.Sprintf("%q", list)
}
