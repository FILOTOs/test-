package main

import "fmt"

type Symbol uint8

const (
	Cross Symbol = iota
	Circle
	Dot
)

type GameStatus uint8

const (
	GameInProgress GameStatus = iota
	GameDraw
	FirstPlayerWin
	SecondPlayerWin
)

type iPlayer interface {
	getSymbol() Symbol
	getNextMove() (int, int, error)
	getID() int
}

var (
	MovesPlayer1 = [4][2]int{{1, 1}, {2, 0}, {2, 2}, {2, 1}}
	MovesPlayer2 = [4][2]int{{1, 2}, {0, 2}, {0, 0}, {0, 0}}
)

type humanPlayer struct {
	symbol Symbol
	index  int
	id     int
}

func (h *humanPlayer) getSymbol() Symbol {
	return h.symbol
}

func (h *humanPlayer) getNextMove() (int, int, error) {
	if h.symbol == Cross {
		h.index = h.index + 1
		return MovesPlayer1[h.index-1][0], MovesPlayer1[h.index-1][1], nil
	} else if h.symbol == Circle {
		h.index = h.index + 1
		return MovesPlayer2[h.index-1][0], MovesPlayer2[h.index-1][1], nil
	}
	return 0, 0, fmt.Errorf("Invalid Symbol")
}

func (h *humanPlayer) getID() int {
	return h.id
}

type computerPlayer struct {
	symbol Symbol
	id     int
}

func (c *computerPlayer) getSymbol() Symbol {
	return c.symbol
}

func (c *computerPlayer) getNextMove() (int, int, error) {
	//To be implemented
	return 0, 0, nil
}

func (c *computerPlayer) getID() int {
	return c.id
}

type board struct {
	square    [][]Symbol
	dimension int
}

func (b *board) printBoard() {
	for i := 0; i < b.dimension; i++ {
		for j := 0; j < b.dimension; j++ {
			if b.square[i][j] == Dot {
				fmt.Print(".")
			} else if b.square[i][j] == Cross {
				fmt.Print("*")
			} else {
				fmt.Print("o")
			}

		}
		fmt.Println("")
	}
}

func (b *board) markSymbol(i, j int, symbol Symbol) (bool, Symbol, error) {
	if i > b.dimension || j > b.dimension {
		return false, Dot, fmt.Errorf("index input is greater than dimension")
	}
	if b.square[i][j] != Dot {
		return false, Dot, fmt.Errorf("input square already marked")
	}
	if symbol != Cross && symbol != Circle {
		return false, Dot, fmt.Errorf("incorrect Symbol")
	}
	b.square[i][j] = symbol
	win := b.checkWin(i, j, symbol)

	return win, symbol, nil
}

func (b *board) checkWin(i, j int, symbol Symbol) bool {
	//Check Row
	rowMatch := true
	for k := 0; k < b.dimension; k++ {
		if b.square[i][k] != symbol {
			rowMatch = false
		}
	}

	if rowMatch {
		return rowMatch
	}

	//Check Row
	columnMatch := true
	for k := 0; k < b.dimension; k++ {
		if b.square[k][j] != symbol {
			columnMatch = false
		}
	}

	if columnMatch {
		return columnMatch
	}

	//Check diagonal
	diagonalMatch := false
	if i == j {
		diagonalMatch = true
		for k := 0; k < b.dimension; k++ {
			if b.square[k][k] != symbol {
				diagonalMatch = false
			}
		}
	}

	return diagonalMatch

}

type game struct {
	board           *board
	firstPlayer     iPlayer
	secondPlayer    iPlayer
	firstPlayerTurn bool
	moveIndex       int
	gameStatus      GameStatus
}

func initGame(b *board, p1, p2 iPlayer) *game {
	game := &game{
		board:           b,
		firstPlayer:     p1,
		secondPlayer:    p2,
		firstPlayerTurn: true,
		gameStatus:      GameInProgress,
	}
	return game
}

func (g *game) play() error {
	var win bool
	var symbol Symbol
	for {
		if g.firstPlayerTurn {
			x, y, err := g.firstPlayer.getNextMove()
			if err != nil {
				return err
			}
			win, symbol, err = g.board.markSymbol(x, y, g.firstPlayer.getSymbol())
			if err != nil {
				return err
			}
			g.firstPlayerTurn = false
			g.printMove(g.firstPlayer, x, y)
		} else {
			x, y, err := g.secondPlayer.getNextMove()
			if err != nil {
				return err
			}
			win, symbol, err = g.board.markSymbol(x, y, g.secondPlayer.getSymbol())
			if err != nil {
				return err
			}
			g.firstPlayerTurn = true
			g.printMove(g.secondPlayer, x, y)
		}
		g.moveIndex = g.moveIndex + 1

		g.setGameStatus(win, symbol)
		if g.gameStatus != GameInProgress {
			break
		}
	}
	return nil
}

func (g *game) setGameStatus(win bool, symbol Symbol) {
	if win {
		if g.firstPlayer.getSymbol() == symbol {
			g.gameStatus = FirstPlayerWin
			return
		} else if g.secondPlayer.getSymbol() == symbol {
			g.gameStatus = SecondPlayerWin
			return
		}
	}
	if g.moveIndex == g.board.dimension*g.board.dimension {
		g.gameStatus = GameDraw
		return
	}
	g.gameStatus = GameInProgress
}

func (g *game) printMove(player iPlayer, x, y int) {
	symbolString := ""
	symbol := player.getSymbol()
	if symbol == Cross {
		symbolString = "*"
	} else if symbol == Circle {
		symbolString = "o"
	}
	fmt.Printf("Player %d Move with Symbol %s at Position X:%d Y:%d\n", player.getID(), symbolString, x, y)
	g.board.printBoard()
	fmt.Println("")
}

func (g *game) printResult() {
	switch g.gameStatus {
	case GameInProgress:
		fmt.Println("Game in Between")
	case GameDraw:
		fmt.Println("Game Drawn")
	case FirstPlayerWin:
		fmt.Println("First Player Win")
	case SecondPlayerWin:
		fmt.Println("Second Player Win")
	default:
		fmt.Println("Invalid Game Status")
	}
	g.board.printBoard()
}

func main() {

	board := &board{
		square:    [][]Symbol{{Dot, Dot, Dot}, {Dot, Dot, Dot}, {Dot, Dot, Dot}},
		dimension: 3,
	}

	player1 := &humanPlayer{
		symbol: Cross,
		id:     1,
	}

	player2 := &humanPlayer{
		symbol: Circle,
		id:     2,
	}

	game := initGame(board, player1, player2)
	game.play()
	game.printResult()

}
