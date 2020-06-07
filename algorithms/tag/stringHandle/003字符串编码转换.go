package stringHandle

import (
	"unicode/utf8"
)


// 字节转utf8
func Bytes2Utf8(data []byte) (res []rune) {
	for len(data) > 0 {
		r, n := utf8.DecodeRune(data)
		res = append(res, r)
		data = data[n:]
	}

	return
}

//func Str2BytesCode(str string) string {
//	res := ""
//	for i := range str {
//		res += "\\x" + hex.EncodeToString([]byte{str[i]})
//	}
//
//	return res
//}
