package main

import "fmt"

func main() {
	fmt.Println("This is for testing.")

	isPalindrome(121)

}

func isPalindrome(x int) bool {
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