package slice

import (
	"log"
	"reflect"
)

// reverse slice, support []int and []string
func Reverse(slice interface{}) {
	value := reflect.ValueOf(slice)
	if value.Kind() != reflect.Slice {
		log.Panic("ERROR : Input is not slice")
	}
	if value.Len() == 0 {
		return
	}
	switch value.Index(0).Kind() {
	case reflect.Int:
		sli := slice.([]int)
		for i, j := 0, len(sli) - 1; i < j; i, j = i + 1, j - 1 {
			sli[i], sli[j] = sli[j], sli[i]
		}
	case reflect.String:
		sli := slice.([]string)
		for i, j := 0, len(sli) - 1; i < j; i, j = i + 1, j - 1 {
			sli[i], sli[j] = sli[j], sli[i]
		}
	default:
		log.Panic("ERROR : Input not supported")
	}
}
