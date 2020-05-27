package stack

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"fmt"
)

func TestSliceStack(t *testing.T) {
	var stack Stack
	stack = NewSliceStack()

	Convey("SliceStack", t, func() {
		Convey("IsEmpty", func() {
			So(stack.IsEmpty(), ShouldEqual, true)
		})
		Convey("Push", func() {
			stack.Push(1)
			So(stack.IsEmpty(), ShouldEqual, false)
			stack.Push(2)
		})
		Convey("Top", func() {
			So(stack.Top().(int), ShouldEqual, 2)
		})
		Convey("Pop", func() {
			So(stack.Pop().(int), ShouldEqual, 2)
			So(stack.Pop().(int), ShouldEqual, 1)
			So(stack.Pop(), ShouldEqual, nil)
			So(stack.Top(), ShouldEqual, nil)
		})
		Convey("Flush", func() {
			stack.Push(1)
			stack.Push(2)
			stack.Push(3)
			stack.Pop()
			So(stack.Top().(int), ShouldEqual, 2)
			So(stack.IsEmpty(), ShouldEqual, false)
			stack.Flush()
			So(stack.IsEmpty(), ShouldEqual, true)
		})
		Convey("String", func() {
			stack.Push(1)
			stack.Push(2)
			stack.Push(3)
			stack.Pop()
			So(fmt.Sprintf("%s", stack), ShouldEqual, "data: [1 2 3], top: 1")
		})
	})
}
