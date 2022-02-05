package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 1, 3, 1}
	fmt.Println(removeDublicatesElement3(s))
}

func removeDublicatesElement3(nums []int) []int {
	result := make([]int, 0, len(nums))

	for _, item := range nums {
		found := false
		for _, r := range result {
			if item == r {
				found = true
				break
			}
		}
		if !found {
			result = append(result, item)
		}

	}

	return result
}

func removeDublicatesElement2(nums []int) []int {
	result := make([]int, 0, len(nums))

	temp := map[int]struct{}{}
	for _, item := range nums {
		temp[item] = struct{}{}
	}
	for k := range temp {
		result = append(result, k)
	}
	return result
}

func removeDublicatesElement(nums []int) []int {
	result := make([]int, 0, len(nums))
	fmt.Println(len(result))
	fmt.Println(cap(result))

	temp := map[int]struct{}{}
	for _, item := range nums {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
