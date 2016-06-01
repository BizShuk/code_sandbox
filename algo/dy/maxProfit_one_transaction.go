package dy

import (
	"math"
)

// only one transaction
func MaxProfit1(prices []int) int {
	var max, cur float64
	for i := 1; i < len(prices); i++ {
		cur += float64(prices[i] - prices[i-1])
		cur = math.Max(0, cur)
		max = math.Max(max, cur)

	}
	return int(max)
}
