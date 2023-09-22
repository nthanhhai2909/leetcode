package main

import "fmt"

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

// back-tracking
func specialPerm(nums []int) int {

	end := 0
	for ind := range nums {
		end = end | (1 << ind)
	}
	return count(0, end, 1, 1, nums)
}

func count(mask, end, last, val int, nums []int) int {
	mod := 1000000007
	if !isGood(val, last) {
		return 0
	}

	if mask == end {
		return 1
	}

	ans := 0
	for cur, nV := range nums {
		if mask&(1<<cur) == 0 {
			ans = (ans + count(mask|(1<<(cur)), end, val, nV, nums)%mod) % mod
		}
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
	nums := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192}
	fmt.Println(specialPerm(nums))
}
