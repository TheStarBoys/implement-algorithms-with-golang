package Tencent

import (
	"sort"
)

/*
牛牛和羊羊正在玩一个纸牌游戏。这个游戏一共有n张纸牌, 第i张纸牌上写着数字ai。
牛牛和羊羊轮流抽牌, 牛牛先抽, 每次抽牌他们可以从纸牌堆中任意选择一张抽出, 直到纸牌被抽完。
他们的得分等于他们抽到的纸牌数字总和。
现在假设牛牛和羊羊都采用最优策略, 请你计算出游戏结束后牛牛得分减去羊羊得分等于多少。
 */
/*
func main() {
	n := 0
	fmt.Scanf("%d\n", &n)
	a := make([]int, n)
	for i := 0; i< n; i++ {
		fmt.Scanf("%d",&a[i])
	}
	fmt.Println(getScore(n, a))
}
*/
func getScore(n int, a[]int) int {
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	sum1, sum2 := 0, 0
	for i := 0; i < n; i++ {
		if i % 2 == 0 {
			sum1 += a[i] // 牛牛
		}else {
			sum2 += a[i] // 羊羊
		}
	}
	return sum1-sum2
}