package interfaces

import "github.com/hajimehoshi/ebiten/v2"

type Updatable interface {
	Update()
	Draw(screen *ebiten.Image)
	Destroy()
	GetId() int64
}
