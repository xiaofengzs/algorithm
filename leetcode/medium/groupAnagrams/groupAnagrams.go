package groupAnagrams

import (
	"fmt"
	"sort"
	"strings"
)

// option3: 统计每个单词中字母出现的次数，字母异位词的字母出现的次数一定相等
// 这里有个技巧是需要巧妙的利用数组来解题，字母'a'减去‘a’得到0，也就是数组的第0位表示a，以此类推
// 将字母出现次数相同的词放到一个切片中，作为值，字母统计作为key，将它们放入map中
// 最后使用map转换结果
func groupAnagrams3(strs []string) [][]string {
	tmpResult := make(map[[26]int][]string)
	for _, word := range strs {
		key := [26]int{}
		for _, letter := range word {
			key[letter-'a']++
		}
		tmpResult[key] = append(tmpResult[key], word)
	}
	var result = make([][]string, 0, len(tmpResult))
	for _, value := range tmpResult {
		result = append(result, value)
	}
	return result
}

// option2: 将每个单词中的字母对应的质数的积算一下，然后如果一样的话就把原单词加入到map中，
// 排序后的单词作为key，slice作为值，把原单词加入到切片中
// 最后再把map转换为切片
func groupAnagrams2(strs []string) [][]string {
	letterPrimeNumberMappings := map[string]int{
		"a": 3, "b": 5, "c": 7, "d": 11, "e": 13, "f": 17, "g": 19,
		"h": 23, "i": 29, "j": 31, "k": 37, "l": 41, "m": 43, "n": 47,
		"o": 53, "p": 59, "q": 61, "r": 67, "s": 71, "t": 73, "u": 79,
		"v": 83, "w": 89, "x": 97, "y": 101, "z": 103,
	}
	var tmpResult map[int][]string = make(map[int][]string)
	for _, word := range strs {
		total := 1
		for _, letter := range word {
			if primeNumber, ok := letterPrimeNumberMappings[string(letter)]; ok {
				total = total * primeNumber
			}
		}
		tmpResult[total] = append(tmpResult[total], word)
	}

	var result = make([][]string, 0, len(tmpResult))
	for _, value := range tmpResult {
		result = append(result, value)
	}

	return result
}

// option1: 将每个单词排序，然后如果一样的话就把原单词加入到map中，
// 排序后的单词作为key，slice作为值，把原单词加入到切片中
// 最后再把map转换为切片
func groupAnagrams1(strs []string) [][]string {
	var tmpResult map[string][]string = make(map[string][]string)

	for _, str := range strs {
		strSliceforSort := strings.Split(str, "")
		sort.Strings(strSliceforSort)
		sortedStr := strings.Join(strSliceforSort, "")

		tmpResult[sortedStr] = append(tmpResult[sortedStr], str)
	}

	var result = make([][]string, 0, len(tmpResult))
	for _, value := range tmpResult {
		result = append(result, value)
	}

	return result
}
