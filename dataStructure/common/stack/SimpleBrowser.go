package stack

import "fmt"

type Browser struct {
	forwardStack Stack
	backStack Stack
}

func NewBrowser() *Browser {
	return &Browser{
		forwardStack: NewSliceStack(),
		backStack: NewLinkedListStack(),
	}
}

func (b *Browser) OpenWeb(url string) {
	fmt.Println("Open web:", url)
	b.backStack.Push(url)
	b.forwardStack.Flush()
}

func (b *Browser) CanBack() bool {
	return b.backStack.IsEmpty() == false
}

func (b *Browser) Back() (url string) {
	if b.backStack.IsEmpty() {
		return ""
	}

	v := b.backStack.Pop()
	b.forwardStack.Push(v)

	return b.Current()
}

func (b *Browser) CanForward() bool {
	return b.forwardStack.IsEmpty() == false
}

func (b *Browser) Forward() (url string) {
	if b.forwardStack.IsEmpty() {
		return ""
	}

	v := b.forwardStack.Pop()
	b.backStack.Push(v)

	return b.Current()
}

func (b *Browser) Current() (url string) {
	if b.backStack.IsEmpty() {
		return ""
	}

	return b.backStack.Top().(string)
}
