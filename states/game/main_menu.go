package gamestates

import (
	"fmt"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hpdobrica/starbreach/game"
	"github.com/hpdobrica/starbreach/state"
	"github.com/hpdobrica/starbreach/tiles"
)

type MainMenuState struct {
	game         *game.Game
	background   *ebiten.Image
	stateMachine *state.StateMachine
}

func NewMainMenuState(game *game.Game) *MainMenuState {
	return &MainMenuState{game: game}
}

func (m *MainMenuState) Init(stateMachine *state.StateMachine) {
	m.stateMachine = stateMachine
}

func (i *MainMenuState) Update() {

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if (x >= 25 && x <= 95) && (y >= 25 && y <= 50) {
			fmt.Println("New Game!")
			mainGameState := NewMainGameState(i.game)
			i.game.StateMachine.PushState(mainGameState)
		}

		if (x >= 25 && x <= 95) && (y >= 55 && y <= 80) {
			fmt.Println("Quit!")
			os.Exit(0)
		}

		fmt.Println(x, y)
	}

}

func (i *MainMenuState) Draw(screen *ebiten.Image) {
	if i.background == nil {
		w, h := screen.Size()
		i.background = ebiten.NewImage(w, h)
	}

	i.background.Fill(color.Black)

	op := &ebiten.DrawImageOptions{}

	screen.DrawImage(i.background, op)

	ebitenutil.DebugPrintAt(screen, "New Game", 30, 30)

	ebitenutil.DebugPrintAt(screen, "Quit", 30, 60)

	// drawMainMenu(screen, i.game.SquareSize, i.game.SquaresInRow, i.game.Tiles)

}

func drawMainMenu(screen *ebiten.Image, squareSize int, squaresInRow int, tilesObj *tiles.Tiles) {

	scale := float64(squareSize) / float64(tilesObj.TileSize)

	for i := 0; i < squaresInRow; i++ {
		for j := 0; j < squaresInRow; j++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Scale(scale, scale)
			op.GeoM.Translate(float64(i*squareSize), float64(j*squareSize))

			screen.DrawImage(tilesObj.GetTileAt(0, 2), op)

		}
	}

}
