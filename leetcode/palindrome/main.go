package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	r := 0
	y := x
	for x > 0 {
		r = r*10 + x%10
		x = x/10
	}
	return r == y
}

type scenario struct {
	Input    int
	Expected bool
}

func main() {
	var scenarios = []scenario{
		{
			Input:    232,
			Expected: true,
		},
		{
			Input:    -323,
			Expected: false,
		},
		{
			Input:    284653,
			Expected: false,
		},
		{
			Input:    1,
			Expected: true,
		},
		{
			Input:    123456789987654321,
			Expected: true,
		},
	}
	for _, s := range scenarios {
		if isPalindrome(s.Input) != s.Expected {
			fmt.Println("test failed: ", s.Input)
			return
		}
	}
	fmt.Println("tests passed")
}
