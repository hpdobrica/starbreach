package gamestates

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hpdobrica/starbreach/game"
	"github.com/hpdobrica/starbreach/tiles"
)

type DeployState struct {
	game *game.Game
}

func NewDeployState(game *game.Game) *DeployState {
	return &DeployState{game: game}
}

func (i DeployState) Init() {

}

func (i *DeployState) Update() {

}

func (i *DeployState) Draw(screen *ebiten.Image) {

	drawBoard(screen, i.game.SquareSize, i.game.SquaresInRow, i.game.Tiles)

}

func drawBoard(screen *ebiten.Image, squareSize int, squaresInRow int, tilesObj *tiles.Tiles) {

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
