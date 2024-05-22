package main

import (
	"sort"
	"strconv"
	"strings"
)

// Link: https://leetcode.com/problems/minimum-time-difference/
// Space O(n)
// Time O(nlog(n))
func findMinDifference(timePoints []string) int {
	minutes := convertTimePointToMinutes(timePoints)
	sort.Slice(minutes, func(i, j int) bool { return minutes[i] < minutes[j] })
	min := minutes[len(minutes)-1] - minutes[0]
	if 24*60-min < min {
		min = 24*60 - min
	}
	for i := 0; i < len(minutes)-1; i++ {
		if minutes[i+1]-minutes[i] < min {
			min = minutes[i+1] - minutes[i]
		}
	}
	return min
}

func convertTimePointToMinutes(timePoints []string) []int {
	ans := make([]int, 0)
	for _, time := range timePoints {
		arr := strings.Split(time, ":")
		hour, _ := strconv.Atoi(arr[0])
		minute, _ := strconv.Atoi(arr[1])
		ans = append(ans, minute+hour*60)
	}
	return ans
}
