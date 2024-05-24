package main

import "fmt"

func twoSum(nums []int, target int) []int {
	indices := make([]int, 0)
	numMap := make(map[int]int)

	for idx, num := range nums {
		diff := target - num
		val, existsInMap := numMap[diff]
		if existsInMap {
			indices = append(indices, idx, val)
			return indices
		} else {
			numMap[num] = idx
		}
		numMap[num] = idx
	}

	return indices
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	indices := twoSum(nums, target)
	fmt.Println(indices)
}
