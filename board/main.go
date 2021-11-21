package board

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hpdobrica/starbreach/game"
	"github.com/hpdobrica/starbreach/tiles"
)

type Board struct {
	layers       [][][]int
	owningPlayer int

	squareSize   int
	squaresInRow int
	tiles        *tiles.Tiles
}

func NewBoard(game *game.Game, terrainSpriteX, terrainSpriteY, owningPlayer, squareSize, squaresInRow int, tiles *tiles.Tiles) *Board {

	layers := [][][]int{
		{
			{-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1},
			{-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1},
			{-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1},
			{-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1},
			{-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1},
			{-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1},
			{-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1},
			{-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1},
		},
		{
			{-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1},
			{-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1},
			{-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1},
			{-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1},
			{-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1},
			{-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1},
			{-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1},
			{-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1}, {-1, -1},
		},
	}

	for i := range layers[0] {
		layers[0][i] = []int{terrainSpriteX, terrainSpriteY}
	}

	return &Board{layers: layers, owningPlayer: owningPlayer, squareSize: squareSize, squaresInRow: squaresInRow, tiles: tiles}
}

func (b Board) Draw(screen *ebiten.Image) {

	scale := float64(b.squareSize) / float64(b.tiles.TileSize)

	for _, layer := range b.layers {
		for n, square := range layer {
			x := n / b.squaresInRow
			y := n % b.squaresInRow

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Scale(scale, scale)
			op.GeoM.Translate(float64(x*b.squareSize), float64(y*b.squareSize))

			screen.DrawImage(b.tiles.GetTileAt(square[0], square[1]), op)
		}
	}

	// for i := 0; i < b.squaresInRow; i++ {
	// 	for j := 0; j < b.squaresInRow; j++ {
	// 		op := &ebiten.DrawImageOptions{}
	// 		op.GeoM.Scale(scale, scale)
	// 		op.GeoM.Translate(float64(i*b.squareSize), float64(j*b.squareSize))

	// 		screen.DrawImage(b.tiles.GetTileAt(0, 2), op)

	// 	}
	// }

}
