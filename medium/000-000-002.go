package main

import (
	"fmt"
	"sort"
)

/*
Leetcode: https://leetcode.com/problems/minimum-number-of-people-to-teach/
On a social network consisting of m users and some friendships between users, two users can communicate with each other if they know a common language.

You are given an integer n, an array languages, and an array friendships where:

There are n languages numbered 1 through n,
languages[i] is the set of languages the i(th) user knows, and
friendships[i] = [ui, vi] denotes a friendship between the users ui and vi.
You can choose one language and teach it to some users so that all friends can communicate with each other. Return the minimum number of users you connect to teach.

Note that friendships are not transitive, meaning if x is a friend of y and y is a friend of z, this doesn't guarantee that x is a friend of z.


Example 1:

Input: n = 2, languages = [[1],[2],[1,2]], friendships = [[1,2],[1,3],[2,3]]
Output: 1
Explanation: You can either teach user 1 the second language or user 2 the first language.
Example 2:

Input: n = 3, languages = [[2],[1,3],[1,2],[3]], friendships = [[1,4],[1,2],[3,4],[2,3]]
Output: 2
Explanation: Teach the third language to users 1 and 3, yielding two users to teach.
*/

// ---------------------------------------------------- Start Version 1 ------------------------------------------------

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	users := make(map[int]struct{})
	for _, lags := range languages {
		sort.Slice(lags, func(i, j int) bool { return lags[i] < lags[j] })
	}
	for i := 0; i < len(friendships); i++ {
		f1 := friendships[i][0]
		f2 := friendships[i][1]
		if !canTalk(languages[f1-1], languages[f2-1]) {
			users[f1] = struct{}{}
			users[f2] = struct{}{}
		}
	}
	cl := make([]int, n)
	max := 0
	for user := range users {
		lag := languages[user-1]
		for _, l := range lag {
			cl[l-1]++
			if max < cl[l-1] {
				max = cl[l-1]
			}
		}
	}
	return len(users) - max
}

func canTalk(l1, l2 []int) bool {
	p1 := 0
	p2 := 0
	for p1 < len(l1) && p2 < len(l2) {
		if l1[p1] == l2[p2] {
			return true
		} else if l1[p1] > l2[p2] {
			p2++
		} else {
			p1++
		}
	}
	return false
}

func main() {
	fmt.Println("Min: ", minimumTeachings(3,
		[][]int{{2}, {1, 3}, {1, 2}, {3}},
		[][]int{{1, 4}, {1, 2}, {3, 4}, {2, 3}})) // expected 3
	//fmt.Println("Min: ", minimumTeachings(2, [][]int{{1}, {2}, {1, 2}}, [][]int{{1, 2}, {1, 3}, {2, 3}}))
}

// ---------------------------------------------------------------------------------------------------------------------
