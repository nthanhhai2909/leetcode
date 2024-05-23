package main

import (
	"fmt"
	"sort"
)

// Link: https://leetcode.com/problems/heaters/
// Solution 1: Sort 2 array [houses, heaters] and then using 2-pointers
// N - Len houses, M - Len Heaters -> NLogN + MLogM + N, assume N ~ M => NLogN
func findRadius(houses []int, heaters []int) int {
	sort.Slice(houses, func(i, j int) bool { return houses[i] < houses[j] })
	sort.Slice(heaters, func(i, j int) bool { return heaters[i] < heaters[j] })
	houseLen := len(houses)
	heatersLen := len(heaters)
	ans := -1
	rHouse := houseLen - 1
	rHeater := heatersLen - 1
	for rHouse >= 0 {
		d := abs(houses[rHouse] - heaters[rHeater])
		tIndex := rHeater - 1
		for tIndex >= 0 {
			temp := abs(houses[rHouse] - heaters[tIndex])
			if temp <= d {
				rHeater--
				d = temp
			} else {
				break
			}
			tIndex--
		}
		if d > ans {
			ans = d
		}
		rHouse--
	}

	return ans
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func main() {
	fmt.Println(findRadius([]int{1, 2, 3, 5, 15}, []int{2, 30}), "13")
	fmt.Println(findRadius([]int{1, 5}, []int{10}), "9")
	fmt.Println(findRadius([]int{1, 1, 1, 1, 1, 1, 999, 999, 999, 999, 999}, []int{499, 500, 501}), "498")
	fmt.Println(findRadius([]int{282475249, 622650073, 984943658, 144108930, 470211272, 101027544, 457850878, 458777923},
		[]int{823564440, 115438165, 784484492, 74243042, 114807987, 137522503, 441282327, 16531729, 823378840, 143542612}), "161834419")
}
