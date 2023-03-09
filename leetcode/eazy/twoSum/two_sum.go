package eazy

import "strconv"

// 方法1: 暴力破解，遍历nums切片，并且在每一次取值时，
// 开启一个新的遍历，去当前值后面的全部数据，并进行计算
// 时间复杂度: O(N^2)
// 空间复杂度: O(1)
func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}


// 方法2: 利用字典存放nums中的数字，然后在遍历下个值时计算出差值，
// 在字典里面寻找有没有这个差值，这样的话只需要便利nums切片一遍，
// 在每一次取值处理之中去字典里面查找一次
// 时间复杂度: O(N)
// 空间复杂度: O(n)
func twoSum2(nums []int, target int) []int {
	backupNums := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		difference := target - nums[i]
		if indexOfNums, ok := backupNums[difference]; ok {
			return []int{i, indexOfNums}
		}
		backupNums[nums[i]] = i
	}
	return nil
}

func twoSum3(nums []int, target int) []int {
	backupNums := make(map[int]int)
	for i, v := range nums {
		difference := target - v
		if indexOfNums, ok := backupNums[difference]; ok {
			return []int{i, indexOfNums}
		}
		backupNums[v] = i
	}
	return nil
}
