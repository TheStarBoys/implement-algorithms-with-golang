package leetcode

// recursiveAlgorithm，超时
func new21Game0837_0(N int, K int, W int) float64 {
	if K > N {
		return 0
	}

	return helper0837_0(N, K, W, 0)
}

func helper0837_0(N, K, W, cur int) float64 {
	if cur >= K {
		if N >= cur {
			return 1
		} else {
			return 0
		}
	}
	var res float64
	for i := 1; i <= W; i++ {
		res += helper0837_0(N, K, W, cur + i)
	}

	return res / float64(W)
}


