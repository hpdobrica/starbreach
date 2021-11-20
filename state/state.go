package state

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type State interface {
	Init(*StateMachine)

	// Pause()
	// Resume()

	// HandleEvents()
	Update()
	Draw(*ebiten.Image)
}
