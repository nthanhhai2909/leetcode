package main

import (
	"sort"
)

/*
Link:https://leetcode.com/problems/next-greater-numerically-balanced-number/description/
2048. Next Greater Numerically Balanced Number
*/
func nextBeautifulNumber(n int) int {
	for {
		n = n + 1
		if isValid(n) {
			return n
		}
	}
}

func isValid(n int) bool {
	m := make(map[int]int)
	for n > 0 {
		k := n % 10
		if v, ok := m[k]; ok {
			m[k] = v + 1
		} else {
			m[k] = 1
		}
		n = n / 10
	}

	for k, v := range m {
		if k != v {
			return false
		}
	}
	return true
}

// ---------------------------------------------------------------------------

func nextBeautifulNumberV2(n int) int {
	size := 0
	temp := n
	for temp > 0 {
		size++
		temp = temp / 10
	}
	var nums [][]int
	data := []int{1, 2, 3, 4, 5, 6}
	findSubNumbers(data, size, size, []int{}, &nums)
	findSubNumbers(data, size+1, size+1, []int{}, &nums)
	ans := possibleNums(nums)
	sort.Slice(ans, func(i, j int) bool { return ans[i] < ans[j] })
	for _, v := range ans {
		if v > n {
			return v
		}
	}
	return -1
}

func findSubNumbers(nums []int, sum, remainingSum int, subNums []int, result *[][]int) {
	if remainingSum == 0 {
		*result = append(*result, append([]int{}, subNums...))
		return
	}
	if remainingSum < 0 {
		return
	}
	for i, num := range nums {
		findSubNumbers(nums[i+1:], sum, remainingSum-num, append(subNums, num), result)
	}
}

func possibleNums(nums [][]int) []int {
	ans := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		m := make(map[int]int)
		sum := 0
		for j := 0; j < len(nums[i]); j++ {
			m[nums[i][j]] = nums[i][j]
			sum += nums[i][j]
		}
		bt(0, 0, sum, m, &ans)
	}

	return ans
}

func bt(num int, cur_count int, sum int, cur_state map[int]int, result *[]int) {
	if cur_count == sum {
		*result = append(*result, num)
		return
	}

	for k, v := range cur_state {
		if v > 0 {
			clonedMap := make(map[int]int)
			for key, value := range cur_state {
				clonedMap[key] = value
			}
			clonedMap[k]--
			bt(num*10+k, cur_count+1, sum, clonedMap, result)
		}

	}
}
