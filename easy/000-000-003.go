package main

import "fmt"

/*
*

Leetcode: https://leetcode.com/problems/c/
You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

Example 1:

Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
Example 2:

Input: nums1 = [1], m = 1, nums2 = [], n = 0
Output: [1]
Explanation: The arrays we are merging are [1] and [].
The result of the merge is [1].
Example 3:

Input: nums1 = [0], m = 0, nums2 = [1], n = 1
Output: [1]
Explanation: The arrays we are merging are [] and [1].
The result of the merge is [1].
Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge result can fit in nums1.

Constraints:

nums1.length == m + n
nums2.length == n
0 <= m, n <= 200
1 <= m + n <= 200
-10^9 <= nums1[i], nums2[j] <= 10^9

Follow up: Can you come up with an algorithm that runs in O(m + n) time?
*/
func main() {
	nums1 := []int{1, 4, 6, 7, 8, 10, 0, 0, 0, 0}
	nums2 := []int{1, 2, 7, 9}
	mergeV2(nums1, 6, nums2, 4)
	fmt.Println(nums1)
}

func mergeV1(nums1 []int, m int, nums2 []int, n int) {
	p1 := 0
	p2 := 0
	ans := make([]int, m+n)
	for p1 < m && p2 < n {
		if nums1[p1] <= nums2[p2] {
			ans[p1+p2] = nums1[p1]
			p1++
		} else {
			ans[p1+p2] = nums2[p2]
			p2++
		}
	}

	for p1 < m {
		ans[p1+p2] = nums1[p1]
		p1++
	}

	for p2 < n {
		ans[p1+p2] = nums2[p2]
		p2++
	}

	for i := 0; i < m+n; i++ {
		nums1[i] = ans[i]
	}
}

func mergeV2(nums1 []int, m int, nums2 []int, n int) {
	ind1 := m - 1
	ind2 := n - 1
	ind3 := m + n - 1

	for ind3 >= 0 && ind1 >= 0 && ind2 >= 0 {
		if nums1[ind1] >= nums2[ind2] {
			nums1[ind3] = nums1[ind1]
			ind1--
		} else {
			nums1[ind3] = nums2[ind2]
			ind2--
		}

		ind3--
	}

	for ind1 >= 0 {
		nums1[ind3] = nums1[ind1]
		ind1--
		ind3--
	}

	for ind2 >= 0 {
		nums1[ind3] = nums2[ind2]
		ind2--
		ind3--
	}
}
