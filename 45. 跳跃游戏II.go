package main

func jump(nums []int) int {
	far, end, step := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		far = max(nums[i]+i, far)
		if end == i {
			step++
			end = far
		}
	}
	return step
}
