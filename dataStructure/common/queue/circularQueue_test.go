package queue

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"

)
func TestCircularQueue(t *testing.T) {
	queue := NewCircularQueue(5)

	Convey("Queue", t, func() {
		Convey("New Queue", func() {
			So(queue.IsEmpty(), ShouldEqual, true)
			So(queue.Front(), ShouldEqual, nil)
			So(queue.Rear(), ShouldEqual, nil)
			So(queue.IsFull(), ShouldEqual, false)
		})
		Convey("EnQueue", func() {
			So(queue.Rear(), ShouldEqual, nil)
			So(queue.EnQueue(1), ShouldEqual, true)
			So(queue.Rear(), ShouldEqual, 1)
			So(queue.EnQueue(2), ShouldEqual, true)
			So(queue.Rear(), ShouldEqual, 2)
			So(queue.EnQueue(3), ShouldEqual, true)
			So(queue.Rear(), ShouldEqual, 3)
			So(queue.EnQueue(4), ShouldEqual, true)
			So(queue.Rear(), ShouldEqual, 4)
			So(queue.EnQueue(5), ShouldEqual, true)
			So(queue.Rear(), ShouldEqual, 5)
			So(queue.EnQueue(6), ShouldEqual, false)
			So(queue.Rear(), ShouldEqual, 5)
		})
		Convey("DeQueue", func() {
			So(queue.Front(), ShouldEqual, 1)
			So(queue.DeQueue(), ShouldEqual, true)
			So(queue.Front(), ShouldEqual, 2)
			So(queue.DeQueue(), ShouldEqual, true)
			So(queue.Front(), ShouldEqual, 3)
			So(queue.DeQueue(), ShouldEqual, true)
			So(queue.Front(), ShouldEqual, 4)
			So(queue.DeQueue(), ShouldEqual, true)
			So(queue.Front(), ShouldEqual, 5)
			So(queue.DeQueue(), ShouldEqual, true)
			So(queue.Front(), ShouldEqual, nil)
			So(queue.DeQueue(), ShouldEqual, false)
		})

	})
}