package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 1, 3, 1}

	fmt.Println(removeDublicatesElement(s))
}

func removeDublicatesElement(nums []int) []int {
	result := make([]int, 0, len(nums))
	temp := map[int]struct{}{}
	for _, item := range nums {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
