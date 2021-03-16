package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum1([]int{-2, 0, 1, 1, 2}))
}

/*
给你一个包含n个整数的数组nums,判断nums中是否存在三个元素a,b,c,使得a+b+c=0?请找出所有和为0且不重复的三元组
答案中不可以包含重复的三元组
*/

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

func threeSum1(nums []int) [][]int {

	sort.Ints(nums)
	fmt.Println(nums)
	var result [][]int
	n := len(nums)
	//这里取巧，减少了两次运算
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		right := n - 1
		left := i + 1
		target := -nums[i]

		for left < right {
			sum := nums[left] + nums[right]
			if target == sum {
				//有解的情况下，如果不去重，重复的left，必将找出重复的right
				result = append(result, []int{nums[i], nums[left], nums[right]})
				//去重
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				//去重后需要偏移指针
				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
