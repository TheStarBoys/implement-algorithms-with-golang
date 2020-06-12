package MI

import (
	"math/rand"
	"testing"
	"time"
)

var maxLength, maxNums = 10000, 10000
func Test_SortStrings(t *testing.T) {
	strs := []string{}
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(maxNums) + 1
	for i := 0; i < length; i++ {
		strs = append(strs, generateString())
	}

	//str := []string{"a", "aba", "abc", "ab", "d", "ba", "bb"}

	strs = SortStrings(strs, 0)

	for i := 1; i < len(strs); i++ {
		if strs[i-1] > strs[i] {
			t.Errorf("expect: %s <= %s", strs[i-1], strs[i])
		}
	}
}


func generateString() string {
	length := rand.Intn(maxLength) + 1
	buf := []byte{}
	for i := 0; i < length; i++ {
		b := byte(rand.Intn(26) + 'a')
		buf = append(buf, b)
	}
	return string(buf)
}