package main

/*
Leetcode: https://leetcode.com/problems/count-square-submatrices-with-all-ones/

Given a m * n matrix of ones and zeros, return how many square submatrices have all ones.

Example 1:

Input: matrix =
[
  [0,1,1,1],
  [1,1,1,1],
  [0,1,1,1]
]
Output: 15
Explanation:
There are 10 squares of side 1.
There are 4 squares of side 2.
There is  1 square of side 3.
Total number of squares = 10 + 4 + 1 = 15.
Example 2:

Input: matrix =
[
  [1,0,1],
  [1,1,0],
  [1,1,0]
]
Output: 7
Explanation:
There are 6 squares of side 1.
There is 1 square of side 2.
Total number of squares = 6 + 1 = 7.


Constraints:

1 <= arr.length <= 300
1 <= arr[0].length <= 300
0 <= arr[i][j] <= 1
*/

func countSquares(matrix [][]int) int {
	ans := 0
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	cur_m := make([]int, n)
	bef_m := make([]int, n)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			cur_m[j] = matrix[i][j]
			if i != m-1 && j != n-1 && matrix[i][j] > 0 {
				cur_m[j] = min(cur_m[j+1], bef_m[j], bef_m[j+1]) + 1
			}
			ans += cur_m[j]
		}
		for k := 0; k < n; k++ {
			bef_m[k] = cur_m[k]
		}
	}
	return ans
}
func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}

	if b <= a && b <= c {
		return b
	}
	return c
}
