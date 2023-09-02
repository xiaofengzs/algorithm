package movezeros

/*
Option3: 剔除中间有0的元素，再后面追加
*/
func moveZeroes3(nums []int) {
	p := 0
	length := len(nums)
	for p < length {
		if nums[p] == 0 {
			nums = append(nums[:p], nums[p+1:]...)
			nums = append(nums, 0)
			length --
		} else {
			p++
		}
	}
}

/*
Optoin2: 左右两指针交换非0到前面
左指针指向0，右指针指非0，将非0和0交换
这样会有个特点，左指针左边的元素都是非0，右指针左边和左指针右边之间的元素都是0
*/
func moveZeroes2(nums []int) {
	l := 0
	r := 0
	length := len(nums)

	for l < length {
		if nums[l] != 0 {
			nums[r], nums[l] = nums[l], nums[r]
			r++
		}
		l++
	}
}

/* Option1 :  两个指针将非0的元素换到前面，自己想出来的
前后两个指针，第一个指针非0的元素，第二个去找0，然后连个互相交换
*/
func moveZeroes1(nums []int) {
	for index, _ := range nums {
		if index == len(nums)-1 {
			break
		}
		first := &nums[index]
		second := &nums[index+1]
		if *first != 0 {
			continue
		}

		for i := index + 1; i < len(nums); i++ {
			second = &nums[i]
			if *second != 0 {
				break
			}
		}

		if *second != 0 {
			t := *first
			*first = *second
			*second = t
		}
	}
}
