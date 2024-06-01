package main

import "fmt"

// Link: https://leetcode.com/problems/minimum-penalty-for-a-shop/

// TC: O(n) SC: O(n)
func bestClosingTime(customers string) int {
	size := len(customers)
	ans := size
	noArr := make([]int, size+1)
	yesArr := make([]int, size+1)
	noCount := 0
	yesCount := 0
	for i := 0; i < size; i++ {
		noArr[i] = noCount
		if customers[i] == 'N' {
			noCount++
		}
	}
	noArr[size] = noArr[size-1]
	for i := size - 1; i >= 0; i-- {
		if customers[i] == 'Y' {
			yesCount++
		}
		yesArr[i] = yesCount
	}
	yesArr[size] = yesArr[size-1] - 1
	if yesArr[size] < 0 {
		yesArr[size] = 0
	}
	cur := size
	for i := 0; i <= size; i++ {
		sum := noArr[i] + yesArr[i]
		if sum < cur {
			ans = i
			cur = sum
		}
	}

	return ans
}

func main() {
	fmt.Println(bestClosingTime("YNYYN"))
}
