package main

// Link: https://leetcode.com/problems/permutations/
/**
Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.



Example 1:

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
Example 2:

Input: nums = [0,1]
Output: [[0,1],[1,0]]
Example 3:

Input: nums = [1]
Output: [[1]]


Constraints:

1 <= nums.length <= 6
-10 <= nums[i] <= 10
All the integers of nums are unique.
*/

/*
	TC: O(n.n!)
	SC: O(n.n!)
*/
func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	size := len(nums)
	backtracking(&ans, nums, make([]int, size), make([]bool, size), 0)
	return ans
}

func backtracking(ans *[][]int, nums []int, cur []int, mem []bool, level int) {
	if level == len(nums) {
		*ans = append(*ans, cur)
		return
	}

	for ind, val := range mem {
		if !val {
			newCur := append([]int(nil), cur...)
			newCur[level] = nums[ind]
			newMem := append([]bool(nil), mem...)
			newMem[ind] = true
			backtracking(ans, nums, newCur, newMem, level+1)
		}
	}
}
