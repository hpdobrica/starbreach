package main

import (
	"bytes"
	"embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hpdobrica/starbreach/game"
	"github.com/hpdobrica/starbreach/state"
	gamestates "github.com/hpdobrica/starbreach/states/game"
	"github.com/hpdobrica/starbreach/tiles"
)

//go:embed assets/*
var resources embed.FS

func main() {
	squaresInRow := 8
	squareSize := 50
	boardSize := squaresInRow * squareSize
	ebiten.SetWindowSize(boardSize*2, boardSize*2)
	ebiten.SetWindowTitle("Hello world")

	Tiles := tiles.New(resources, "assets/nature-tileset.png", 32)

	gameObj := game.NewGame(Tiles, squaresInRow, squareSize)

	gameStateMachine := buildGameStateMachine(gameObj)

	gameObj.Init(gameStateMachine)

	if err := ebiten.RunGame(gameObj); err != nil {
		log.Fatal(err)
	}
}

func buildGameStateMachine(gameObj *game.Game) *state.StateMachine {
	stateMachine := state.NewStateMachine()

	introState := gamestates.NewIntroState(300, gameObj)

	stateMachine.PushState(introState)

	return stateMachine

}

// func drawBoard(screen *ebiten.Image, Tiles *tiles.Tiles) {

// 	scale := float64(squareSize) / float64(32)

// 	for i := 0; i < squaresInRow; i++ {
// 		for j := 0; j < squaresInRow; j++ {
// 			op := &ebiten.DrawImageOptions{}
// 			op.GeoM.Scale(scale, scale)
// 			op.GeoM.Translate(float64(i*squareSize), float64(j*squareSize))

// 			screen.DrawImage(Tiles.GetTileAt(0, 2), op)

// 		}
// 	}

// }

// func tmpWriteNumbersOnBoard(screen *ebiten.Image) {
// 	for i := 0; i < squaresInRow; i++ {
// 		for j := 0; j < squaresInRow; j++ {
// 			id := j*squaresInRow + i + 1
// 			ebitenutil.DebugPrintAt(screen, strconv.Itoa(id), i*squareSize, j*squareSize)
// 		}
// 	}
// }

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

// func idToCoordinates(id int) (x, y int) {
// 	x = ((id-1)%squaresInRow)*squareSize + (squareSize / 2)
// 	y = ((id-1)/squaresInRow)*squareSize + (squareSize / 2)

// 	fmt.Println(id%squaresInRow, id/squaresInRow)
// 	fmt.Println(x, y)
// 	// fmt.Println(20 % 8)

// 	return
// }

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
