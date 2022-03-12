package main

import "fmt"

type Game struct {
	Playground       *playground
	Player1, Player2 *player
}

func NewGame() *Game {
	g := &Game{
		Playground: NewPlayground(),
		Player1:    NewPlayer(1),
		Player2:    NewPlayer(2),
	}
	g.Playground.Init()
	g.Player1.RequestName()
	g.Player2.RequestName()
	return g
}

func (g *Game) move(p *player) {
	for {
		i, j := p.RequestMove()
		if g.Playground.SetMove(i, j, p.CommandType) {
			break
		} else {
			fmt.Println("incorrect move, make another move")
		}
	}
	g.Playground.Print()
}

func (g *Game) Start() {
	for {
		g.move(g.Player1)
		if g.checkGame() {
			return
		}
		g.move(g.Player2)
		if g.checkGame() {
			return
		}
	}
}

func (g *Game) checkGame() bool {
	for i := 0; i < 3; i++ {
		win := true
		for j := 0; j < 3; j++ {
			if g.Playground.Board[i][j] != 1 {
				win = false
				break
			}
		}
		if win {
			fmt.Println("Player1 won")
			return true
		}
	}
	for i := 0; i < 3; i++ {
		win := true
		for j := 0; j < 3; j++ {
			if g.Playground.Board[j][i] != 1 {
				win = false
				break
			}
		}
		if win {
			fmt.Println("Player1 won")
			return true
		}
	}
	for i := 0; i < 3; i++ {
		win := true
		for j := 0; j < 3; j++ {
			if g.Playground.Board[i][j] != 2 {
				win = false
				break
			}
		}
		if win {
			fmt.Println("Player2 won")
			return true
		}
	}
	for i := 0; i < 3; i++ {
		win := true
		for j := 0; j < 3; j++ {
			if g.Playground.Board[j][i] != 2 {
				win = false
				break
			}
		}
		if win {
			fmt.Println("Player2 won")
			return true
		}
	}
	for i := 0; i < 3; i++ {
		win := true
		if g.Playground.Board[i][i] != 1 {
			win = false
			break
		}
		if win {
			fmt.Println("Player1 won")
			return true
		}
	}
	for i := 0; i < 3; i++ {
		win := true
		if g.Playground.Board[i][2-i] != 1 {
			win = false
			break
		}
		if win {
			fmt.Println("Player1 won")
			return true
		}
	}
	return false
}
