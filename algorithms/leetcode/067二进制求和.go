package leetcode

func addBinary067_0(a string, b string) string {
	result := ""
	i, j := len(a) - 1, len(b) - 1
	flag := 0
	for i >= 0 || j >= 0 {
		t1, t2 := 0, 0
		if i >= 0 {
			t1 = int(a[i] - '0')
		}
		if j >= 0 {
			t2 = int(b[j] - '0')
		}
		switch t1 + t2 + flag {
		case 3 : flag = 1; result = "1" + result
		case 2 : flag = 1; result = "0" + result
		case 1 : flag = 0; result = "1" + result
		case 0 : flag = 0; result = "0" + result
		}
		i--
		j--
	}
	if flag == 1 {
		return "1" + result
	}
	return result
}