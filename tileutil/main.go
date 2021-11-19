package tileutil

import (
	"bytes"
	"embed"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Tiles struct {
	asset    *ebiten.Image
	tileSize int
}

func BuildTiles(resources embed.FS, path string, tileSize int) *Tiles {
	img := getImage(resources, path)

	return &Tiles{
		asset:    img,
		tileSize: tileSize,
	}
}

func (t Tiles) GetTileAt(x, y int) *ebiten.Image {
	x1 := x * t.tileSize
	x2 := y * t.tileSize

	y1 := x1 + t.tileSize
	y2 := x2 + t.tileSize

	return t.asset.SubImage(image.Rect(x1, x2, y1, y2)).(*ebiten.Image)

}

func getImage(resources embed.FS, path string) *ebiten.Image {
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
