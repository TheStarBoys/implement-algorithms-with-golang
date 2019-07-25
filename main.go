package main

import (
	"fmt"
	ds "github.com/TheStarBoys/implement-data-structures-and-algorithms-with-golang/dataStructure"
)

func main() {
	ll := ds.LinkList{}
	ll.ListInsert(0, 1)
	ll.ListInsert(0, 2)
	fmt.Println(ll.GetElem(0))
	ll.ListDelete(0)
	fmt.Println(ll.GetElem(1))

}
