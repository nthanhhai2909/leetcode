package main

import "fmt"

/*
Leetcode: https://leetcode.com/problems/find-in-mountain-array/

You may recall that an array arr is a mountain array if and only if:

arr.length >= 3
There exists some i with 0 < i < arr.length - 1 such that:
arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
Given a mountain array mountainArr, return the minimum index such that mountainArr.get(index) == target. If such an index does not exist, return -1.

You cannot access the mountain array directly. You may only access the array using a MountainArray interface:

MountainArray.get(k) returns the element of the array at index k (0-indexed).
MountainArray.length() returns the length of the array.
Submissions making more than 100 calls to MountainArray.get will be judged Wrong Answer. Also, any solutions that attempt to circumvent the judge will result in disqualification.



Example 1:

Input: array = [1,2,3,4,5,3,1], target = 3
Output: 2
Explanation: 3 exists in the array, at index=2 and index=5. Return the minimum index, which is 2.
Example 2:

Input: array = [0,1,2,4,2,1], target = 3
Output: -1
Explanation: 3 does not exist in the array, so we return -1.


Constraints:

3 <= mountain_arr.length() <= 104
0 <= target <= 109
0 <= mountain_arr.get(index) <= 109
*/

type MountainArray struct {
	arr   []int
	count int
}

func NewMountainArray(arr []int) *MountainArray {
	return &MountainArray{
		arr:   arr,
		count: 0,
	}
}

var m map[int]int

func (this *MountainArray) get(index int) int {
	this.count++
	return this.arr[index]
}

func (this *MountainArray) length() int {
	return len(this.arr)
}

// -------------------------------------------------------- Version 1 ---------------------------------------------------
func binarySearch(mountainArray *MountainArray, start int, end int, target int, compare func(i, j int) int) int {
	index := mountainArray.length()
	for start <= end {
		mid := (start + end) / 2
		sort := compare(get(mid, mountainArray), target)
		if sort < 0 {
			start = mid + 1
		} else if sort > 0 {
			end = mid - 1
		} else {
			return mid
		}
	}
	return index
}

func findInMountainArray(target int, mountainArr *MountainArray) int {
	m = make(map[int]int)
	ans := doSearch(0, mountainArr.length()-1, target, mountainArr)
	if ans == mountainArr.length() {
		return -1
	}
	return ans
}

func doSearch(start, end, target int, mountainArr *MountainArray) int {
	ans := mountainArr.length()
	mid := (start + end) / 2
	if mid == 0 {
		return min(binarySearch(mountainArr, start, end, target, less), ans)
	}

	if mid == mountainArr.length()-1 {
		return min(binarySearch(mountainArr, start, end, target, bigger), ans)
	}

	midVal := get(mid, mountainArr)
	if mid > 0 && midVal >= get(mid-1, mountainArr) {
		if target <= midVal {
			ans = min(binarySearch(mountainArr, start, mid, target, less), ans)
		}
	} else {
		ans = min(doSearch(start, mid-1, target, mountainArr), ans)
	}

	if mid < mountainArr.length()-1 && midVal >= get(mid+1, mountainArr) {
		if target <= midVal {
			ans = min(binarySearch(mountainArr, mid, end, target, bigger), ans)
		}
	} else {
		ans = min(doSearch(mid+1, end, target, mountainArr), ans)
	}
	return ans
}

func less(i, j int) int {
	return i - j
}

func bigger(i, j int) int {
	return j - i
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func get(ind int, mountainArr *MountainArray) int {
	if val, ok := m[ind]; ok {
		return val
	}
	val := mountainArr.get(ind)
	m[ind] = val
	return val
}

// ---------------------------------------------------------------------------------------------------------------------

// -------------------------------------------------------- Version 2 --------------------------------------------------
func binarySearchV2(mountainArray *MountainArray, start int, end int, target int, compare func(i, j int) int) int {
	index := mountainArray.length()
	for start <= end {
		mid := (start + end) / 2
		sort := compare(mountainArray.get(mid), target)
		if sort < 0 {
			start = mid + 1
		} else if sort > 0 {
			end = mid - 1
		} else {
			return mid
		}
	}
	return index
}

func findInMountainArrayV2(target int, mountainArr *MountainArray) int {
	l := 0
	r := mountainArr.length() - 1
	m := 0
	for l <= r {
		m = (l + r) / 2
		mV := mountainArr.get(m)
		rV := mountainArr.get(m + 1)

		if mV <= rV {
			l = m
		}
		if mV >= rV {
			r = m
		}
		if r == l+1 {
			m = r
			break
		}
	}

	ans1 := binarySearchV2(mountainArr, 0, m, target, less)
	if ans1 != mountainArr.length() {
		return ans1
	}
	ans2 := binarySearchV2(mountainArr, m+1, mountainArr.length()-1, target, bigger)
	if ans2 != mountainArr.length() {
		return ans2
	}
	return -1
}

func main() {
	arr := NewMountainArray([]int{1, 2, 3, 4, 5, 3, 1})
	fmt.Println(findInMountainArrayV2(3, arr))
}

// ---------------------------------------------------------------------------------------------------------------------
