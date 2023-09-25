package main

import (
	"fmt"
)

/*
2741. Special Permutations: https://leetcode.com/problems/special-permutations/
You are given a 0-indexed integer array nums containing n distinct positive integers. A permutation of nums is called special if:

For all indexes 0 <= i < n - 1, either nums[i] % nums[i+1] == 0 or nums[i+1] % nums[i] == 0.
Return the total number of special permutations. As the answer could be large, return it modulo 109 + 7.



Example 1:

Input: nums = [2,3,6]
Output: 2
Explanation: [3,6,2] and [2,6,3] are the two special permutations of nums.
Example 2:

Input: nums = [1,4,3]
Output: 2
Explanation: [3,1,4] and [4,1,3] are the two special permutations of nums.


Constraints:

2 <= nums.length <= 14
1 <= nums[i] <= 109
*/

const mod = 1000000007

func specialPerm(nums []int) int {
	size := len(nums)
	dp := make(map[int][]int, size)
	for _, val := range nums {
		dp[val] = make([]int, 1<<size)
		for ind := range dp[val] {
			dp[val][ind] = -1
		}
	}
	ans := count(0, 1, 1, nums, dp)
	return ans
}

func count(mask, last, val int, nums []int, dp map[int][]int) int {
	if !isGood(val, last) {
		return 0
	}

	if mask == 1<<len(nums)-1 {
		return 1
	}

	if _, ok := dp[val]; ok && dp[val][mask] != -1 {
		return dp[val][mask]
	}

	ans := 0
	for cur, nV := range nums {
		if mask&(1<<cur) == 0 {
			ans = (ans + count(mask|(1<<(cur)), val, nV, nums, dp)%mod) % mod
		}
	}

	if _, ok := dp[val]; ok {
		dp[val][mask] = ans
	}
	return ans
}

func isGood(a, b int) bool {
	if a%b == 0 || b%a == 0 {
		return true
	}
	return false
}

func main() {
	nums := []int{2, 4, 8, 16}
	fmt.Println(specialPerm(nums))
}
