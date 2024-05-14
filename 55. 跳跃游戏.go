package main

func canJump(nums []int) bool {
	far := nums[0]
	for i := 0; i < len(nums)-1; i++ {
		far = max(far, i+nums[i])
		if far <= i {
			return false
		}
	}
	return far >= len(nums)-1
}
