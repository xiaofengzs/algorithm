package main

import (
	"fmt"
	// "sort"
)

func main() {
	fmt.Println("This is for testing.")

	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	fmt.Println(maxArea2(height))
}

/*
Option2: 设置左右两个指针，指向头和尾，计算面积，然后移动值较小的那个指针
*/
func maxArea2(height []int) int {
	left, right, result := 0, len(height) - 1, 0
	for left != right {
		result = Max(Min(height[left], height[right]) * (right - left), result)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return result
}

func Min(x, y int) int {
	if x < y {
	  return x
	}
	return y
}

	func Max(x, y int) int{
		if x > y {
		return x
		}
		return y
	}

/*
Option1: 嵌套遍历两边数组
*/
func maxArea1(height []int) int {
	length := len(height)
	result := 0
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			var h int
			if height[i] < height[j] {
				h = height[i]
			} else {
				h = height[j]
			}
			tmpResult := h * (j - i)
			if tmpResult > result {
				result = tmpResult
			}
		}
	}
	return result
}
