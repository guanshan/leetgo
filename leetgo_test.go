package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test 714. 买卖股票的最佳时机含手续费
func TestMaxProfitFee(t *testing.T) {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2
	profit := maxProfitFee(prices, fee)
	assert.Equal(t, profit, 8)

	prices = []int{1, 3, 7, 5, 10, 3}
	fee = 3
	profit = maxProfitFee(prices, fee)
	assert.Equal(t, profit, 6)

	prices = []int{1, 3, 2, 8, 4, 9}
	fee = 2
	profit = maxProfitFee(prices, fee)
	assert.Equal(t, profit, 8)

	prices = []int{1, 3, 7, 5, 10, 3}
	fee = 3
	profit = maxProfitFee(prices, fee)
	assert.Equal(t, profit, 6)
}

// Test 55. 跳跃游戏
func TestCanJump(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	can := canJump(nums)
	assert.True(t, can)

	nums = []int{3, 2, 1, 0, 4}
	can = canJump(nums)
	assert.False(t, can)
}

// Test 122. 买卖股票的最佳时机 II
func TestMaxProfit2(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	profit := maxProfit2(prices)
	assert.Equal(t, profit, 7)

	prices = []int{1, 2, 3, 4, 5}
	profit = maxProfit2(prices)
	assert.Equal(t, profit, 4)

	prices = []int{7, 6, 4, 3, 1}
	profit = maxProfit2(prices)
	assert.Equal(t, profit, 0)
}

// Test 121. 买卖股票的最佳时机
func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	profit := maxProfit(prices)
	assert.Equal(t, profit, 5)

	prices = []int{7, 6, 4, 3, 1}
	profit = maxProfit(prices)
	assert.Equal(t, profit, 0)

	prices = []int{1, 2}
	profit = maxProfit(prices)
	assert.Equal(t, profit, 1)
}

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
