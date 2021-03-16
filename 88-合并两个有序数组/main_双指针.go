package main

import "fmt"

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	merge([]int{2, 0}, 1, []int{1}, 1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	i, j := m-1, n-1
	for i >= 0 && j >= 0 {
		n1 := nums1[i]
		n2 := nums2[j]

		if n1 < n2 {
			nums1[i+j+1] = n2
			j--
		} else {
			nums1[i+j+1] = n1
			i--
		}
	}

	if j >= 0 {
		for k := j; k >= 0; k-- {
			nums1[k] = nums2[k]
		}
	}

	fmt.Println(nums1)
}
