package MI

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var maxLength, maxNums = 10000, 10000
func Test_SortStrings(t *testing.T) {
	str := []string{}
	ti := time.Now()
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(maxNums) + 1
	for i := 0; i < length; i++ {
		str = append(str, generateString())
	}

	//str := []string{"a", "aba", "abc", "ab", "d", "ba", "bb"}

	str = SortStrings(str, 0)
	//str = SortStrings2(str) // 2.3161787s

	fmt.Println(str)
	//t.Logf("running time : %s", time.Since(ti))
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