package leetcode

import "strconv"

func fizzBuzz412_0(n int) []string {
	res := make([]string, n)
	for i := 1; i <= n; i++ {
		if i % 3 == 0 {
			res[i-1] = "Fizz"
		}
		if i % 5 == 0 {
			res[i-1] += "Buzz"
			continue
		}
		if res[i-1] == "" {
			res[i-1] = strconv.Itoa(i)
		}
	}
	return res
}
