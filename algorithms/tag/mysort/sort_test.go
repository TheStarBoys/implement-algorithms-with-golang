package mysort

import (
	"testing"
	"errors"
	"fmt"
	"math/rand"
	"time"
	"sort"
)

func TestSort(t *testing.T) {
	var err error
	err = testSort(bubbleSort)
	if err != nil {
		t.Errorf("bubbleSort err: %v", err)
	}

	err = testSort(InsertionSort_Up)
	if err != nil {
		t.Errorf("insertionSort err: %v", err)
	}

	err = testSort(selectionSort)
	if err != nil {
		t.Errorf("selectionSort err: %v", err)
	}

	err = testSort(shellSort)
	if err != nil {
		t.Errorf("shellSort err: %v", err)
	}

	err = testSort(func(ints []int) {
		mergeSort(ints, 0, len(ints)-1)
	})
	if err != nil {
		t.Errorf("mergeSort err: %v", err)
	}

	err = testSort(func(ints []int) {
		quickSort(ints, 0, len(ints)-1)
	})
	if err != nil {
		t.Errorf("quickSort err: %v", err)
	}
}

func testSort(f func([]int)) (err error) {
	defer func() {
		if p := recover(); p != nil {
			str, ok := p.(string)
			if ok {
				err = errors.New(str)
			} else {
				err = errors.New("panic")
			}
		}
	}()
	// 测试数组为空和nil的情况
	f([]int{})
	f(nil)
	rand.Seed(time.Now().UnixNano())
	// 测试一千种随机情况
	for i := 0; i < 1000; i++ {
		length := rand.Intn(1000)
		arr := make([]int, length)
		for j := range arr {
			arr[j] = rand.Intn(1000)
		}
		input := append([]int{}, arr...)
		if f(arr); sort.IntsAreSorted(arr) == false {
			return errors.New(fmt.Sprintf("input: %v, output: %v, expect: upSorted", input, arr))
		}
	}

	return err
}

// 测试是否是升序排序
//func isUpSorted(nums []int) bool {
//	for i := 1; i < len(nums); i++ {
//		if nums[i-1] > nums[i] {
//			return false
//		}
//	}
//
//	return true
//}