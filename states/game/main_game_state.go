package gamestates

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hpdobrica/starbreach/board"
	"github.com/hpdobrica/starbreach/game"
	"github.com/hpdobrica/starbreach/state"
	turnstates "github.com/hpdobrica/starbreach/states/turn"
)

type MainGameState struct {
	game             *game.Game
	boards           []*board.Board
	stateMachine     *state.StateMachine
	turnStateMachine *state.StateMachine
}

func NewMainGameState(game *game.Game) *MainGameState {
	boards := make([]*board.Board, game.NumberOfPlayers)

	boards[0] = board.NewBoard(game, 0, 2, 1, game.SquareSize, game.SquaresInRow, game.Tiles)
	boards[1] = board.NewBoard(game, 2, 2, 2, game.SquareSize, game.SquaresInRow, game.Tiles)

	turnStateMachine := state.NewStateMachine()

	return &MainGameState{game: game, boards: boards, turnStateMachine: turnStateMachine}
}

func (m *MainGameState) Init(stateMachine *state.StateMachine) {
	m.stateMachine = stateMachine

	turnTransition := turnstates.NewTurnTransitionState(m.game, 0, 1)

	m.turnStateMachine.PushState(turnTransition)

}

func (m *MainGameState) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyJ) {
		// m.currentBoard = m.boards[]
	}
	(*m.turnStateMachine.CurrentState()).Update()

}

func (m *MainGameState) Draw(screen *ebiten.Image) {
	// drawBoard(screen, m.game.SquareSize, m.game.SquaresInRow, m.game.Tiles)
	m.boards[1].Draw(screen)
	(*m.turnStateMachine.CurrentState()).Draw(screen)

}

// func drawBoard(screen *ebiten.Image, squareSize int, squaresInRow int, tilesObj *tiles.Tiles) {

// 	scale := float64(squareSize) / float64(tilesObj.TileSize)

// 	for i := 0; i < squaresInRow; i++ {
// 		for j := 0; j < squaresInRow; j++ {
// 			op := &ebiten.DrawImageOptions{}
// 			op.GeoM.Scale(scale, scale)
// 			op.GeoM.Translate(float64(i*squareSize), float64(j*squareSize))

// 			screen.DrawImage(tilesObj.GetTileAt(0, 2), op)

// 		}
// 	}

// }
