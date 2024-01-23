package main

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
