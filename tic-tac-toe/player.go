package main

import "fmt"

type player struct {
	CommandType uint32
	Name        string
}

func NewPlayer(commandType uint32) *player {
	return &player{
		CommandType: commandType,
	}
}

func (p *player) RequestName() {
	fmt.Printf("input name for Player %d: ", p.CommandType)
	fmt.Scanf("%s\n", &p.Name)
}

func (p *player) RequestMove() (int, int) {
	i, j := 0, 0
	fmt.Printf("Player%d. input move: ", p.CommandType)
	fmt.Scanf("%d\n", &i)
	fmt.Scanf("%d\n", &j)
	return i, j
}
