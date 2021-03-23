package main

import "fmt"

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

func maxSlidingWindow(nums []int, k int) []int {
	q := []int{}
	push := func(i int) {
		//nums[i]大于队列的最后一位元素，就弹出最后的元素
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		//直到最后，不小，就把i索引添加到q的数组中
		q = append(q, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q[0]]
	//i代表滑动窗口最右边的值
	for i := k; i < n; i++ {
		push(i)
		//这个有点不好理解,排除不属于当前滑动窗口值的索引
		for q[0] <= i-k {
			q = q[1:]
		}
		//q队列对首的元素是最大元素
		ans = append(ans, nums[q[0]])
	}
	return ans
}
