package main

import (
	"fmt"
)

type scenario struct {
	Expected []int
	Input    []int
}

func compareSlices(sl1, sl2 []int) bool {
	if len(sl1) != len(sl2) {
		return false
	}
	for i, v := range sl1 {
		if v != sl2[i] {
			return false
		}
	}
	return true
}

func sort(input []int) []int {
	for i := 0; i < len(input)-1; i++ {
		for j := 0; j < len(input)-1; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}
	return input
}

func main() {
	var scenarios = []*scenario{
		{
			Expected: []int{1, 2, 3, 4, 5},
			Input:    []int{5, 4, 3, 2, 1},
		},
		{
			Expected: []int{-4, -2, 1, 3, 5},
			Input:    []int{-2, 1, 5, -4, 3},
		},
		{
			Expected: []int{21, 43,65, 89, 100},
			Input:    []int{21, 89, 65, 43, 100},
		},
	}
	for _, s := range scenarios {
		actual := sort(s.Input)
		if !compareSlices(actual, s.Expected) {
			fmt.Println("test failed")
			return
		}
	}
	fmt.Println("tests passed")
}
