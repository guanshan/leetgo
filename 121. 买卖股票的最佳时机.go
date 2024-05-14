package main

import (
	"math"
)

func maxProfit(prices []int) int {
	buy, sell := math.MinInt64, 0
	for i := 0; i < len(prices); i++ {
		buy = max(buy, -prices[i])
		sell = max(sell, buy+prices[i])
	}
	return sell
}
