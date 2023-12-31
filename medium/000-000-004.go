package main

import "fmt"

/*
Leetcode: https://leetcode.com/problems/powx-n/
Implement pow(x, n), which calculates x raised to the power n (i.e., xn).
*/

func myPowV1(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / myPowV1(x, -n)
	}
	return x * myPowV1(x, n-1)
}

func myPowV2(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / myPowV2(x, -n)
	}
	if n%2 == 0 {
		return myPowV2(x*x, n/2)
	} else {
		return x * myPowV2(x*x, (n-1)/2)
	}
}

func main() {
	fmt.Println(myPowV1(2.00000, -10))
	fmt.Println(myPowV2(2.00000, -10))
	fmt.Println("-------------------------------------")
	fmt.Println(myPowV1(2.00000, 10))
	fmt.Println(myPowV2(2.00000, 10))
	fmt.Println("-------------------------------------")
	fmt.Println(myPowV1(5.00000, 7))
	fmt.Println(myPowV2(5.00000, 7))
	fmt.Println("-------------------------------------")
	fmt.Println(myPowV1(5.00000, 20))
	fmt.Println(myPowV2(5.00000, 20))
}
