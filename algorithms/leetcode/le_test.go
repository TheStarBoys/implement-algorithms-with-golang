package leetcode

import (
	"testing"
	"os"
	"io/ioutil"
	"fmt"
)

func TestLe(t *testing.T) {
	dir, _ := os.Getwd()

	fmt.Println(GetAllFile(dir))
}

func GetAllFile(pathname string) error {
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			fmt.Printf("[%s]\n", pathname+"\\"+fi.Name())
			GetAllFile(pathname + fi.Name() + "\\")
		} else {
			fmt.Println(fi.Name())

		}
	}
	return err
}