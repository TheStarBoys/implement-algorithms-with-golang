package stringHandle

import (
	"testing"
	"fmt"
)

func TestBytes2Unicode(t *testing.T) {
	runes := Bytes2Utf8([]byte("hello, 世界"))
	for i := range runes {
		fmt.Printf("%c", runes[i])
	}
	fmt.Println()
	//fmt.Printf("%#v\n", []byte("hello, 世界"))
	//for _, b := range []byte("hello, 世界") {
	//	fmt.Println(hex.EncodeToString([]byte{b}))
	//}
	//fmt.Println(Str2BytesCode("hello, 世界"))
	//fmt.Println("\x68\x65\x6c\x6c\x6f\x2c\x20\xe4\xe7")
}

