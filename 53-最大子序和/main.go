package main

<<<<<<< HEAD
import (
	"fmt"
)
=======
import "fmt"
>>>>>>> 20400e72f6b5c14362dc03729df340c899d1cd9f

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
<<<<<<< HEAD
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
=======
		//这里求出索引i结尾的最大数
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		//索引i为结尾的最大数，跟当前最大值得对比
>>>>>>> 20400e72f6b5c14362dc03729df340c899d1cd9f
		if nums[i] > max {
			max = nums[i]
		}
	}
<<<<<<< HEAD
=======

>>>>>>> 20400e72f6b5c14362dc03729df340c899d1cd9f
	return max
}
