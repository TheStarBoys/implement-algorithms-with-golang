package main

import "fmt"

func main() {
	m := map[string]bool{}
	m["hello"] = true
	fmt.Println(m["hello"])
}