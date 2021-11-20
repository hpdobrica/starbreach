package state

import "github.com/hajimehoshi/ebiten/v2"

type State interface {
	Init()

	// Pause()
	// Resume()

	// HandleEvents()
	Update()
	Draw(*ebiten.Image)
}
