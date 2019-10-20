package MI

import (
	"math/rand"
	"testing"
	"time"
)

func Test_SortStrings(t *testing.T) {
	str := []string{}
	ti := time.Now()
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(100000) + 1
	for i := 0; i < length; i++ {
		str = append(str, generateString())
	}
	str = SortStrings(str, 0)
	t.Logf("running time : %s", time.Since(ti))
}


func generateString() string {
	length := rand.Intn(1000) + 1
	buf := []byte{}
	for i := 0; i < length; i++ {
		b := byte(rand.Intn(26) + 'a')
		buf = append(buf, b)
	}
	return string(buf)
}