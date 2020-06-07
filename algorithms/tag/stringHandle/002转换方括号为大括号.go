package stringHandle

import "strings"

func convertToBig(str string) string {
	str = strings.ReplaceAll(str, "[", "{")
	str = strings.ReplaceAll(str, "]", "}")
	return str
}