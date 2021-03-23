package main

import "fmt"

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow([]int{1, -1}, 1))
}

func maxSlidingWindow(nums []int, k int) []int {
	var s []int
	push := func(i int) {
		for len(s) > 0 && nums[i] > nums[s[len(s)-1]] {
			s = s[:len(s)-1]
		}

		s = append(s, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	res := make([]int, n-k+1)
	res[0] = nums[s[0]]
	for i := k; i < n; i++ {
		push(i)

		for s[0] <= i-k {
			s = s[1:]
		}
		//res = append(res, nums[s[0]])
		res[i-k+1] = nums[s[0]]
	}
	return res
}
