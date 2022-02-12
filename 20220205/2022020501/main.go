package main

import "fmt"

func checkFormula(f string) bool {
	stc := newStack()
	symbolCount := 0
	for _, r := range f {
		if r == '(' {
			stc.push(0)
			symbolCount = 0
			continue
		}
		if r == ')' {
			_, err := stc.pop()
			if err != nil || symbolCount == 0 {
				return false
			}
			continue
		}
		symbolCount++
	}
	return true
}

func main() {
	tests := []string{
		"((1+2)*(3+4)*(5+6))*(2+3)",
		"(1+2))",
		"))1*2",
		"((1+2)*(3+4)*(5+6))*(2+3))))",
		"(()(2+3)",
		"(2+3)",
	}
	for _, t := range tests {
		fmt.Println(t, ":", checkFormula(t))
	}
}
