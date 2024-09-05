package main

func removeDuplicates2(nums []int) int {
	slow, fast, step := 0, 0, 1
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			if step <= 2 {
				slow++
			}

			step = 1
			nums[slow] = nums[fast]
		} else if step <= 2 {
			slow++
			step++
		} else {
			step++
		}
		fast++
	}
	return slow + 1
}
