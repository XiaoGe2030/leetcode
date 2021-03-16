package main

import "sort"

func main() {
	//	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	merge([]int{2, 0}, 1, []int{1}, 1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m; i < len(num1); i++ {
		nums[i] = nums2[i-m]
	}

	sort.Ints(nums1)
}
