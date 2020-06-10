package ac_automation

import (
	"testing"
	"fmt"
)

func TestACAutomation(t *testing.T) {
	ac := New()
	ac.Insert([]rune("abcd"))
	ac.Insert([]rune("bc"))
	ac.Insert([]rune("bcd"))
	ac.Insert([]rune("c"))
	if ac.Search([]rune("abcd")) == false {
		t.Errorf("expect: true")
	}
	if ac.Search([]rune("bc")) == false {
		t.Errorf("expect: true")
	}
	if ac.Search([]rune("bcd")) == false {
		t.Errorf("expect: true")
	}
	if ac.Search([]rune("c")) == false {
		t.Errorf("expect: true")
	}
	if ac.Search([]rune("ab")) == true {
		t.Errorf("expect: false")
	}

	ac.Insert([]rune("我草"))
	ac.Insert([]rune("尼玛"))
	ac.Insert([]rune("狗东西"))
	ac.Insert([]rune("贱人"))
	ac.Insert([]rune("屎"))
	ac.Insert([]rune("操"))
	ac.BuildFailurePointer()
	fmt.Println(ac.Match([]rune("我草狗东西，你个贱人，吃屎吧你，操")))
	str := "我草狗东西，你个贱人，吃屎吧你，操"

	s, ok := ac.Replace([]rune(str), '*')
	if !ok {
		t.Errorf("expect: true")
	}
	fmt.Printf("%s ==> %s\n", str, string(s))

	s, ok = ac.ReplaceWithN([]rune(str), '*', 3)
	if !ok {
		t.Errorf("expect: true")
	}
	fmt.Printf("%s ==> %s\n", str, string(s))

	str = "你狗东西，你个贱人，吃屎吧你，md"
	s, ok = ac.ReplaceWithN([]rune(str), '*', 2)
	if !ok {
		t.Errorf("expect: true")
	}
	fmt.Printf("%s ==> %s\n", str, string(s))

	if ac.Search([]rune("abcd")) == false {
		t.Errorf("expect: true")
	}
	if ac.Search([]rune("bc")) == false {
		t.Errorf("expect: true")
	}
	if ac.Search([]rune("bcd")) == false {
		t.Errorf("expect: true")
	}
	if ac.Search([]rune("c")) == false {
		t.Errorf("expect: true")
	}
	if ac.Search([]rune("ab")) == true {
		t.Errorf("expect: false")
	}
}
