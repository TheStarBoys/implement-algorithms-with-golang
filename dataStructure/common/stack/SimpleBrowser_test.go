package stack

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBrowser(t *testing.T) {
	browser := NewBrowser()
	Convey("Browser", t, func() {
		browser.OpenWeb("www.qq.com")
		So(browser.CanBack(), ShouldEqual, true)
		So(browser.CanForward(), ShouldEqual, false)
		So(browser.Current(), ShouldEqual, "www.qq.com")

		browser.OpenWeb("www.baidu.com")
		So(browser.CanBack(), ShouldEqual, true)
		So(browser.CanForward(), ShouldEqual, false)
		So(browser.Current(), ShouldEqual, "www.baidu.com")

		browser.OpenWeb("www.tencent.com")
		So(browser.CanBack(), ShouldEqual, true)
		So(browser.CanForward(), ShouldEqual, false)
		So(browser.Current(), ShouldEqual, "www.tencent.com")

		So(browser.Back(), ShouldEqual, "www.baidu.com")
		So(browser.Back(), ShouldEqual, "www.qq.com")
		So(browser.Forward(), ShouldEqual, "www.baidu.com")

		browser.OpenWeb("www.sina.com")
		So(browser.CanBack(), ShouldEqual, true)
		So(browser.CanForward(), ShouldEqual, false)
		So(browser.Current(), ShouldEqual, "www.sina.com")
		So(browser.Forward(), ShouldEqual, "")

		So(browser.Back(), ShouldEqual, "www.baidu.com")
		So(browser.Back(), ShouldEqual, "www.qq.com")
		So(browser.Back(), ShouldEqual, "")
	})
}
