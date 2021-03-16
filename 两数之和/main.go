package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum1([]int{3, 2, 4}, 6))
}

func twoSum(nums []int, target int) []int {

	var result []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i, j)
				return result
			}
		}
	}

	return result

}

func twoSum1(nums []int, target int) []int {
	hashMap := make(map[int]int, 0)
	for i, n := range nums {
		if n1, ok := hashMap[target-n]; ok {
			return []int{n1, i}
		}

		hashMap[n] = i
	}
	return nil
}
