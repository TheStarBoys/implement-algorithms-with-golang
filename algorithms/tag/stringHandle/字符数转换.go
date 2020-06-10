package stringHandle

import (
	"regexp"
	"unicode"
)

// GetStrLength 返回输入的字符串的字数，汉字和中文标点算 2 个字数，英文和其他字符 1 个算 1 个字数
func GetStrLength(str string) int {
	var total int

	reg := regexp.MustCompile("/·|，|。|《|》|‘|’|”|“|；|：|【|】|？|（|）|、/")

	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || reg.Match([]byte(string(r))) {
			total = total + 2
		} else {
			total++
		}
	}

	return total
}