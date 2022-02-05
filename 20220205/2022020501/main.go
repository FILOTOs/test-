package main

import "fmt"

type stack struct {
	slc []int
}

func newStack() *stack {
	return &stack{
		slc: []int{},
	}
}

func (s *stack) push(item int) {
	fmt.Println("push")
	s.slc = append(s.slc, item)
}

func (s *stack) print() {
	fmt.Println("stack: ", s.slc)
}

func (s *stack) pop() int {
	fmt.Println("pop")
	if len(s.slc) == 0 {
		fmt.Println("stack is empty")
		return 0
	}
	res := s.slc[len(s.slc)-1]
	s.slc = s.slc[:len(s.slc)-1]
	return res
}

func main() {
	stc := newStack()
	stc.push(1)
	stc.print()
	stc.push(2)
	stc.print()
	stc.push(3)
	stc.print()
	fmt.Println(stc.pop())
	stc.print()
	fmt.Println(stc.pop())
	stc.print()
	fmt.Println(stc.pop())
	stc.print()
	stc.pop()
}
