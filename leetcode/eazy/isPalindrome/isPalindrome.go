package ispalindrome

import "strconv"


// 方法1: 同时从这个数字的开头以及结尾进行比较
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)
	length := len(s)
	for i := 0; i < length; i++ {
		j := length - i - 1
		if (i >= j ) {
			return true
		}
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

// 方法2： 同时从这个数字的开头以及结尾进行比较，使用递归进行比较
func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	return exec(s, 0, len(s)-1)
}

func exec(s string, i int, j int) bool {
	if i >= j {
		return true
	}

	if s[i] != s[j] {
		return false
	}
	return exec(s, i + 1, j - 1)
}

// 方法3：通过将数字对10取余的结果乘10反转数字，判断反转之后的数字是否与目标数字相等
func isPalindrome3(x int) bool {
	if x < 0 {
		return false
	}
	reverse := 0
	res := x
	ctrl := x
	for {
		res = res % 10
		reverse = reverse * 10 + res
		ctrl = ctrl / 10
		res = ctrl
		if ctrl < 1 {
			break
		}
	}
	return reverse == x
}