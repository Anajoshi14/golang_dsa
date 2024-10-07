package main

import "fmt"

func main() {

	s1 := twoSum([]int{2, 11, 15, 7}, 9)
	fmt.Println(s1)
}
func twoSum(nums []int, target int) []int {

	//step1 - create a map with keys as array elements and values as index

	mapElements := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		reqValue := target - nums[i]
		reqIndex, isExist := mapElements[reqValue]
		if isExist {
			return []int{i, reqIndex}
		}

		mapElements[nums[i]] = i
	}
	return []int{}
}
