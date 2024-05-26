package main

import (
	"sort"
)

// Link: https://leetcode.com/problems/permutations-ii/description/
/*
Given a collection of numbers, nums, that might contain duplicates, return all possible unique permutations in any order.



Example 1:

Input: nums = [1,1,2]
Output:
[[1,1,2],
 [1,2,1],
 [2,1,1]]
Example 2:

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]


Constraints:

1 <= nums.length <= 8
-10 <= nums[i] <= 10
*/

/*
	TC: O(n.n!)
	SC: O(n.n!)
*/

func permuteUnique(nums []int) [][]int {
	ans := make([][]int, 0)
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	size := len(nums)
	bkPermuteUnique(&ans, nums, make([]int, size), make([]bool, size), 0)
	return ans
}

func bkPermuteUnique(ans *[][]int, nums []int, cur []int, mem []bool, level int) {
	if level == len(nums) {
		*ans = append(*ans, cur)
		return
	}
	fMatch := false
	var prev int
	for ind, val := range mem {
		if !val {
			if !fMatch {
				fMatch = true
				prev = nums[ind]
			} else {
				if nums[ind] == prev {
					continue
				} else {
					prev = nums[ind]
				}
			}

			newCur := append([]int(nil), cur...)
			newCur[level] = nums[ind]
			newMem := append([]bool(nil), mem...)
			newMem[ind] = true
			bkPermuteUnique(ans, nums, newCur, newMem, level+1)
		}
	}
}
