package stringMatching

import (
	"testing"
	"fmt"
)

func TestStringMatching(t *testing.T) {
	str, sub := "ababbcafababacd", "ababacd"

	fmt.Println("-----Front-----")
	fmt.Printf("BF:\t\t\t%d\n", BFSearch(str, sub))
	fmt.Printf("bfSearch:\t%d\n", bfSearch(str, sub))
	fmt.Printf("Sunday:\t\t%d\n", SundaySearch(str, sub))
	fmt.Printf("RK:\t\t\t%d\n", RKSearch(str, sub))
	fmt.Printf("BM:\t\t\t%d\n", BMSearch(str, sub))
	fmt.Printf("KMP:\t\t%d\n", KMPSearch(str, sub))

	fmt.Println("-----Mid-----")
	str, sub = "ababbcafababacd", "cafa"
	fmt.Printf("BF:\t\t\t%d\n", BFSearch(str, sub))
	fmt.Printf("bfSearch:\t%d\n", bfSearch(str, sub))
	fmt.Printf("Sunday:\t\t%d\n", SundaySearch(str, sub))
	fmt.Printf("RK:\t\t\t%d\n", RKSearch(str, sub))
	fmt.Printf("BM:\t\t\t%d\n", BMSearch(str, sub))
	fmt.Printf("KMP:\t\t%d\n", KMPSearch(str, sub))

	fmt.Println("-----End-----")
	str, sub = "ababbcafababacd", "bacd"
	fmt.Printf("BF:\t\t\t%d\n", BFSearch(str, sub))
	fmt.Printf("bfSearch:\t%d\n", bfSearch(str, sub))
	fmt.Printf("Sunday:\t\t%d\n", SundaySearch(str, sub))
	fmt.Printf("RK:\t\t\t%d\n", RKSearch(str, sub))
	fmt.Printf("BM:\t\t\t%d\n", BMSearch(str, sub))
	fmt.Printf("KMP:\t\t%d\n", KMPSearch(str, sub))

	fmt.Println("-----Less-----")
	str, sub = "he", "wor"
	fmt.Printf("BF:\t\t\t%d\n", BFSearch(str, sub))
	fmt.Printf("bfSearch:\t%d\n", bfSearch(str, sub))
	fmt.Printf("Sunday:\t\t%d\n", SundaySearch(str, sub))
	fmt.Printf("RK:\t\t\t%d\n", RKSearch(str, sub))
	fmt.Printf("BM:\t\t\t%d\n", BMSearch(str, sub))
	fmt.Printf("KMP:\t\t%d\n", KMPSearch(str, sub))

	fmt.Println("-----No Exist-----")
	str, sub = "hello, world", "dfaf"
	fmt.Printf("BF:\t\t\t%d\n", BFSearch(str, sub))
	fmt.Printf("bfSearch:\t%d\n", bfSearch(str, sub))
	fmt.Printf("Sunday:\t\t%d\n", SundaySearch(str, sub))
	fmt.Printf("RK:\t\t\t%d\n", RKSearch(str, sub))
	fmt.Printf("BM:\t\t\t%d\n", BMSearch(str, sub))
	fmt.Printf("KMP:\t\t%d\n", KMPSearch(str, sub))
}
