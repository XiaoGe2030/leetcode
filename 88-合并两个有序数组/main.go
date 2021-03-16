package main

import "fmt"

func main() {
	//	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	merge([]int{2, 0}, 1, []int{1}, 1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	var res []int
	var i, j int
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}

	if i < m {
		res = append(res, nums1[i:]...)
	}

	if j < n {
		res = append(res, nums2[j:]...)
	}

	for i := 0; i < len(nums1); i++ {
		nums1[i] = res[i]
	}

	fmt.Println(nums1)
}
