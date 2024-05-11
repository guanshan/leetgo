package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test 27. 移除元素
func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	length := removeElement(nums, val)
	assert.Equal(t, nums[:length], []int{2, 2})

	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	val = 2
	length = removeElement(nums, val)
	assert.Equal(t, nums[:length], []int{0, 1, 3, 0, 4})
}

// Test 88. 合并两个有序数组
func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	merge(nums1, m, nums2, n)
	assert.Equal(t, nums1, []int{1, 2, 2, 3, 5, 6})

	nums1 = []int{1}
	m = 1
	nums2 = []int{}
	n = 0

	merge(nums1, m, nums2, n)
	assert.Equal(t, nums1, []int{1})

	nums1 = []int{0}
	m = 0
	nums2 = []int{1}
	n = 1

	merge(nums1, m, nums2, n)
	assert.Equal(t, nums1, []int{1})

	nums1 = []int{4, 5, 6, 0, 0, 0}
	m = 3
	nums2 = []int{1, 2, 3}
	n = 3

	merge(nums1, m, nums2, n)
	assert.Equal(t, nums1, []int{1, 2, 3, 4, 5, 6})
}
