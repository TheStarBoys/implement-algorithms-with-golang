package leetcode

// 回溯的一种写法
func subsets078_0(nums []int) [][]int {
	var result [][]int
	result = backtrack078_0(nums, 0, result, []int{})
	return result
}

func backtrack078_0(nums []int, indx int, result [][]int, tmp []int) [][]int {
	// append底层扩容的时候，slice引用的底层数组会发生变化
	// slice是值传递，能修改引用的底层数组的值是因为slice底层里有一个指针，指向的数组元素
	// 当函数内的slice发生改变，由于slice并不是引用，slice里的指针才是引用, 所以函数内部的append扩容
	// 分配了一个更大容量的slice，将数组元素进行copy到新的slice的底层数组指针里，再添加新的元素值
	// 至此，产生了两个不同的slice，指向了不同的底层数组
	// 导致函数外传进的result感知不到这种变化
	result = append(result, tmp)

	for i := indx; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
		t := make([]int, len(tmp))
		copy(t, tmp)
		result = backtrack078_0(nums, i+1, result, t)
		// 回溯
		tmp = tmp[:len(tmp)-1]
	}
	return result
}

// 回溯的第二种写法

func subsets078_1(nums []int) [][]int {
	var result [][]int
	backtrack078_1(nums, 0, &result, []int{})
	return result
}

func backtrack078_1(nums []int, indx int, result *[][]int, tmp []int) {
	*result = append(*result, tmp)

	for i := indx; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
		t := make([]int, len(tmp))
		copy(t, tmp)
		backtrack078_1(nums, i+1, result, t)
		// 回溯
		tmp = tmp[:len(tmp)-1]
	}
}