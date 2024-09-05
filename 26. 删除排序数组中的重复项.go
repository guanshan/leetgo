package main

func removeDuplicates(nums []int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow += 1
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}
