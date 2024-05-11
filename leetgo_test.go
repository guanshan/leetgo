package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
