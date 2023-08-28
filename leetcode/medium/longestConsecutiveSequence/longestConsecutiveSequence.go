package longestconsecutivesequence

import (
	"sort"
)

/*
Option1: 排序法。

	先排序，然后不断的去比较下一个元素是否是比当前的元素大1
	由于是顺序的，所以下个元素和当前元素不相邻时，直接从下一个元素继续找
*/
func longestConsecutive1(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)
	result := 0
	length := 1

	for i := 0; i < len(nums)-1; i++ {
		if nums[i]+1 == nums[i+1] {
			length++
		} else if nums[i] == nums[i+1] {
			continue
		} else {
			length = 1
		}
		if length > result {
			result = length
		}
	}

	if length > result {
		result = length
	}

	return result
}

/*
Option2: 使用哈希表
思路：先将元素添加到哈希表中，在Go中只能将元素当作map的key。
然后开始遍历给的数字，每一次遍历都在哈希表里找一下，如果有的话，就继续找，并且记录长度
*/
func longestConsecutive2(nums []int) int {
	var m = make(map[int]bool)
	for _, value := range nums {
		m[value] = true
	}

	result := 0

	for _, value := range nums {
		result = loopUp2(m, value, result)
	}
	return result
}

func loopUp2(m map[int]bool, target int, result int) int {
	longth := 1
	_, existed := m[target+1]
	for existed {
		longth++
		target++
		_, existed = m[target+1]
	}
	if longth > result {
		return longth
	}
	return result
}

/*
Option3: 在哈希表上改进了一下，
思路，在添加到哈希表之前，先排序，然后加到哈希表中，并且记录索引位置
遍历这些数字，并在哈希表中查找，如果发现下一个不联系的话，就记录好索引位置，直接从下一个开始继续找
*/
func longestConsecutive3(nums []int) int {
	// sort.Ints(nums)
	var m = make(map[int]int)
	for index, value := range nums {
		m[value] = index
	}

	finalResult := 0

	for i := 0; i < len(nums); i++ {
		newResult, index := loopUp3(m, nums[i], finalResult)
		finalResult = newResult
		if i < index {
			i = index
		}
	}

	return finalResult
}

func loopUp3(m map[int]int, target int, result int) (newResult int, index int) {
	longth := 1
	cursor := 0
	index, existed := m[target+1]
	cursor = index
	for existed {
		longth++
		target++
		index, existed = m[target+1]
		if cursor < index {
			cursor = index
		}
	}
	if longth > result {
		return longth, cursor
	}
	return result, cursor
}

/*
Option4: 哈希表改进方法
在使用法希表查找之前，判断一下前一个连续的元素是否存在(当前元素减1)
*/
func longestConsecutive4(nums []int) int {
	var m = make(map[int]bool)
	for _, value := range nums {
		m[value] = true
	}

	result := 0

	for _, value := range nums {
		result = loopUp2(m, value, result)
	}
	return result
}

func loopUp4(m map[int]bool, target int, result int) int {
	longth := 1
	if _, existed := m[target-1]; !existed {
		_, existed := m[target+1]
		for existed {
			longth++
			target++
			_, existed = m[target+1]
		}
		if longth > result {
			return longth
		}
	}
	return result
}