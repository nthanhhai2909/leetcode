package main

/*
*

Leetcode: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

Example 1:
Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

Example 2:
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.
*/

// Brute force
func maxProfitV1(prices []int) int {
	size := len(prices)
	if size <= 1 {
		return 0
	}
	ans := 0
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			val := prices[j] - prices[i]
			if val > ans {
				ans = val
			}
		}
	}

	if ans > 0 {
		return ans
	}
	return 0
}

func maxProfitV2(prices []int) int {
	size := len(prices)
	if size <= 1 {
		return 0
	}
	ans := 0
	b := 0
	s := 1

	for s < size {
		profit := prices[s] - prices[b]
		if profit > ans {
			ans = profit
		}
		if profit < 0 {
			b = s
		}
		s++
	}

	return ans
}
