package main

/*
Leetcode: https://leetcode.com/problems/buddy-strings/
Given two strings s and goal, return true if you can swap two letters in s so the result is equal to goal, otherwise, return false.

Swapping letters is defined as taking two indices i and j (0-indexed) such that i != j and swapping the characters at s[i] and s[j].

For example, swapping at indices 0 and 2 in "abcd" results in "cbad".


Example 1:

Input: s = "ab", goal = "ba"
Output: true
Explanation: You can swap s[0] = 'a' and s[1] = 'b' to get "ba", which is equal to goal.
Example 2:

Input: s = "ab", goal = "ab"
Output: false
Explanation: The only letters you can swap are s[0] = 'a' and s[1] = 'b', which results in "ba" != goal.
Example 3:

Input: s = "aa", goal = "aa"
Output: true
Explanation: You can swap s[0] = 'a' and s[1] = 'a' to get "aa", which is equal to goal.


Constraints:

1 <= s.length, goal.length <= 2 * 104
s and goal consist of lowercase letters.
*/

func buddyStrings(s string, goal string) bool {
	sSize := len(s)
	goalSize := len(goal)
	if sSize != goalSize {
		return false
	}

	if s == goal {
		m := make(map[uint8]int)
		for i := 0; i < sSize; i++ {
			if count, ok := m[s[i]]; ok {
				count++
				if count >= 2 {
					return true
				}
			} else {
				m[s[i]] = 1
			}
		}
		return false
	}

	arrs := make([]uint8, 0)
	dif := 0
	for i := 0; i < sSize; i++ {
		if s[i] != goal[i] {
			arrs = append(arrs, s[i], goal[i])
			dif++
		}

		if dif > 2 {
			return false
		}
	}
	return dif == 2 && arrs[0] == arrs[3] && arrs[1] == arrs[2]
}

func buddyStringsV2(s string, goal string) bool {
	sSize := len(s)
	goalSize := len(goal)
	if sSize != goalSize {
		return false
	}

	arrs := make([]int, 0)
	mems := make([]int, 26)
	for i := 0; i < sSize; i++ {
		mems[s[i]-97]++
		if s[i] != goal[i] {
			arrs = append(arrs, i)
		}

		if len(arrs) > 2 {
			return false
		}
	}
	if len(arrs) == 2 && s[arrs[0]] == goal[arrs[1]] && s[arrs[1]] == goal[arrs[0]] {
		return true
	}

	if len(arrs) == 0 {
		for i := 0; i < 26; i++ {
			if mems[i] >= 2 {
				return true
			}
		}
	}
	return false
}
