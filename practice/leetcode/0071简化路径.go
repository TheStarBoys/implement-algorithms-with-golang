package leetcode

import "strings"

func simplifyPath0071_0(path string) string {
	abs := []string{}
	prevIndx := 0
	for {
		for ; prevIndx < len(path) && path[prevIndx] == '/'; prevIndx++ {}
		if prevIndx == len(path) {
			break
		}
		i := prevIndx
		for ; i < len(path) && path[i] != '/'; i++ {}
		p := path[prevIndx:i]
		prevIndx = i
		switch p {
		case "..":
			if len(abs) > 0 {
				abs = abs[:len(abs)-1]
			}
			continue
		case "." :
			continue
		}
		abs = append(abs, p)
	}

	return "/" + strings.Join(abs, "/")
}
