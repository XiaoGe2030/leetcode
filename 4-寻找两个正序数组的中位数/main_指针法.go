package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3, 4, 9}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}

//没有利用nums1,nums2有序的条件
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)

	length := n + m
	k := length / 2
	if length&1 == 1 {
		return getKth(nums1, nums2, 0, n-1, 0, m-1, k+1)
	}
	return (getKth(nums1, nums2, 0, n-1, 0, m-1, k) + getKth(nums1, nums2, 0, n-1, 0, m-1, k+1)) * 0.5
}

func getKth(nums1, nums2 []int, start1, end1, start2, end2, k int) float64 {

	//计算数组长度
	len1 := end1 - start1 + 1
	len2 := end2 - start2 + 1

	if len1 > len2 {
		return getKth(nums2, nums1, start2, end2, start1, end1, k)
	}

	//如果第一个数组的长度为0,就直接返回第二个数组的第k大的元素就行
	if len1 == 0 {
		return float64(nums2[start2+k-1])
	}

	//k=1,时返回数组的最小元素
	if k == 1 {
		return float64(min(nums1[start1], nums2[start2]))
	}

	//这里主要考虑到数组的长度可能小于k/2,所以要去最小值
	i := start1 + min(len1, k/2) - 1
	j := start2 + min(len2, k/2) - 1

	if nums1[i] > nums2[j] {
		//切割nums2的元素,
		return getKth(nums1, nums2, start1, end1, j+1, end2, k-(j-start2+1))
	} else {
		return getKth(nums1, nums2, i+1, end1, start2, end2, k-(i-start1+1))
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
