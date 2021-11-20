package turnstates

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hpdobrica/starbreach/game"
	"github.com/hpdobrica/starbreach/state"
)

type TurnTransitionState struct {
	game         *game.Game
	background   *ebiten.Image
	stateMachine *state.StateMachine
	turnNumber   int
	player       int
}

func NewTurnTransitionState(game *game.Game, turnNumber, player int) *TurnTransitionState {
	return &TurnTransitionState{game: game}
}

func (t *TurnTransitionState) Init(stateMachine *state.StateMachine) {
	t.stateMachine = stateMachine
}

func (t *TurnTransitionState) Update() {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		actionState := NewTurnActionState(t.game, t.turnNumber, t.player)
		t.stateMachine.PopState()
		t.stateMachine.PushState(actionState)
	}
}

func (t *TurnTransitionState) Draw(screen *ebiten.Image) {
	w, h := screen.Size()

	if t.background == nil {
		t.background = ebiten.NewImage(w, h)
	}

	t.background.Fill(color.Black)

	op := &ebiten.DrawImageOptions{}

	screen.DrawImage(t.background, op)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Player %d, Turn %d - Press space to continue", t.player, t.turnNumber), 10, h/2-10)
}
