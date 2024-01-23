package main

import (
	"sort"
)

/*
*

Link: https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone/description/
2037. Minimum Number of Moves to Seat Everyone
*/

func minMovesToSeat(seats []int, students []int) int {
	n := len(seats)
	if n == 0 {
		return 0
	}

	sort.Slice(seats, func(i, j int) bool { return seats[i] < seats[j] })
	sort.Slice(students, func(i, j int) bool { return students[i] < students[j] })

	ans := 0
	for i := 0; i < n; i++ {
		ans += abs(seats[i] - students[i])
	}

	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
