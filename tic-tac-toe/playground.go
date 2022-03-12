package main

import "fmt"

type playground struct {
	Board [3][3]int
}

// Init - initializes a new play
func (p *playground) Init() {
	// init empty board
	for i, r := range p.Board {
		for j, _ := range r {
			p.Board[i][j] = 0
		}
	}
}

func (p *playground) Print() {
	for i, r := range p.Board {
		for j, _ := range r {
			if p.Board[i][j] == 0 {
				fmt.Print("_")
			}
			if p.Board[i][j] == 1 {
				fmt.Print("X")
			}
			if p.Board[i][j] == 2 {
				fmt.Print("O")
			}
		}
		fmt.Println()
	}

}

func NewPlayground() *playground {
	return &playground{}
}

func (p *playground) SetMove(i, j int, commandType uint32) bool {
	if p.Board[i][j] == 0 {
		p.Board[i][j] = int(commandType)
		return true
	}
	return false
}
