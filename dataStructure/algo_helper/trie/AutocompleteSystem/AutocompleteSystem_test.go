package AutocompleteSystem

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAutocompleteSystem(t *testing.T) {
	Convey("AutocompleteSystem", t, func() {
		as := Constructor([]string{"i love you","island","iroman","i love leetcode"}, []int{5,3,2,2})
		So(hotest(as.findHot()), ShouldEqual, []string{"i love you","island","iroman"})
	})
}
