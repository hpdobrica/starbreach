package turnstates

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hpdobrica/starbreach/game"
	"github.com/hpdobrica/starbreach/state"
)

type TurnActionState struct {
	game         *game.Game
	background   *ebiten.Image
	stateMachine *state.StateMachine
	turnNumber   int
	player       int
}

func NewTurnActionState(game *game.Game, turnNumber, player int) *TurnActionState {
	return &TurnActionState{game: game}
}

func (t *TurnActionState) Init(stateMachine *state.StateMachine) {
	t.stateMachine = stateMachine
}

func (t *TurnActionState) Update() {

}

func (t *TurnActionState) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "now you play out your turn")
}
