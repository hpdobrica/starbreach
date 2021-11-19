package main

import (
	"bytes"
	"embed"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hpdobrica/starbreach/tileutil"
)

//go:embed assets/*
var resources embed.FS

type Game struct {
	tiles *tileutil.Tiles
}

const squaresInRow = 8
const squareSize = 50
const boardSize = squareSize * squaresInRow

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello world!")
	drawBoard(screen, g.tiles)
	tmpWriteNumbersOnBoard(screen)

	// drawPawn(screen, 1)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return boardSize, boardSize
}

func main() {
	ebiten.SetWindowSize(boardSize*2, boardSize*2)
	ebiten.SetWindowTitle("Hello world")

	Tiles := tileutil.BuildTiles(resources, "assets/nature-tileset.png", 32)

	if err := ebiten.RunGame(&Game{tiles: Tiles}); err != nil {
		log.Fatal(err)
	}
}

func drawBoard(screen *ebiten.Image, Tiles *tileutil.Tiles) {

	scale := float64(squareSize) / float64(32)

	for i := 0; i < squaresInRow; i++ {
		for j := 0; j < squaresInRow; j++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Scale(scale, scale)
			op.GeoM.Translate(float64(i*squareSize), float64(j*squareSize))

			screen.DrawImage(Tiles.GetTileAt(0, 2), op)

		}
	}

}

func tmpWriteNumbersOnBoard(screen *ebiten.Image) {
	for i := 0; i < squaresInRow; i++ {
		for j := 0; j < squaresInRow; j++ {
			id := j*squaresInRow + i + 1
			ebitenutil.DebugPrintAt(screen, strconv.Itoa(id), i*squareSize, j*squareSize)
		}
	}
}

// func drawPawn(screen *ebiten.Image, id int) {
// 	pawn, _, err := ebitenutil.NewImageFromFile("img/pawn.png")
// 	w, h := pawn.Size()

// 	if err != nil {
// 		log.Fatal("something went wrong", err)
// 	}

// 	x, y := idToCoordinates(4)

// 	op := &ebiten.DrawImageOptions{}

// 	op.GeoM.Translate(float64(-w/2), float64(-h/2))

// 	op.GeoM.Translate(float64(x), float64(y))

// 	screen.DrawImage(pawn, op)

// }

func idToCoordinates(id int) (x, y int) {
	x = ((id-1)%squaresInRow)*squareSize + (squareSize / 2)
	y = ((id-1)/squaresInRow)*squareSize + (squareSize / 2)

	fmt.Println(id%squaresInRow, id/squaresInRow)
	fmt.Println(x, y)
	// fmt.Println(20 % 8)

	return
}

func getImage(path string) *ebiten.Image {
	imgByte, errRead := resources.ReadFile(path)
	if errRead != nil {
		log.Fatal(errRead)
	}

	imgDecoded, _, errDecode := image.Decode(bytes.NewReader(imgByte))

	if errDecode != nil {
		log.Fatal(errRead)
	}

	return ebiten.NewImageFromImage(imgDecoded)
}
