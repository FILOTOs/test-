package main

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

func (g *Game) Start() {
	for {
		i, j := g.Player1.RequestMove()
		g.Playground.SetMove(i, j, g.Player1.CommandType)
		g.Playground.Print()
	}
}
