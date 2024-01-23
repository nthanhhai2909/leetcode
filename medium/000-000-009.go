package main

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

// ------------------------------------------------------------

func nextBeautifulNumberv2(n int) int {
	for {
		n = n + 1
		if isValid(n) {
			return n
		}
	}
}
