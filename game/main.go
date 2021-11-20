package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hpdobrica/starbreach/state"
	"github.com/hpdobrica/starbreach/tiles"
)

type Game struct {
	Tiles           *tiles.Tiles
	SquaresInRow    int
	SquareSize      int
	StateMachine    *state.StateMachine
	NumberOfPlayers int
}

func NewGame(tiles *tiles.Tiles, squaresInRow, squareSize int) *Game {
	return &Game{
		Tiles:           tiles,
		SquaresInRow:    squaresInRow,
		SquareSize:      squareSize,
		NumberOfPlayers: 2,
	}
}

func (g *Game) Init(gameStateMachine *state.StateMachine) {
	g.StateMachine = gameStateMachine
}

const squaresInRow = 8
const squareSize = 50
const boardSize = squareSize * squaresInRow

func (g *Game) Update() error {
	(*g.StateMachine.CurrentState()).Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	(*g.StateMachine.CurrentState()).Draw(screen)

	// ebitenutil.DebugPrint(screen, "Hello world!")
	// drawBoard(screen, g.tiles)
	// tmpWriteNumbersOnBoard(screen)

	// drawPawn(screen, 1)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.SquareSize * g.SquaresInRow, g.SquareSize * g.SquaresInRow
}
