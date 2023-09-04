package longestsubstringwithoutrepeatingcharacters

import "strings"



/*
option3: 使用滑动窗口，左指针和右指针分别代表边界
左右指针从指向第一元素开始，右指针开始向后去找不重复的最长的字串，
当发现重复，则不断的将左指针向右移，这样移除开头的元素，直至窗口中没有重复的元素
当窗口中没有重复的元素，则继续右移左指针去寻找下一个最长不重复字串
*/ 
func lengthOfLongestSubstring3(s string) int {
	length := len(s)
	p := 0
	m := map[byte]int{}
	result := 0
	for i := 0; i < length; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for p < length && m[s[p]] == 0 {
			m[s[p]]++
			p++
		}
		result = Max(result, p - i)
	}
	
	return result
}
/*
option2: 在option1上改了一下，但是改进失败
*/
func lengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 1
	}

	length := len(s)
	result := 0

	for i := 0; i < length; i++ {
		tmpResult := 1
		m := make(map[byte]int)
		m[s[i]] = i
		p := i + 1
		for p < length {
			if index, ok := m[s[p]]; ok {
				i = index
				result = Max(result, tmpResult)
				break
			}
			m[s[p]] = p
			p++
			tmpResult++
		}
	}
	return result
}


/*
option1: 基于一个字母作为开头，去寻找不重复的字串，如果重复了就记录好长度，然后跳到下一个字母重复操作
*/
func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 1
	}

	length := len(s)
	result := 0
	for index, value := range s {
		tmpResult := 1
		m := make(map[rune]int)
		m[value] = index
		p := index + 1
		for p < length {
			if _, ok := m[rune(s[p])]; ok {
				result = Max(result, tmpResult)
				break
			}
			m[rune(s[p])] = p
			p++
			tmpResult++
		}
	}
	return result
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}