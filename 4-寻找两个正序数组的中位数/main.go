package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}

//没有利用nums1,nums2有序的条件
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	n := len(nums1)

	if n&1 == 0 {
		return (float64(nums1[n/2]+nums1[n/2-1]) / 2.0)
	} else {
		return float64(nums1[n/2])
	}
}
