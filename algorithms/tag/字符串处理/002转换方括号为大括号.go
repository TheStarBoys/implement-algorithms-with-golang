package 字符串处理

import "strings"

func convertToBig(str string) string {
	str = strings.ReplaceAll(str, "[", "{")
	str = strings.ReplaceAll(str, "]", "}")
	return str
}