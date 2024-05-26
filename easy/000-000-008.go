package main

import "fmt"

// Link: https://leetcode.com/problems/existence-of-a-substring-in-a-string-and-its-reverse/

/*
Given a string s, find any
substring
 of length 2 which is also present in the reverse of s.

Return true if such a substring exists, and false otherwise.



Example 1:

Input: s = "leetcode"

Output: true

Explanation: Substring "ee" is of length 2 which is also present in reverse(s) == "edocteel".

Example 2:

Input: s = "abcba"

Output: true

Explanation: All of the substrings of length 2 "ab", "bc", "cb", "ba" are also present in reverse(s) == "abcba".

Example 3:

Input: s = "abcd"

Output: false

Explanation: There is no substring of length 2 in s, which is also present in the reverse of s.



Constraints:

1 <= s.length <= 100
s consists only of lowercase English letters.
*/

type Set map[byte]struct{}

func NewSet() Set {
	return make(Set)
}

func (s Set) Add(value byte) {
	s[value] = struct{}{}
}

func (s Set) Contains(value byte) bool {
	_, exists := s[value]
	return exists
}

func isSubstringPresent(s string) bool {
	size := len(s)
	if size < 2 {
		return false
	}
	m := make(map[byte]Set, 0)
	for ind := 0; ind < size-1; ind++ {
		if set, ok := m[s[ind]]; ok {
			set.Add(s[ind+1])
		} else {
			values := NewSet()
			values.Add(s[ind+1])
			m[s[ind]] = values
		}
	}
	for ind := size - 1; ind > 0; ind-- {
		if val, ok := m[s[ind]]; ok && val.Contains(s[ind-1]) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(isSubstringPresent("iaskslikc"))
}
