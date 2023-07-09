package main

import (
	"fmt"
)

/*

Leetcode: https://leetcode.com/problems/adding-two-negabinary-numbers/

Negabinary docs: https://en.wikipedia.org/wiki/Negative_base
Given two numbers arr1 and arr2 in base -2, return the result of adding them together.

Each number is given in array format:  as an array of 0s and 1s, from most significant bit to least significant bit.
For example, arr = [1,1,0,1] represents the number (-2)^3 + (-2)^2 + (-2)^0 = -3.
A number arr in array, format is also guaranteed to have no leading zeros: either arr == [0] or arr[0] == 1.

Return the result of adding arr1 and arr2 in the same format: as an array of 0s and 1s with no leading zeros.



Example 1:

Input: arr1 = [1,1,1,1,1], arr2 = [1,0,1]
Output: [1,0,0,0,0]
Explanation: arr1 represents 11, arr2 represents 5, the output represents 16.
Example 2:

Input: arr1 = [0], arr2 = [0]
Output: [0]
Example 3:

Input: arr1 = [0], arr2 = [1]
Output: [1]


Constraints:

1 <= arr1.length, arr2.length <= 1000
arr1[i] and arr2[i] are 0 or 1
arr1 and arr2 have no leading zeros
-----------------------------------------------------------------------------------------------------------------------'

Addition
Adding negabinary numbers proceeds bitwise, starting from the least significant bits;
the bits from each addend are summed with the (balanced ternary) carry from the previous bit (0 at the LSB).
This sum is then decomposed into an output bit and carry for the next iteration as show in the table:

Sum	Output	Comment
		        Bit	Carry
−2	010 (−2)	0	1	01 (−2)	−2 occurs only during subtraction.
−1	011 (−2)	1	1	01 (−2)
0	000 (−2)	0	0	00 (−2)
1	001 (−2)	1	0	00 (−2)
2	110 (−2)	0	−1	11 (−2)
3	111 (−2)	1	−1	11 (−2)	3 occurs only during addition.
The second row of this table, for instance, expresses the fact that −1 = 1 + 1 × −2; the fifth row says 2 = 0 + −1 × −2; etc.

As an example, to add 1010101−2 (1 + 4 + 16 + 64 = 85) and 1110100−2 (4 + 16 − 32 + 64 = 52),

Carry:          1 −1  0 −1  1 −1  0  0  0
First addend:         1  0  1  0  1  0  1
Second addend:        1  1  1  0  1  0  0 +
               --------------------------
Number:         1 −1  2  0  3 −1  2  0  1
Bit (result):   1  1  0  0  1  1  0  0  1
Carry:          0  1 −1  0 −1  1 −1  0  0
so the result is 110011001−2 (1 − 8 + 16 − 128 + 256 = 137).

*/

func main() {
	fmt.Println(addNegabinary([]int{0}, []int{1, 0}))
	fmt.Println(addNegabinary([]int{0}, []int{0}))
	fmt.Println(addNegabinary([]int{0}, []int{1}))
	fmt.Println(addNegabinary([]int{1, 0, 1, 0, 1, 0, 1}, []int{1, 1, 1, 0, 1, 0, 0}))
	fmt.Println(convertDecimalToNegabinary(22))
	fmt.Println(addNegabinary([]int{1}, []int{1, 1, 0, 1}))
	fmt.Println(addNegabinaryV2([]int{1}, []int{1, 1, 0, 1}))

}

func addNegabinary(arr1 []int, arr2 []int) []int {
	ind1 := len(arr1) - 1
	ind2 := len(arr2) - 1
	carry := 0
	ans := make([]int, 0)
	m := map[int][]int{
		-2: {0, 1},
		-1: {1, 1},
		0:  {0, 0},
		1:  {1, 0},
		2:  {0, -1},
		3:  {1, -1},
	}
	for ind1 >= 0 || ind2 >= 0 || carry != 0 {
		val1 := 0
		val2 := 0
		if ind1 >= 0 {
			val1 = arr1[ind1]
			ind1--
		}

		if ind2 >= 0 {
			val2 = arr2[ind2]
			ind2--
		}
		number := val1 + val2 + carry
		bit := m[number][0]
		ans = append(ans, bit)
		carry = m[number][1]
	}
	// remove leading zeros
	ind := len(ans) - 1
	for len(ans) > 1 && ans[ind] == 0 {
		ans = ans[:len(ans)-1]
		ind--
	}
	reverse(ans)
	return ans
}

func addNegabinaryV2(arr1 []int, arr2 []int) []int {
	ind1 := len(arr1) - 1
	ind2 := len(arr2) - 1
	carry := 0
	ans := make([]int, 0)
	for ind1 >= 0 || ind2 >= 0 || carry != 0 {
		val1 := 0
		val2 := 0
		if ind1 >= 0 {
			val1 = arr1[ind1]
			ind1--
		}

		if ind2 >= 0 {
			val2 = arr2[ind2]
			ind2--
		}

		number := val1 + val2 + carry
		bit := number & 1
		ans = append(ans, bit)
		carry = -(number >> 1)
	}

	// remove leading zeros
	ind := len(ans) - 1
	for len(ans) > 1 && ans[ind] == 0 {
		ans = ans[:len(ans)-1]
		ind--
	}
	reverse(ans)
	return ans
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

/*
This func just use to support debug process.
*/
func convertDecimalToNegabinary(val int) []int {
	ans := make([]int, 0)
	for val != 0 {
		remainder := val % -2
		val = val / -2
		fmt.Println("Before: ", remainder, val)
		if remainder < 0 {
			remainder += 2
			val += 1
		}
		fmt.Println("After: ", remainder, val)
		fmt.Println("---------------------------")
		ans = append(ans, remainder)
	}
	reverse(ans)
	return ans
}
