package main

import (
	"fmt"
)

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

func (s *stack) pop() (int, error) {
	fmt.Println("pop")
	if len(s.slc) == 0 {
		return -1, fmt.Errorf("pop error: stack is empty")
	}
	res := s.slc[len(s.slc)-1]
	s.slc = s.slc[:len(s.slc)-1]
	return res, nil
}

func main() {
	stc := newStack()
	stc.push(1)
	stc.print()
	stc.push(2)
	stc.print()
	stc.push(3)
	stc.print()
	val, err := stc.pop()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	stc.print()
	val, err = stc.pop()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	stc.print()
	val, err = stc.pop()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	stc.print()
	val, err = stc.pop()
	if err != nil {
		fmt.Println(err)
		return
	}
}
