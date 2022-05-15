package main

import (
	"fmt"
)

type scenario struct {
	Expected int
	Input    []int
	Target   int
}

func searching(tr int, inp []int, offset int) int {
	if len(inp) == 1 {
		if inp[0] == tr {
			return 0
		} else {
			return -1
		}
	}
	mid := len(inp) / 2
	if inp[mid] > tr {
		return searching(tr, inp[:mid], offset)
	} else if inp[mid] == tr {
		return mid + offset
	} else {
		return searching(tr, inp[mid:], mid+offset)
	}
}

func main() {
	var scenarios = []scenario{
		{
			Expected: 3,
			Input:    []int{-2, 1, 5, 12, 35},
			Target:   12,
		},
		{
			Expected: 6,
			Input:    []int{0, 1, 3, 4, 7, 9, 12},
			Target:   12,
		},
		{
			Expected: 2,
			Input:    []int{-3, 2, 5, 8, 9, 18, 26, 30},
			Target:   5,
		},
		{
			Expected: -1,
			Input:    []int{2, 10, 16, 35, 54},
			Target:   25,
		},
	}
	for _, s := range scenarios {
		actual := searching(s.Target, s.Input, 0)
		if s.Expected != actual {
			fmt.Println("tests failed")
			return
		}
	}
	fmt.Println("tests passed")
}
