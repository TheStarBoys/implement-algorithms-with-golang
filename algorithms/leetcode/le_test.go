package leetcode

import (
	"testing"
	"os"
	"io/ioutil"
	"fmt"
	"path/filepath"
)

func TestLe(t *testing.T) {
	dir, _ := os.Getwd()

	fmt.Println(GetAllFile(dir))
}


// 文件改名
func GetAllFile(pathname string) error {
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			fmt.Printf("[%s]\n", pathname+"\\"+fi.Name())
			GetAllFile(pathname + fi.Name() + "\\")
		} else {
			fmt.Println(fi.Name())
			err := os.Rename(filepath.Join(pathname, fi.Name()), filepath.Join(pathname, "0" + fi.Name()))
			if err != nil {
				fmt.Println("rename err:", err)
				continue
			}

		}
	}
	return err
}