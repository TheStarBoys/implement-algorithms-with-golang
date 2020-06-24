package offer

func reverseLeftWords(s string, n int) string {
	bts := []byte(s)
	bts = append(bts[n:], bts[:n]...)

	return string(bts)
}