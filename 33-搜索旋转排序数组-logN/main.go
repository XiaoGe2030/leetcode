package main

import "fmt"

func main() {
	//fmt.Println(search([]int{4,5,6,7,0,1,2},0))
//	fmt.Println(search([]int{1,3},3))
	fmt.Println(search([]int{4,5,6,7,8,1,2,3},8))
}

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	if n == 1 {
		if nums[0]== target {
			return 0
		}
		return -1 
	}

	for l,r := 0,n-1 ;l <= r ; {
		mid := (l+r)/2
		fmt.Println(l,r,mid)
		if nums[mid]==target {
			return mid
		}

		if nums[0] <=  nums[mid]{ 
			//如果左边有序
			if nums[0] <= target && target <=  nums[mid] {
				r = mid -1 
			}else {
				l = mid +1
			}
		}else {
			if  nums[mid] < target &&  target<= nums[n-1] {
				l = mid +1 
			}else {
				r = mid -1
			}
		}
	}

	return -1 

}
