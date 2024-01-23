package main

/*
Leetcode: https://leetcode.com/problems/steps-to-make-array-non-decreasing/

You are given a 0-indexed integer array nums. In one step, remove all elements nums[i] where nums[i - 1] > nums[i] for all 0 < i < nums.length.

Return the number of steps performed until nums becomes a non-decreasing array.
Example 1:

Input: nums = [5,3,4,4,7,3,6,11,8,5,11]
Output: 3
Explanation: The following are the steps performed:
- Step 1: [5,3,4,4,7,3,6,11,8,5,11] becomes [5,4,4,7,6,11,11]
- Step 2: [5,4,4,7,6,11,11] becomes [5,4,7,11,11]
- Step 3: [5,4,7,11,11] becomes [5,7,11,11]
[5,7,11,11] is a non-decreasing array. Therefore, we return 3.
Example 2:

Input: nums = [4,5,7,7,13]
Output: 0
Explanation: nums is already a non-decreasing array. Therefore, we return 0.


Constraints:

1 <= nums.length <= 105
1 <= nums[i] <= 109

*/

func totalSteps(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	m := make(map[int]struct{})
	ans := 0
	start := 0
	for {
		p1 := start
		p2 := start + 1
		temp := make([]int, 0)
		isSet := false
		for p2 < len(nums) {
			if _, ok := m[p1]; ok {
				p1++
				p2++
				continue
			}
			if _, ok := m[p2]; ok {
				p2++
				continue
			}
			if nums[p1] > nums[p2] {
				temp = append(temp, p2)
				if !isSet {
					start = p1
				}
				isSet = true
			}
			p1 = p2
			p2++
		}
		if len(temp) == 0 {
			break
		}
		for _, val := range temp {
			m[val] = struct{}{}
		}
		ans++
	}
	return ans
}

// ----------------------------------------------------------------------------------
type Pair struct {
	Key   int
	Value int
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func totalStepsV2(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	stack := make([]Pair, 0)
	stack = append(stack, Pair{nums[size-1], 0})
	ans := 0
	for i := size - 2; i >= 0; i-- {
		c := 0
		for len(stack) > 0 && nums[i] > stack[len(stack)-1].Key {
			value := stack[len(stack)-1].Value
			stack = stack[:len(stack)-1]
			c = max(c+1, value)
		}
		stack = append(stack, Pair{nums[i], c})
		ans = max(ans, c)
	}
	return ans
}
