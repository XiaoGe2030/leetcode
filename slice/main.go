package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 4}
	nums2 := make([]int, 4)
	copy(nums2, nums1)
	fmt.Println(nums2)
}
