package jsonHandle

import (
	"encoding/json"
	"log"
	"reflect"
)

func CompareJson(json1, json2 []byte) bool {
	path := "CompareJson"
	m1 := make(map[string]interface{})
	m2 := make(map[string]interface{})
	if !CheckJson(json.Unmarshal(json1, &m1), path, "json parse err") {
		return false
	}
	if !CheckJson(json.Unmarshal(json2, &m2), path, "json parse err") {
		return false
	}
	return CompareJsonMap(m1, m2)
}
// 	将json.Unmarshal()后的map作为参数传进来。比较两个json_map是否相同
func CompareJsonMap(d1, d2 interface{}) bool {
	if reflect.ValueOf(d1).Kind() == reflect.Map {
		if len(d1.(map[string]interface{})) != len(d2.(map[string]interface{})) {
			return false
		}
		if !compareMap(d1.(map[string]interface{}), d2.(map[string]interface{})) {
			return false
		}
	} else if reflect.ValueOf(d1).Kind() == reflect.Slice {
		if len(d1.([]interface{})) != len(d2.([]interface{})) {
			return false
		}
		for sli, slv := range d1.([]interface{}) {
			if !CompareJsonMap(slv, d2.([]interface{})[sli]) {
				return false
			}
		}
	} else {
		if d2 != d1 {
			return false
		}
	}
	return true
}


func compareMap(m1, m2 map[string]interface{}) bool {
	for k, v := range m1 {
		if reflect.ValueOf(v).Kind() == reflect.Map {
			if len(m1) != len(m2) {
				return false
			}
			if !compareMap(v.(map[string]interface{}), m2[k].(map[string]interface{})) {
				return false
			}
		}else if reflect.ValueOf(v).Kind() == reflect.Slice { // Slice不能直接比较
			if len(v.([]interface{})) != len(m2[k].([]interface{})) {
				return false
			}
			for sli, slv := range v.([]interface{}) {
				if reflect.ValueOf(slv).Kind() == reflect.Map {
					if !compareMap(slv.(map[string]interface{}), m2[k].([]interface{})[sli].(map[string]interface{})) {
						return false
					}
				}
			}
			continue
		}else if v != m2[k] {
			return false
		}
	}
	return true
}
// 检验json错误
// 报错位置和错误信息
func CheckJson(err error,path, msg string) (isSuc bool){
	if err != nil {
		log.Println(msg, "path", path, "err", err)
		return
	}
	return true
}