package byte字节跳动

/*
https://www.nowcoder.com/test/question/66b68750cf63406ca1db25d4ad6febbf?pid=8537209&tid=26496289


5
1 2 3 3 5
3
1 2 1
2 4 5
3 5 3

func main() {
	n := 0
	fmt.Scanf("%d", &n)
	users := make([]int, n)
	for i := 0; i < n; i ++ {
		fmt.Scanf("%d", &users[i])
	}
	q := 0
	l, r, k := 0, 0, 0
	fmt.Scanf("%d", &q)
	for i := 0; i < q; i++ {
		fmt.Scanf("%d %d %d", &l, &r, &k)
		fmt.Println(getLikes(l, r, k, users))
	}
}
*/


func getLikes(l, r, k int, users []int) int {
	count := 0
	for x := l-1; x < r; x++ {
		if users[x] == k {
			count++
		}
	}
	return count
}
