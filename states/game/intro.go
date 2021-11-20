package gamestates

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hpdobrica/starbreach/game"
	"github.com/hpdobrica/starbreach/util"
)

type IntroState struct {
	duration        int
	background      *ebiten.Image
	initialDuration int
	game            *game.Game
}

func NewIntroState(duration int, game *game.Game) *IntroState {
	return &IntroState{duration: duration, initialDuration: duration, game: game}
}

func (i IntroState) Init() {

}

func (i *IntroState) Update() {
	i.duration -= 1
	fmt.Println(i.duration)

	if i.duration <= 0 {
		mainMenuState := NewMainMenuState(i.game)
		i.game.StateMachine.PopState()
		i.game.StateMachine.PushState(mainMenuState)
		// transition to other state
	}
}

func (i *IntroState) Draw(screen *ebiten.Image) {
	w, h := screen.Size()

	if i.background == nil {
		i.background = ebiten.NewImage(w, h)
	}

	var currentShade uint8
	if i.duration > 100 {
		t := float64(i.duration-100) / float64(i.initialDuration)

		currentShade = uint8(util.Lerp(0, 255, t))
	} else {
		currentShade = 0
	}

	i.background.Fill(color.RGBA{R: currentShade, G: currentShade, B: currentShade, A: 255})

	op := &ebiten.DrawImageOptions{}

	screen.DrawImage(i.background, op)
	ebitenutil.DebugPrintAt(screen, "Starbreach", w/2, h/2)
}
