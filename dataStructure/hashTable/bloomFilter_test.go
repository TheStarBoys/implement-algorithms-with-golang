package hashTable

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
)

func TestBloomFilter(t *testing.T) {
	var bf *BloomFilter
	size := 100
	dataSize := 100
	Convey("BloomFilter", t, func() {
		Convey("New Bloom Filter", func() {
			bf = NewBloomFilter(size)
			So(len(bf.data), ShouldEqual, size)
		})
		Convey("Add Data and MightContain", func() {
			for i := 0; i < dataSize; i++ {
				bf.Add(strconv.Itoa(i))
			}
			for i := 0; i < dataSize; i++ {
				So(bf.MightContain(strconv.Itoa(i)), ShouldEqual, true)
			}
			So(bf.MightContain("102"), ShouldEqual, false)
			So(bf.MightContain("101"), ShouldEqual, true) // 误报了
		})
	})
}
