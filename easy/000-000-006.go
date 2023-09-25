package main

import "fmt"

/**
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.


Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false


Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/

func isValid(s string) bool {
	m := make(map[int32]int32, 0)
	m[')'] = '('
	m[']'] = '['
	m['}'] = '{'
	stack := make([]int32, 0)
	for _, ch := range s {
		if len(stack) != 0 && m[ch] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()[]{}"))
}
