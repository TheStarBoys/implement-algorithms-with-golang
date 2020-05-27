package hashTable

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
)

func TestBloomFilter_GoConvey(t *testing.T) {
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
			So(bf.MightContain("10000000045871"), ShouldEqual, false)
			So(bf.MightContain("101"), ShouldEqual, true) // 误报了
		})
	})
}


func TestBloomFilter(t *testing.T) {
	// 百万数据：0.14s
	// 千万数据：1.48s
	// 一亿数据：15.15s
	var bf *BloomFilter
	size := 1000000
	dataSize := 1000000
	bf = NewBloomFilter(size)
	if len(bf.data) != size {
		t.Errorf("length not equal")
	}
	for i := 0; i < dataSize; i++ {
		bf.Add(strconv.Itoa(i))
	}
	for i := 0; i < dataSize; i++ {
		if actual := bf.MightContain(strconv.Itoa(i)); actual != true {
			t.Errorf("MightContain, actual: %v, expect: %v\n", actual, true)
		}
	}
	if actual := bf.MightContain("10000000045871"); actual != false {
		t.Errorf("MightContain, actual: %v, expect: %v\n", actual, false)
	}
}
