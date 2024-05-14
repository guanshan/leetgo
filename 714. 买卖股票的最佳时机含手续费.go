package main

import (
	"math"
)

func maxProfitFee(prices []int, fee int) int {
	buy, sell := math.MinInt64, 0
	for i := 0; i < len(prices); i++ {
		currProfit := sell
		buy = max(buy, currProfit-prices[i])
		sell = max(sell, buy+prices[i]-fee)
	}
	return sell
}
